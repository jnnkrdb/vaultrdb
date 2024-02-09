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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// VRDBRequest is the Schema for the vrdbrequests API
type VRDBRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	NamespaceSelector VRDBNamespaceSelector `json:"namespaceSelector,omitempty"`
	Data              map[string]string     `json:"data,omitempty"`

	Status VRDBStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VRDBRequestList contains a list of VRDBRequest
type VRDBRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VRDBRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VRDBRequest{}, &VRDBRequestList{})
}

// validate function
func (r *VRDBRequest) validate() error {

	return nil
}
