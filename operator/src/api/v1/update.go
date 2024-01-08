package v1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// default update function for the object, with an retry configuration
func Update(ctx context.Context, c client.Client, statusUpdate bool, obj client.Object) error {
	var _log = log.FromContext(ctx).WithValues("statusUpdate", statusUpdate)

	_log.Info("updating object")

	// configuring the default retry option for the upüdate function, since there is
	// the possibility, that updated objects, are already updated, since the last time they
	// were cached in the function
	if retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {

		if v, ok := ctx.Value(VRDBKey{}).(types.NamespacedName); ok {
			if err := c.Get(ctx, v, obj); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("could not parse context into types.NamespacedName: %v", ctx.Value(VRDBKey{}))
		}

		if statusUpdate {
			if err := c.Status().Update(ctx, obj); err != nil {
				return err
			}
		} else {
			if err := c.Update(ctx, obj); err != nil {
				return err
			}
		}

		return nil
	}); retryErr != nil {
		_log.Error(retryErr, "failed to update object")
		return retryErr
	}

	return nil
}
