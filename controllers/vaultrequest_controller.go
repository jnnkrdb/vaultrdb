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

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
	"github.com/jnnkrdb/vaultrdb/controllers/vaultrequest"
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
	var _log = log.FromContext(ctx).WithValues("vaultrequest", req.NamespacedName)

	// intro log
	_log.V(0).Info("starting reconcilation")

	// create caching object
	// cache vaultrequests
	var vaultreq = &jnnkrdbdev1.VaultRequest{}

	// parse the ctrl.Request into a vaultrequest
	if err := r.Get(ctx, req.NamespacedName, vaultreq); err != nil {
		// if the error is an "NotFound" error, then the vaultrequest probably was deleted
		// returning no error
		if errors.IsNotFound(err) {
			_log.V(3).Info("vaultrequest not found")
			return ctrl.Result{}, nil
		}
		_log.V(0).Error(err, "error reconciling vaultrequest")
		// if the error is something else, return the error
		return ctrl.Result{}, err
	}

	// run reconcilation of the identified vaultrequest
	return vaultrequest.Reconcile(_log, ctx, r.Client, vaultreq)
}

// SetupWithManager sets up the controller with the Manager.
func (r *VaultRequestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&jnnkrdbdev1.VaultRequest{}).
		// this eventfilter is set, to prevent reconcilation loops, because, if unset, the
		// reconcilation controller gets called, everytime the vaultrequest gets updated,
		// even if the update occurs in metadata- or status-fields
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		// trying to watch namespace, so when a namespace gets created, all the vaultrequests
		// will be synchronized again. That means, the rescheduling can be deactivated
		// Owns(&v1.Namespace{}).
		Complete(r)
}
