package logging

import (
	"flag"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// default logger for vrdb services
var Log logr.Logger

// initialize the logger with a name
func InitLogger(name string) {

	var opts = zap.Options{Development: true}

	opts.BindFlags(flag.CommandLine)

	Log = zap.New(zap.UseFlagOptions(&opts)).WithName(name)
}
