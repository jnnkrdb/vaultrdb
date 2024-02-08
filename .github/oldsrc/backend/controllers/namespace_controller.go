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
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
	"github.com/jnnkrdb/vaultrdb/controllers/vaultrequest"
)

// NamespaceReconciler reconciles a Namespace object
type NamespaceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Namespace object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *NamespaceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	// initialize the logger
	var _log = log.FromContext(ctx).WithName(fmt.Sprintf("namespace [%s]", req.NamespacedName))

	// create caching object
	// cache namespace
	var ns = &v1.Namespace{}

	// parse the ctrl.Request into a namespace
	if err := r.Get(ctx, req.NamespacedName, ns, &client.GetOptions{}); err != nil {
		// if the error is an "NotFound" error, then the namespace probably was deleted
		// returning no error
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		_log.V(0).Error(err, "error reconciling vaultrequest at namespace event")
		// if the error is something else, return the error
		return ctrl.Result{}, err
	}

	// since the namespace was found, start the checks for the namespace
	_log.V(0).Info("namespace changed")

	// requesting the list of vaultrequests, existing in the cluster
	var vaultrequestList = &jnnkrdbdev1.VaultRequestList{}
	if err := r.List(ctx, vaultrequestList, &client.ListOptions{}); err != nil {
		_log.V(0).Error(err, "error listing the vaultrequests")
		return ctrl.Result{Requeue: true}, err
	}

	// for every item in the vaultrequestslist, start the reconcilation
	for _, vr := range vaultrequestList.Items {
		_log.V(1).Info("identified vaultrequest", "vr.name", vr.Name, "vr.Namespace", vr.Namespace)

		if _, err := vaultrequest.Reconcile(_log, ctx, r.Client, &vr); err != nil {
			return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, err
		}

		_log.V(1).Info("successfully reconciled vaultrequest", "vr.name", vr.Name, "vr.Namespace", vr.Namespace)
	}

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NamespaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.Namespace{}).
		// this eventfilter is set, to prevent reconcilation loops, because, if unset, the
		// reconcilation controller gets called, everytime the namespace gets updated,
		// even if the update occurs in metadata- or status-fields
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
}
