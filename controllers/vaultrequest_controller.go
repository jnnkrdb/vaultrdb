/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
)

// VaultRequestReconciler reconciles a VaultRequest object
type VaultRequestReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=jnnkrdb.de,resources=vaultrequests,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=jnnkrdb.de,resources=vaultrequests/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=jnnkrdb.de,resources=vaultrequests/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the VaultRequest object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *VaultRequestReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	// initialize the logger
	var _log = log.FromContext(ctx).WithName(fmt.Sprintf("vaultrequest [%s]", req.NamespacedName))

	// create caching object
	// cache vaultrequests
	var vaultreq = &jnnkrdbdev1.VaultRequest{}

	// parse the ctrl.Request into a vaultrequest
	if err := r.Get(ctx, req.NamespacedName, vaultreq); err != nil {
		// if the error is an "NotFound" error, then the vaultrequest probably was deleted
		// returning no error
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		_log.Error(err, "error reconciling vaultrequest")
		// if the error is something else, return the error
		return ctrl.Result{}, err
	}

	// check if the object contains the finalization flags, or has to be terminated
	if finalized, err := checkFinalizationVaultRequest(_log, ctx, r, vaultreq); err != nil || finalized {
		return ctrl.Result{Requeue: !finalized}, err
	}

	_log.Info("start reconciling")

	// validate the date, provided by the dataref fields. First check if there are any keys provided, if not
	// return with a nil error and no requeuing
	if len(vaultreq.Spec.DataMap) <= 0 {
		_log.Info("empty datamaps are not allowed")
		return ctrl.Result{}, nil
	}

	var objectData = make(map[string]string)

	// start processing the datamap fields
	for datamapKey, datamap := range vaultreq.Spec.DataMap {

		_log = _log.WithValues("datamapkey", datamapKey)

		// get the data for the key
		dat, e := datamap.GetData(_log)
		if e != nil {
			_log.Error(e, "invalid datamap object")
			return ctrl.Result{Requeue: false}, e
		}

		objectData[datamapKey] = dat
	}

	// calculating the namespaces
	match, avoid, err := vaultreq.Spec.Namespaces.CalculateNamespaces(_log, ctx, r.Client)
	if err != nil {
		_log.Error(err, "couldn't calculate namespaces")
		return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
	}

	// remove the configmap/secrets, which should not exist anymore
	for _, statusObject := range vaultreq.Status.Deployed {

		nsLog := _log.WithValues("kind", statusObject.Kind, "namespace", statusObject.Namespace, "name", statusObject.Name)

		// check if the namespace of the status object is in the avoid namespaces
		// array
		for _, avoidNamespace := range avoid {
			if avoidNamespace.Name != statusObject.Namespace {
				continue
			}
		}

		switch statusObject.Kind {
		case "Secret":
			// check if the secret exists
			var s = &v1.Secret{}
			if err = r.Get(ctx, types.NamespacedName{Namespace: statusObject.Namespace, Name: statusObject.Name}, s, &client.GetOptions{}); err != nil && !errors.IsNotFound(err) {
				nsLog.Error(err, "couldn't receive object information")
				return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
			}

			// remove the secret from the cluster
			if err = r.Delete(ctx, s, &client.DeleteOptions{}); err != nil {
				nsLog.Error(err, "couldn't remove object")
				return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
			}

		case "ConfigMap":
			// check if the configmap exists
			var cm = &v1.ConfigMap{}
			if err = r.Get(ctx, types.NamespacedName{Namespace: statusObject.Namespace, Name: statusObject.Name}, cm, &client.GetOptions{}); err != nil && !errors.IsNotFound(err) {
				nsLog.Error(err, "couldn't receive object information")
				return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
			}

			// remove the configmap from the cluster
			if err = r.Delete(ctx, cm, &client.DeleteOptions{}); err != nil {
				nsLog.Error(err, "couldn't remove object")
				return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
			}
		}

	}

	// create or update the configmaps/secrets from in the match namespaces
	for _, matchNamespace := range match {

		nsLog := _log.WithValues("kind", vaultreq.Spec.Namespaces.Kind, "namespace", matchNamespace.Name)

		var lbls = make(map[string]string)
		for k := range objectData {
			lbls[k] = "validkey"
		}

		// check the kind of the requested resource
		switch vaultreq.Spec.Namespaces.Kind {
		case "Secret":

			var s = &v1.Secret{}

			// checking existance of the secret
			if err = r.Get(ctx, types.NamespacedName{Namespace: matchNamespace.Name, Name: vaultreq.Name}, s, &client.GetOptions{}); err != nil && !errors.IsNotFound(err) {
				nsLog.Error(err, "couldn't validate, if object exists or not")
				return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
			}

			s.StringData = objectData
			s.Labels = lbls
			s.Type = v1.SecretType(vaultreq.Spec.Namespaces.Type)

			// if the object is not found, then create the object
			// else, update the obkect
			if errors.IsNotFound(err) {

				// fill the neccessary informations
				s.Name = vaultreq.Name
				s.Namespace = matchNamespace.Name

				if err = r.Create(ctx, s, &client.CreateOptions{}); err != nil {
					nsLog.Error(err, "couldn't creating existing object")
					return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
				}

			} else {

				if err = r.Update(ctx, s, &client.UpdateOptions{}); err != nil {
					nsLog.Error(err, "couldn't update existing object")
					return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
				}
			}

		case "ConfigMap":

			var cm = &v1.ConfigMap{}

			// checking existance of the configmap
			if err = r.Get(ctx, types.NamespacedName{Namespace: matchNamespace.Name, Name: vaultreq.Name}, cm, &client.GetOptions{}); err != nil && !errors.IsNotFound(err) {
				nsLog.Error(err, "couldn't validate, if object exists or not")
				return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
			}

			cm.Data = objectData
			cm.Labels = lbls

			// if the object is not found, then create the object
			// else, update the obkect
			if errors.IsNotFound(err) {

				// fill the neccessary informations
				cm.Name = vaultreq.Name
				cm.Namespace = matchNamespace.Name

				if err = r.Create(ctx, cm, &client.CreateOptions{}); err != nil {
					nsLog.Error(err, "couldn't creating existing object")
					return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
				}

			} else {

				if err = r.Update(ctx, cm, &client.UpdateOptions{}); err != nil {
					nsLog.Error(err, "couldn't update existing object")
					return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
				}
			}
		}
	}

	// TODO(user): your logic here

	return ctrl.Result{
		Requeue:      true,
		RequeueAfter: 5 * time.Minute,
	}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *VaultRequestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&jnnkrdbdev1.VaultRequest{}).
		// this eventfilter is set, to prevent reconcilation loops, because, if unset, the
		// reconcilation controller gets called, everytime the deployrequest gets updated,
		// even if the update occurs in metadata- or status-fields
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		// trying to watch namespace, so when a namespace gets created, all the vaultrequests
		// will be synchronized again. That means, the rescheduling can be deactivated
		Owns(&v1.Namespace{}).
		Complete(r)
}
