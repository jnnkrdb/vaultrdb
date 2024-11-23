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
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// reconcilation function
func (v *VRDBRequest) Reconcile(ctx context.Context, c client.Client) (ctrl.Result, error) {
	var _log = log.FromContext(ctx)

	_log.Info("reconciling vrdbrequest")

	// calculate the neccessary namespace collections
	toAvoid, toMatch, err := v.NamespaceSelector.CalculateCollections(ctx, c)
	if err != nil {
		_log.Error(err, "error calculating the namespace collections")
		return ctrl.Result{RequeueAfter: 5 * time.Minute}, err
	}

	// remove the unwanted objects from the cluster
	if res, err := RemoveUnwantedObjectFromCluster(v, ctx, c, &v.Status, toAvoid, &v1.Secret{}); err != nil {
		return res, nil
	}

	// create or update the wanted objects
	_log.Info("lists", "toMatch", toMatch, "toAvoid", toAvoid)

	_log.Info("finished reconciling vrdbrequest")

	return ctrl.Result{}, nil
}
