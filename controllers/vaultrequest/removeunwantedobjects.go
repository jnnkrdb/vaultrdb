package vaultrequest

import (
	"context"

	"github.com/go-logr/logr"
	jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// remove secrets/configmaps from the cluster, which shouldn't exist anymore
func RemoveUnwantedObjects(_log logr.Logger, c client.Client, ctx context.Context, vr *jnnkrdbdev1.VaultRequest, avoidList []v1.Namespace) error {

	for indexDeployed, _do := range vr.Status.Deployed {
		_log = _log.WithValues("kind", _do.Kind, "namespace", _do.Namespace, "name", _do.Name)

		var remove bool = false

		// check all the namespace from the avoidList
		// if the avoidList contains the namespace of the deployed object,
		// then remove the deployed object
		for i := range avoidList {
			if avoidList[i].Name != _do.Namespace {
				remove = true
				break
			}
		}

		if !remove {
			_log.Info("object must not be removed")
			continue
		}

		var removeObject client.Object

		switch _do.Kind {
		case "ConfigMap":
			removeObject = &v1.ConfigMap{}
		case "Secret":
			removeObject = &v1.Secret{}
		default:
			_log.Info("object has unknown kind")
			continue
		}

		// check if the object exists
		if err := c.Get(ctx, types.NamespacedName{
			Namespace: _do.Namespace,
			Name:      _do.Name,
		}, removeObject, &client.GetOptions{}); err != nil {
			if errors.IsNotFound(err) {
				_log.Info("object not found")
				continue
			}
			_log.Error(err, "couldn't receive object informations")
			return err
		}

		// remove the actual object
		if err := c.Delete(ctx, removeObject, &client.DeleteOptions{}); err != nil {
			_log.Error(err, "couldn't remove object")
			return err
		}

		// create the new status from the original vaultrequest.Status
		var newStatus = append(
			append(
				make([]jnnkrdbdev1.DeployedObject, len(vr.Status.Deployed)-1),
				vr.Status.Deployed[:indexDeployed]...,
			),
			vr.Status.Deployed[indexDeployed+1:]...,
		)

		// re-cache the current vaultrequest
		if err := c.Get(ctx, types.NamespacedName{
			Namespace: vr.Namespace,
			Name:      vr.Name,
		}, vr, &client.GetOptions{}); err != nil {
			_log.Error(err, "FATAL error while re-caching the vaultrequest")
			return err
		}

		// update the status of the current vautlrequest
		vr.Status.Deployed = newStatus
		if err := c.Status().Update(ctx, vr); err != nil {
			_log.Error(err, "error updating status object")
			return err
		}
	}

	return nil
}
