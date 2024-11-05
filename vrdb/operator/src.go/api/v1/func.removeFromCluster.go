package v1

import (
	"context"
	"fmt"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// remove the unwanted objects from the cluster
func RemoveUnwantedObjectFromCluster(
	to client.Object,
	ctx context.Context,
	c client.Client,
	s *VRDBStatus,
	toAvoid []string,
	o client.Object) (ctrl.Result, error) {

	var _log = log.FromContext(ctx)

	vrdbObjectNamespacedName, ok := ctx.Value(types.NamespacedName{}).(types.NamespacedName)
	if !ok {
		return ctrl.Result{RequeueAfter: 1 * time.Minute}, fmt.Errorf("could not parse context into types.NamespacedName: %v", ctx.Value(types.NamespacedName{}))
	}

	for _, namespace := range s.Namespaces {
		var __log = _log.WithValues("namespace", namespace)

		if !strings.Contains(strings.Join(toAvoid, ";"), namespace) {
			continue
		}

		__log.Info("removing object from namespace, if exists")
		vrdbObjectNamespacedName.Namespace = namespace

		// receive obj information
		if err := c.Get(ctx, vrdbObjectNamespacedName, o, &client.GetOptions{}); err != nil {
			if errors.IsNotFound(err) {
				_log.Info("object does not exist in namespace", "namespace", namespace)
				continue
			}
			_log.Error(err, "error receiving object information", "namespace", namespace)
			return ctrl.Result{RequeueAfter: 1 * time.Minute}, err
		}

		// remove requested obj from cluster
		if err := c.Delete(ctx, o, &client.DeleteOptions{}); err != nil {
			_log.Error(err, "error removing object from cluster", "namespace", namespace)
			return ctrl.Result{RequeueAfter: 1 * time.Minute}, err
		}

		// updating object status
		s.RemoveNamespace(namespace)
		if err := Update(ctx, c, true, to); err != nil {
			return ctrl.Result{RequeueAfter: 1 * time.Minute}, err
		}
	}

	return ctrl.Result{}, nil
}
