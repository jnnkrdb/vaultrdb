/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/int/server"
	"github.com/jnnkrdb/vaultrdb/libs/logging"
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	//+kubebuilder:scaffold:imports
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	//+kubebuilder:scaffold:scheme
}

func main() {
	var enableLeaderElection bool
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")

	logging.InitLogger("vaultrdb")

	ctrl.SetLogger(logging.Log)

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     "", //":8080",
		Port:                   9443,
		HealthProbeBindAddress: "", //":8081",
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "vaultrdb.jnnkrdb.de",
		// LeaderElectionReleaseOnCancel defines if the leader should step down voluntarily
		// when the Manager ends. This requires the binary to immediately end when the
		// Manager is stopped, otherwise, this setting is unsafe. Setting this significantly
		// speeds up voluntary leader transitions as the new leader don't have to wait
		// LeaseDuration time first.
		//
		// In the default scaffold provided, the program ends immediately after
		// the manager stops, so would be fine to enable this option. However,
		// if you are doing or is intended to do any operation such as perform cleanups
		// after the manager stops then its usage might be unsafe.
		// LeaderElectionReleaseOnCancel: true,
	})
	if err != nil {
		logging.Log.Error(err, "unable to start vaultrdb")
		os.Exit(1)
	}

	//+kubebuilder:scaffold:builder

	/*
		removing healthz checks first, because they are
		implemented in the http backend server
		TODO: maybe implement different stages for storagebackend, ui and operator with flags

		if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
			logging.Log.Error(err, "unable to set up health check")
			os.Exit(1)
		}
		if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
			logging.Log.Error(err, "unable to set up ready check")
			os.Exit(1)
		}
	*/

	logging.Log.Info("starting vaultrdb http backend async")
	go func() {
		var listFuncs = []func(*mux.Router){}

		if err := server.StartHTTP(listFuncs...); err != nil {
			logging.Log.Error(err, "error keeping up the http server")
			os.Exit(1)
		}
	}()

	logging.Log.Info("starting vaultrdb operator")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		logging.Log.Error(err, "problem running vaultrdb")
		server.StopHTTP()
		os.Exit(1)
	}
}
