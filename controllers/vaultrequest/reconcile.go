package vaultrequest

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func Reconcile(_log logr.Logger, ctx context.Context, c client.Client, vr *jnnkrdbdev1.VaultRequest) (ctrl.Result, error) {

	_log.Info("start reconciling")

	// check if the object contains the finalization flags, or has to be terminated
	if finalized, err := Finalize(_log, ctx, c, vr); err != nil || finalized {
		return ctrl.Result{Requeue: !finalized}, err
	}

	// calculating the namespaces
	match, avoid, err := vr.Spec.Namespaces.CalculateNamespaces(_log, ctx, c)
	if err != nil {
		_log.Error(err, "couldn't calculate namespaces")
		return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Minute}, err
	}

	// remove the configmap/secrets, which should not exist anymore
	if err = RemoveUnwantedObjects(_log, c, ctx, vr, avoid); err != nil {
		return ctrl.Result{Requeue: true, RequeueAfter: 2 * time.Minute}, err
	}

	// create or update the configmaps/secrets from in the match namespaces
	rec, result, err := CreateOrUpdateObjects(_log, ctx, c, vr, match)
	if err != nil || rec {
		return result, err
	}

	return ctrl.Result{Requeue: false}, nil
}
