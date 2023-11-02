package logging

import (
	"flag"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	IsDevelopment bool
)

// set the logging configs for the operator
func SetLogging() {

	opts := zap.Options{
		Development: IsDevelopment,
	}

	opts.BindFlags(flag.CommandLine)

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))
}
