/*
Copyright 2024.

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

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
)

// VRDBRequestReconciler reconciles a VRDBRequest object
type VRDBRequestReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=jnnkrdb.de,resources=vrdbrequests,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=jnnkrdb.de,resources=vrdbrequests/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=jnnkrdb.de,resources=vrdbrequests/finalizers,verbs=update

func (r *VRDBRequestReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var _log = log.FromContext(ctx).WithName("vrdbrequest")
	ctx = log.IntoContext(ctx, _log)
	ctx = context.WithValue(ctx, types.NamespacedName{}, req.NamespacedName)

	var vrdbrequest = &jnnkrdbdev1.VRDBRequest{}

	// checking the requested object
	if res, err := jnnkrdbdev1.GetObjectFromCluster(ctx, r.Client, req, vrdbrequest, &client.GetOptions{}); err != nil {
		return res, err
	}

	// check finalization
	if res, finalized, err := jnnkrdbdev1.Finalize(ctx, r.Client, vrdbrequest); err != nil || finalized {
		return res, err
	}

	// reconcile
	return vrdbrequest.Reconcile(ctx, r.Client)
}

// SetupWithManager sets up the controller with the Manager.
func (r *VRDBRequestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&jnnkrdbdev1.VRDBRequest{}).
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
}
