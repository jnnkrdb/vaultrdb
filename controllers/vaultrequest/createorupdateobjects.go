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

	// run pre flight validations and get the datamap
	objectData, valid, err := receiveDataMap(_log, vr.Spec.DataMap)
	switch {
	case err != nil:
		return true, ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
	case !valid:
		return true, ctrl.Result{Requeue: false}, nil
	}

	// creating the validation labels
	var lbls = make(map[string]string)
	for k := range objectData {
		lbls[k] = "validkey"
	}

	// set the annotation, to identify the source vaultrequest, if neccessary
	var annotations = map[string]string{
		"v1.vaultrequest.jnnkrdb.de/source": fmt.Sprintf("%s/%s", vr.Namespace, vr.Name),
	}

	// create or update the configmap/secret from the matching namespace
	for i := range matchList {
		_log = _log.WithValues("kind", vr.Spec.Namespaces.Kind, "namespace", matchList[i].Name, "name", vr.Name)

		var obj client.Object
		var reqErr error

		switch vr.Spec.Namespaces.Kind {
		case "Secret":
			_log.Info("checking existance of the secret")
			var s = &v1.Secret{}

			if reqErr = c.Get(ctx, types.NamespacedName{
				Namespace: matchList[i].Name,
				Name:      vr.Name,
			}, s, &client.GetOptions{}); reqErr != nil && !errors.IsNotFound(reqErr) {
				_log.Error(reqErr, "error requesting object information")
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
			_log.Info("checking existance of the configmap")
			var cm = &v1.ConfigMap{}

			if reqErr = c.Get(ctx, types.NamespacedName{
				Namespace: matchList[i].Name,
				Name:      vr.Name,
			}, cm, &client.GetOptions{}); reqErr != nil && !errors.IsNotFound(reqErr) {
				_log.Error(reqErr, "error requesting object information")
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
		if errors.IsNotFound(reqErr) {
			_log.Info("try creating object")
			err = c.Create(ctx, obj, &client.CreateOptions{})
		} else {
			_log.Info("try updating object")
			err = c.Update(ctx, obj, &client.UpdateOptions{})
		}

		if err != nil {
			_log.Error(err, "couldn't execute objectprocess")
			return true, ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
		}

		// updating status of the vaultrequest
		if err = c.Get(ctx, types.NamespacedName{
			Namespace: vr.Namespace,
			Name:      vr.Name,
		}, vr, &client.GetOptions{}); err != nil {
			_log.Error(err, "couldn't re-cache vaultrequest")
			return true, ctrl.Result{Requeue: true}, err
		}

		var alreadyStatusObject bool = false
		for o := range vr.Status.Deployed {
			if vr.Status.Deployed[o].Namespace == obj.GetNamespace() && vr.Status.Deployed[o].Name == obj.GetName() {
				alreadyStatusObject = true
				break
			}
		}

		if !alreadyStatusObject {
			vr.Status.Deployed = append(vr.Status.Deployed, jnnkrdbdev1.DeployedObject{
				Kind:      vr.Spec.Namespaces.Kind,
				Namespace: matchList[i].Name,
				Name:      vr.Name,
			})
			if err = c.Status().Update(ctx, vr); err != nil {
				_log.Error(err, "could not update the status of the vaultrequest")
				return true, ctrl.Result{Requeue: true}, err
			}
		}
	}

	return false, ctrl.Result{}, nil
}
