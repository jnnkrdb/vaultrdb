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

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var vrdbrequestlog = logf.Log.WithName("vrdbrequest-resource")

func (r *VRDBRequest) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-jnnkrdb-de-v1-vrdbrequest,mutating=true,failurePolicy=fail,sideEffects=None,groups=jnnkrdb.de,resources=vrdbrequests,verbs=create;update,versions=v1,name=mvrdbrequest.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &VRDBRequest{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *VRDBRequest) Default() {
	vrdbrequestlog.Info("default", "name", r.Name, "namespace", r.Namespace, "required", false)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-jnnkrdb-de-v1-vrdbrequest,mutating=false,failurePolicy=fail,sideEffects=None,groups=jnnkrdb.de,resources=vrdbrequests,verbs=create;update,versions=v1,name=vvrdbrequest.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &VRDBRequest{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *VRDBRequest) ValidateCreate() error {
	vrdbrequestlog.Info("validate create", "name", r.Name, "namespace", r.Namespace, "required", true)
	return r.validate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *VRDBRequest) ValidateUpdate(old runtime.Object) error {
	vrdbrequestlog.Info("validate update", "name", r.Name, "namespace", r.Namespace, "required", true)
	return r.validate()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *VRDBRequest) ValidateDelete() error {
	vrdbrequestlog.Info("validate delete", "name", r.Name, "namespace", r.Namespace, "required", false)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
