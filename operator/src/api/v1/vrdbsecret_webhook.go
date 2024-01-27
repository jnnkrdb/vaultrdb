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
	"encoding/base64"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var vrdbsecretlog = logf.Log.WithName("vrdbsecret-resource")

func (r *VRDBSecret) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-jnnkrdb-de-v1-vrdbsecret,mutating=true,failurePolicy=fail,sideEffects=None,groups=jnnkrdb.de,resources=vrdbsecrets,verbs=create;update,versions=v1,name=mvrdbsecret.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &VRDBSecret{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *VRDBSecret) Default() {
	vrdbsecretlog.Info("default", "name", r.Name, "namespace", r.Namespace, "required", true)

	// update the stringData into base64 data
	for k, v := range r.StringData {

		// skip the key:value pair, if there is already a same key in data-field
		if _, ok := r.Data[k]; !ok {

			r.Data[k] = base64.StdEncoding.EncodeToString([]byte(v))
		}
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-jnnkrdb-de-v1-vrdbsecret,mutating=false,failurePolicy=fail,sideEffects=None,groups=jnnkrdb.de,resources=vrdbsecrets,verbs=create;update,versions=v1,name=vvrdbsecret.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &VRDBSecret{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *VRDBSecret) ValidateCreate() error {
	vrdbsecretlog.Info("validate create", "name", r.Name, "namespace", r.Namespace, "required", true)
	return r.validate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *VRDBSecret) ValidateUpdate(old runtime.Object) error {
	vrdbsecretlog.Info("validate update", "name", r.Name, "namespace", r.Namespace, "required", true)
	return r.validate()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *VRDBSecret) ValidateDelete() error {
	vrdbsecretlog.Info("validate delete", "name", r.Name, "namespace", r.Namespace, "required", false)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
