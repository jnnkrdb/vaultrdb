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

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
)

// VRDBConfigReconciler reconciles a VRDBConfig object
type VRDBConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=jnnkrdb.de,resources=vrdbconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=jnnkrdb.de,resources=vrdbconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=jnnkrdb.de,resources=vrdbconfigs/finalizers,verbs=update

func (r *VRDBConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var _log = log.FromContext(ctx).WithName("vrdbconfig reconciler").WithValues("namespace", req.Namespace, "name", req.Name)
	ctx = log.IntoContext(ctx, _log)
	ctx = context.WithValue(ctx, jnnkrdbdev1.VRDBKey{}, req.NamespacedName)

	var vrdbconfig = &jnnkrdbdev1.VRDBConfig{}

	// checking the requested object
	if err := r.Get(ctx, req.NamespacedName, vrdbconfig, &client.GetOptions{}); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		_log.Error(err, "error reconciling vrdbconfig")
		return ctrl.Result{}, err
	}

	// check finalization
	if res, finalized, err := jnnkrdbdev1.Finalize(ctx, r.Client, vrdbconfig); err != nil || finalized {
		return res, err
	}

	// reconcile
	return vrdbconfig.Reconcile(ctx)
}

// SetupWithManager sets up the controller with the Manager.
func (r *VRDBConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&jnnkrdbdev1.VRDBConfig{}).
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
}
