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

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
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

	// start processing the datamap fields
	for datamapKey, datamap := range vaultreq.Spec.DataMap {

		_log = _log.WithValues("datamapkey", datamapKey)

		// first validate the datafield
		if err := datamap.Validate(_log); err != nil {
			_log.Error(err, "validating datafields failed")
			return ctrl.Result{Requeue: true, RequeueAfter: time.Minute * 2}, err
		}
	}

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *VaultRequestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&jnnkrdbdev1.VaultRequest{}).
		// this eventfilter is set, to prevent reconcilation loops, because, if unset, the
		// reconcilation controller gets called, everytime the deployrequest gets updated,
		// even if the update occurs in metadata- or status-fields
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
}