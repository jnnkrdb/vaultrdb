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
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type VRDBStatus struct {
	Namespaces []string `json:"namespaces,omitempty"`
}

// remove a namespace from the status
func (s *VRDBStatus) RemoveNamespace(n string) {
	if len(s.Namespaces) > 0 {
		var new []string

		for i := range s.Namespaces {
			if s.Namespaces[i] == n {
				continue
			}

			new = append(new, s.Namespaces[i])
		}

		s.Namespaces = new
	}
}

func (s VRDBStatus) ContainsNamespace(n string) bool {
	if len(s.Namespaces) > 0 {
		for i := range s.Namespaces {
			if s.Namespaces[i] == n {
				return true
			}
		}
	}
	return false
}

const (
	VRDBFinalizer = "vrdb.jnnkrdb.de/finalizer"
)

// finalizes the object if neccessary and returns the following states:
//
// 1. the return result for the reconcilation loop
//
// 2. a boolean, if the object was finalized or not
//
// 3. the error, if any occured
func Finalize(ctx context.Context, c client.Client, obj client.Object) (ctrl.Result, bool, error) {
	var _log = log.FromContext(ctx).WithValues("kind", obj.GetObjectKind().GroupVersionKind().Kind)

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

		_log.Info("finalizing object")

		findAndRemove := func(_o client.Object, status *VRDBStatus) (ctrl.Result, bool, error) {

			// find all objects and remove them from the cluster
			for _, ns := range status.Namespaces {
				if e := c.Get(ctx, types.NamespacedName{Namespace: ns, Name: obj.GetName()}, _o); e != nil {
					if errors.IsNotFound(e) {
						_log.Error(e, "object not found", "namespace", ns, "name", obj.GetName())
						continue
					}

					_log.Error(e, "error receiving object with namespace and name", "namespace", ns, "name", obj.GetName())
					return ctrl.Result{}, false, e
				}

				if e := c.Delete(ctx, _o); e != nil {
					_log.Error(e, "error removing object from cluster", "namespace", ns, "name", obj.GetName())
					return ctrl.Result{}, false, e
				}

				status.RemoveNamespace(ns)
				if e := Update(ctx, c, true, obj); e != nil {
					return ctrl.Result{}, false, e
				}
			}

			// remove the finalizer mark
			controllerutil.RemoveFinalizer(obj, VRDBFinalizer)
			if err := Update(ctx, c, false, obj); err != nil {
				_log.Error(err, "error removing finalizer")
				return ctrl.Result{}, false, err
			}

			return ctrl.Result{}, true, nil
		}

		switch obj.GetObjectKind().GroupVersionKind().Kind {
		case "VRDBConfig":
			if _, ok := obj.(*VRDBConfig); !ok {
				_log.Error(fmt.Errorf("error parsing object"), "couldn't parse object into VRDBConfig")
				return ctrl.Result{}, false, nil
			} else {
				return findAndRemove(&v1.ConfigMap{}, &obj.(*VRDBConfig).Status)
			}

		case "VRDBSecret":
			if _, ok := obj.(*VRDBSecret); !ok {
				_log.Error(fmt.Errorf("error parsing object"), "couldn't parse object into VRDBConfig")
				return ctrl.Result{}, false, nil
			} else {
				return findAndRemove(&v1.ConfigMap{}, &obj.(*VRDBSecret).Status)
			}

		case "VRDBRequest":
			if _, ok := obj.(*VRDBRequest); !ok {
				_log.Error(fmt.Errorf("error parsing object"), "couldn't parse object into VRDBConfig")
				return ctrl.Result{}, false, nil
			} else {
				return findAndRemove(&v1.ConfigMap{}, &obj.(*VRDBRequest).Status)
			}
		}
	}

	return ctrl.Result{}, false, nil
}
