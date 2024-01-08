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

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type VRDBType struct{}

type VRDBStatus struct {
	Namespaces []string `json:"namespaces,omitempty"`
}

const (
	VRDBFinalizer = "vrdb.jnnkrdb.de/finalizer"
)

func Finalize(ctx context.Context, c client.Client, obj client.Object) (ctrl.Result, bool, error) {
	var _log = log.FromContext(ctx).WithValues("kind", obj.GetObjectKind().GroupVersionKind().Kind)

	_log.Info("finalizing object")

	// checking the object for the finalizer, if the finalizer does not exist
	// then append it to the object
	if !controllerutil.ContainsFinalizer(obj, VRDBFinalizer) {
		_log.Info("appending the finalizer to the object")

		controllerutil.AddFinalizer(obj, VRDBFinalizer)

		if err := Update(ctx, c, false, obj); err != nil {
			_log.Error(err, "error adding finalizer")
			return ctrl.Result{}, false, err
		}
	}

	// finalize if requested
	if obj.GetDeletionTimestamp() != nil &&
		controllerutil.ContainsFinalizer(obj, VRDBFinalizer) {

		switch obj.GetObjectKind().GroupVersionKind().Kind {
		case "VRDBConfig":

		case "VRDBSecret":

		case "VRDBRequest":

		}
	}

	return ctrl.Result{}, false, nil
}
