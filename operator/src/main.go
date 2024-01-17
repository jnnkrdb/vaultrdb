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
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
	"github.com/jnnkrdb/vaultrdb/controllers"
	//+kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")

	// startup configs
	disableVRDBConfigs    bool
	disableVRDBSecrets    bool
	disableVRDBRequests   bool
	disableNamespaces     bool
	disableLeaderElection bool
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(jnnkrdbdev1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {

	flag.BoolVar(&disableVRDBConfigs, "disable-vrdb-configs", false, "Disables the VRDBConfig Reconcilation.")
	flag.BoolVar(&disableVRDBSecrets, "disable-vrdb-secrets", false, "Disables the VRDBSecret Reconcilation.")
	flag.BoolVar(&disableVRDBRequests, "disable-vrdb-requests", false, "Disables the VRDBRequest Reconcilation.")
	flag.BoolVar(&disableVRDBRequests, "disable-namespace-watcher", false, "Disables the Namespace Listener.")
	flag.BoolVar(&disableLeaderElection, "disable-leader-election", false, "If disabled, then the Controllers will not wait for a Lease to expire.")

	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                        scheme,
		MetricsBindAddress:            ":8080",
		Port:                          9443,
		HealthProbeBindAddress:        ":8081",
		LeaderElection:                !disableLeaderElection,
		LeaderElectionID:              "vrdb.jnnkrdb.de",
		LeaderElectionReleaseOnCancel: false,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if !disableVRDBConfigs {
		if err = (&controllers.VRDBConfigReconciler{
			Client: mgr.GetClient(),
			Scheme: mgr.GetScheme(),
		}).SetupWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "VRDBConfig")
			os.Exit(1)
		}
	}

	if !disableVRDBSecrets {
		if err = (&controllers.VRDBSecretReconciler{
			Client: mgr.GetClient(),
			Scheme: mgr.GetScheme(),
		}).SetupWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "VRDBSecret")
			os.Exit(1)
		}
	}

	// disabled because it is not finished
	if !disableVRDBSecrets {
		//	if err = (&controllers.VRDBRequestReconciler{
		//		Client: mgr.GetClient(),
		//		Scheme: mgr.GetScheme(),
		//	}).SetupWithManager(mgr); err != nil {
		//		setupLog.Error(err, "unable to create controller", "controller", "VRDBRequest")
		//		os.Exit(1)
		//	}
	}

	if !disableNamespaces {
		if err = (&controllers.NamespaceReconciler{
			Client: mgr.GetClient(),
			Scheme: mgr.GetScheme(),
		}).SetupWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Namespace")
			os.Exit(1)
		}
	}

	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}