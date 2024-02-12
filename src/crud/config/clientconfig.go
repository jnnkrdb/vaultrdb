package config

import (
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var KClient client.Client

var CrudLog = ctrl.Log.WithName("crud")
