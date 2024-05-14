package v1

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func GetObjectFromCluster(ctx context.Context, c client.Client, req ctrl.Request, obj client.Object, opts ...client.GetOption) (ctrl.Result, error) {
	var l = log.FromContext(ctx)

	// checking the requested object
	if err := c.Get(ctx, req.NamespacedName, obj, opts...); err != nil {
		l.Error(err, "error reconciling object")
		if !errors.IsNotFound(err) {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}
