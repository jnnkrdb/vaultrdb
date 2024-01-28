package vaultrequest

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateOrUpdateObjects(_log logr.Logger, ctx context.Context, c client.Client, vr *jnnkrdbdev1.VaultRequest, matchList []v1.Namespace) (bool, ctrl.Result, error) {

	// validate the data, provided by the dataref fields
	// if the datamap [mp] is empty, then there is no error while validation,
	// but the validation check will be false
	if length := len(vr.Spec.DataMap); length <= 0 {
		_log.V(1).Info("empty datamaps are not allowed", "len(datamap)", length)
		return true, ctrl.Result{Requeue: false}, nil
	}

	var objectData = make(map[string]string)
	// start processing the datamap fields
	for dmKey, dmValue := range vr.Spec.DataMap {
		// get the data for the key
		if dat, err := dmValue.GetData(); err != nil {
			_log.V(0).Error(err, "invalid datamap object", "datamapkey", dmKey)
			return true, ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
		} else {
			objectData[dmKey] = dat
		}
	}

	// creating the validation labels
	var lbls = make(map[string]string)
	for k := range objectData {
		lbls[fmt.Sprintf("v1.vaultrequest.jnnkrdb.de_%s.key", k)] = "validated"
	}

	// set the annotation, to identify the source vaultrequest, if neccessary
	var annotations = map[string]string{
		"v1.vaultrequest.jnnkrdb.de/source.name":      vr.Name,
		"v1.vaultrequest.jnnkrdb.de/source.namespace": vr.Namespace,
	}

	// create or update the configmap/secret from the matching namespace
	for i := range matchList {
		var l = _log.WithValues("kind", vr.Spec.Namespaces.Kind, "namespace", matchList[i].Name, "name", vr.Name)

		var obj client.Object
		var reqErr error

		switch vr.Spec.Namespaces.Kind {
		case "Secret":
			l.V(3).Info("checking existance of the secret")
			var s = &v1.Secret{}

			if reqErr = c.Get(ctx, types.NamespacedName{
				Namespace: matchList[i].Name,
				Name:      vr.Name,
			}, s, &client.GetOptions{}); reqErr != nil && !errors.IsNotFound(reqErr) {
				l.V(0).Error(reqErr, "error requesting object information")
				return true, ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, reqErr
			}

			s.Name = vr.Name
			s.Namespace = matchList[i].Name
			s.Type = v1.SecretType(vr.Spec.Namespaces.Type)
			s.StringData = objectData
			s.Labels = lbls
			s.Annotations = annotations

			obj = s

		case "ConfigMap":
			l.V(3).Info("checking existance of the configmap")
			var cm = &v1.ConfigMap{}

			if reqErr = c.Get(ctx, types.NamespacedName{
				Namespace: matchList[i].Name,
				Name:      vr.Name,
			}, cm, &client.GetOptions{}); reqErr != nil && !errors.IsNotFound(reqErr) {
				l.V(0).Error(reqErr, "error requesting object information")
				return true, ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, reqErr
			}

			cm.Name = vr.Name
			cm.Namespace = matchList[i].Name
			cm.Data = objectData
			cm.Labels = lbls
			cm.Annotations = annotations

			obj = cm
		}

		// if the object is not found, then create the object
		// else, update the object
		var err error
		if errors.IsNotFound(reqErr) {
			l.V(3).Info("try creating object")
			err = c.Create(ctx, obj, &client.CreateOptions{})
		} else {
			l.V(3).Info("try updating object")
			err = c.Update(ctx, obj, &client.UpdateOptions{})
		}

		if err != nil {
			l.V(0).Error(err, "couldn't execute objectprocess")
			return true, ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
		}

		// updating status of the vaultrequest
		if err = c.Get(ctx, types.NamespacedName{
			Namespace: vr.Namespace,
			Name:      vr.Name,
		}, vr, &client.GetOptions{}); err != nil {
			l.V(0).Error(err, "couldn't re-cache vaultrequest")
			return true, ctrl.Result{Requeue: true}, err
		}

		if !jnnkrdbdev1.Contains(vr.Status.Deployed, obj.GetObjectKind().GroupVersionKind().Kind, obj.GetNamespace()) {

			vr.Status.Deployed = jnnkrdbdev1.Append(vr.Status.Deployed, vr.Spec.Namespaces.Kind, matchList[i].Name)

			if err = c.Status().Update(ctx, vr); err != nil {
				l.V(0).Error(err, "could not update the status of the vaultrequest")
				return true, ctrl.Result{Requeue: true}, err
			}
		}
	}

	return false, ctrl.Result{}, nil
}
