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

package v1

import (
	"context"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// reconcilation function
func (v *VRDBSecret) Reconcile(ctx context.Context, c client.Client) (ctrl.Result, error) {
	var _log = log.FromContext(ctx)

	_log.Info("reconciling vrdbsecret")

	// calculate the neccessary namespace collections
	toMatch, toAvoid, err := v.NamespaceSelector.CalculateCollections(ctx, c)
	if err != nil {
		_log.Error(err, "error calculating the namespace collections")
		return ctrl.Result{RequeueAfter: 5 * time.Minute}, err
	}

	// remove the unwanted objects from the cluster
	if res, err := RemoveUnwantedObjectFromCluster(v, ctx, c, &v.Status, toAvoid, &v1.Secret{}); err != nil {
		return res, nil
	}

	// create or update the wanted objects
	for _, ns := range toMatch {
		var (
			__log       = _log.WithValues("requestedNamespace", ns)
			tempS       = &v1.Secret{}
			create bool = false
		)

		if err := c.Get(ctx, types.NamespacedName{Name: v.Name, Namespace: ns}, tempS, &client.GetOptions{}); err != nil {
			if create = errors.IsNotFound(err); !create {
				__log.Error(err, "error while requesting specific secret")
				return ctrl.Result{RequeueAfter: 5 * time.Minute}, err
			}
		}

		// set the secret type
		tempS.Type = v.Type
		// copy the data
		tempS.Data = make(map[string][]byte)
		for k, v := range v.Data {
			tempS.Data[k] = []byte(v)
		}

		// update the existing secret
		if !create {
			if err := c.Update(ctx, tempS, &client.UpdateOptions{}); err != nil {
				__log.Error(err, "error updating existing secret")
				return ctrl.Result{RequeueAfter: 1 * time.Minute}, err
			}
			continue
		}

		// create the object
		tempS.Name = v.Name
		tempS.Namespace = ns
		tempS.Annotations = map[string]string{
			Annotation_SourceNamespace: v.Namespace,
			Annotation_SourceName:      v.Name,
		}

		if err := c.Create(ctx, tempS, &client.CreateOptions{}); err != nil {
			__log.Error(err, "error creating new configmap")
			return ctrl.Result{RequeueAfter: 1 * time.Minute}, err
		}

		if !v.Status.ContainsNamespace(ns) {
			v.Status.Namespaces = append(v.Status.Namespaces, ns)

			if err := Update(ctx, c, true, v); err != nil {
				return ctrl.Result{RequeueAfter: 1 * time.Minute}, err
			}
		}
	}

	_log.Info("finished reconciling vrdbsecret")

	return ctrl.Result{}, nil
}
