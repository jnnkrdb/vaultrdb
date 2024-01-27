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
	"fmt"
	"regexp"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// VRDBSecret is the Schema for the vrdbsecrets API
type VRDBSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	NamespaceSelector VRDBNamespaceSelector `json:"namespaceSelector,omitempty"`
	Data              map[string]string     `json:"data,omitempty"`
	StringData        map[string]string     `json:"stringData,omitempty"`

	Type   v1.SecretType `json:"type,omitempty" protobuf:"bytes,3,opt,name=type,casttype=SecretType"`
	Status VRDBStatus    `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VRDBSecretList contains a list of VRDBSecret
type VRDBSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VRDBSecret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VRDBSecret{}, &VRDBSecretList{})
}

// validate function
func (r *VRDBSecret) validate() error {

	var base64rx = regexp.MustCompile(`^(?:[A-Za-z0-9+\/]{4})*(?:[A-Za-z0-9+\/]{2}==|[A-Za-z0-9+\/]{3}=|[A-Za-z0-9+\/]{4})$`)

	// validate the namespace selector
	if err := r.NamespaceSelector.Validate(); err != nil {
		vrdbsecretlog.Error(err, "error validating vrdbsecret")
		return err
	}

	// validating the base64 encoded fields
	for k, v := range r.Data {

		if len(k) == 0 {
			err := fmt.Errorf("key cannot be nil")
			vrdbsecretlog.Error(err, "error validating stringdata", "key", k)
			return err
		}

		if !base64rx.MatchString(v) {
			vrdbsecretlog.Error(fmt.Errorf("[%s] value contains forbidden character", k), "error validating base64 encoding", "key", k, "value", v)
			return fmt.Errorf("unable to validate base64 encoding of key [%s]", k)
		}

		if _, err := base64.StdEncoding.DecodeString(v); err != nil {
			vrdbsecretlog.Error(err, "error validating base64 encoding", "key", k, "value", v)
			return fmt.Errorf("unable to validate base64 encoding of key [%s]: %v", k, err)
		}
	}

	// validating the stringData fields
	for k := range r.StringData {
		if len(k) == 0 {
			err := fmt.Errorf("key cannot be nil")
			vrdbsecretlog.Error(err, "error validating stringdata", "key", k)
			return err
		}
	}

	return nil
}
