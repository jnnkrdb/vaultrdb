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

	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// startupconfigs are configured via commandline args

func loadArgs(opts *zap.Options) {

	// configuring logs
	*opts = zap.Options{
		Development: true,
	}

	opts.BindFlags(flag.CommandLine)

	// configuring the flags from the args list
	flag.BoolVar(&disableVRDBRequests, "disable-vrdb-requests", false, "Disables the VRDBRequest Reconcilation.")
	flag.BoolVar(&disableLeaderElection, "disable-leader-election", false, "If disabled, then the Controllers will not wait for a Lease to expire.")

	flag.Parse()
}
