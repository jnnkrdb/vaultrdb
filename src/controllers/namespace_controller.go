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

	jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// NamespaceReconciler reconciles a Namespace object
type NamespaceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;

func (r *NamespaceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var _log = log.FromContext(ctx).WithName("namespace").WithValues("namespace", req.NamespacedName)
	ctx = log.IntoContext(ctx, _log)
	var namespace = &v1.Namespace{}

	if err := r.Get(ctx, req.NamespacedName, namespace, &client.GetOptions{}); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		_log.Error(err, "error reconciling vrdb types")
		return ctrl.Result{}, err
	}

	_log.Info("namespace changed")

	var vrdbconfigs = &jnnkrdbdev1.VRDBConfigList{}
	if err := r.List(ctx, vrdbconfigs, &client.ListOptions{}); err != nil {
		_log.Error(err, "error receiving list of vrdbconfigs")
		return ctrl.Result{}, err
	} else {
		for _, item := range vrdbconfigs.Items {
			return item.Reconcile(ctx, r.Client)
		}
	}

	var vrdbsecrets = &jnnkrdbdev1.VRDBSecretList{}
	if err := r.List(ctx, vrdbsecrets, &client.ListOptions{}); err != nil {
		_log.Error(err, "error receiving list of vrdbsecrets")
		return ctrl.Result{}, err
	} else {
		for _, item := range vrdbsecrets.Items {
			return item.Reconcile(ctx, r.Client)
		}
	}

	//var vrdbrequests = &jnnkrdbdev1.VRDBRequestList{}
	//if err := r.List(ctx, vrdbrequests, &client.ListOptions{}); err != nil {
	//	_log.Error(err, "error receiving list of vrdbrequests")
	//	return ctrl.Result{}, err
	//} else {
	//	for _, item := range vrdbrequests.Items {
	//		return item.Reconcile(ctx, r.Client)
	//	}
	//}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NamespaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.Namespace{}).
		Complete(r)
}
