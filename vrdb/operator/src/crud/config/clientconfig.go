package config

import (
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	// default logger for http requests
	CrudLog logr.Logger = ctrl.Log.WithName("crud")

	// default client for requests against the kubernets
	// api server, used for vrdb structs, etc.
	KClient client.Client
)
