package vaultrequest

import (
	"context"

	"github.com/go-logr/logr"
	jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

//+kubebuilder:rbac:groups=core,resources=secrets;configmaps,verbs=get;create;update;patch;delete;list

// check the finalization for the objects
//
// if the finalizer is not appended to the object, it
// will be added, else, if the object is marked, to be deleted,
// the operations to remove the connected objects will be launched
func Finalize(_log logr.Logger, ctx context.Context, c client.Client, vr *jnnkrdbdev1.VaultRequest) (bool, error) {
	_log.V(0).Info("check finalization")

	const _finalizer string = "vaultrequest.jnnkrdb.de/v1.finalizer"

	// check, wether the vaultrequest has the required finalizer or not
	// if not, then add the finalizer
	if controllerutil.ContainsFinalizer(vr, _finalizer) {
		_log.V(1).Info("appending finalizer")

		// add the desired finalizer and update the object
		controllerutil.AddFinalizer(vr, _finalizer)

		// update the vaultrequest, with new finalizer
		if err := c.Update(ctx, vr); err != nil {
			_log.V(0).Error(err, "error adding finalizer")
			return false, err
		}
	}

	// receive the new version of the updated vaultrequest
	if err := c.Get(ctx, types.NamespacedName{Namespace: vr.Namespace, Name: vr.Name}, vr); err != nil {
		_log.V(0).Error(err, "error updating cached object")
		return false, err
	}

	// check, if the vaultrequest is marked to be deleted
	if vr.GetDeletionTimestamp() != nil {
		// check, wether the vaultrequest has the required finalizer or not
		if controllerutil.ContainsFinalizer(vr, _finalizer) {
			// start the finalizing routine
			_log.V(1).Info("finalizing vaultrequest")

			// remove all objects from the status.Deployed field
			for _, i := range vr.Status.Deployed {

				var l = _log.V(3).WithValues("kind", i.Kind, "namespace", i.Namespace, "name", vr.Name)
				l.V(3).Info("finalizing object")

				// get the kind of the object and remove the actual object
				switch i.Kind {

				case "ConfigMap":
					var cm = &v1.ConfigMap{}
					// get the object from declarations
					if err := c.Get(ctx, types.NamespacedName{Namespace: i.Namespace, Name: vr.Name}, cm); err != nil {
						// if the error is an "NotFound" error, then the configmap probably was deleted
						// returning no error
						if errors.IsNotFound(err) {
							l.V(0).Info("object not found")
							continue
						}

						l.V(0).Error(err, "error receiving object from namespace and name")
						return false, err
					}
					// remove the cached object from the cluster
					if err := c.Delete(ctx, cm); err != nil {
						l.V(0).Error(err, "error removing the object")
						return false, err
					}

				case "Secret":
					var scrt = &v1.Secret{}
					// get the object from declarations
					if err := c.Get(ctx, types.NamespacedName{Namespace: i.Namespace, Name: vr.Name}, scrt); err != nil {
						// if the error is an "NotFound" error, then the secret probably was deleted
						// returning no error
						if errors.IsNotFound(err) {
							l.V(3).Info("object not found")
							continue
						}

						l.V(0).Error(err, "error receiving object from namespace and name")
						return false, err
					}
					// remove the cached object from the cluster
					if err := c.Delete(ctx, scrt); err != nil {
						l.V(0).Error(err, "error removing the object")
						return false, err
					}

				default:
					l.V(3).Info("the kind is unknown... skipped")
					continue
				}

				// implement the status update
				var newStatus = jnnkrdbdev1.Remove(vr.Status.Deployed, i.Kind, i.Namespace)
				_log.V(3).Info("new status identified", "newStatus", newStatus)

				// receive the new version of the updated vaultrequest
				if err := c.Get(ctx, types.NamespacedName{Namespace: vr.Namespace, Name: vr.Name}, vr); err != nil {
					_log.V(0).Error(err, "error updating cached object")
					return false, err
				}

				vr.Status.Deployed = newStatus
				if err := c.Status().Update(ctx, vr); err != nil {
					_log.V(0).Error(err, "error updating status from vaultrequest")
					return false, err
				}

				l.V(3).Info("object removed")
			}

			_log.V(1).Info("finished finalizing vaultrequests")

			// receive the new version of the updated vaultrequest
			if err := c.Get(ctx, types.NamespacedName{Namespace: vr.Namespace, Name: vr.Name}, vr); err != nil {
				_log.V(0).Error(err, "error updating cached object")
				return false, err
			}

			// remove the finalizer from the vaultrequest
			controllerutil.RemoveFinalizer(vr, _finalizer)
			if err := c.Update(ctx, vr); err != nil {
				_log.V(0).Error(err, "error removing finalizer flag from cluster object")
				return false, err
			}

			// return the finalized state (true) and nil error
			return true, nil
		}
	}

	return false, nil
}
