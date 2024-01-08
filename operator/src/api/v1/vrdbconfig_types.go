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

// VRDBConfigSpec defines the desired state of VRDBConfig
type VRDBConfigSpec struct {

	// +kubebuilder:default={}
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	MustAvoidRegex []string `json:"mustavoidregex"`

	// +kubebuilder:default={}
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	MustMatchRegex []string `json:"mustmatchregex"`

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Data map[string]string `json:"data,omitempty"`
}

func (specs VRDBConfigSpec) GetAvoidingRegexList() []string {
	return specs.MustAvoidRegex
}

func (specs VRDBConfigSpec) GetMatchingRegexList() []string {
	return specs.MustMatchRegex
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// VRDBConfig is the Schema for the vrdbconfigs API
type VRDBConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VRDBConfigSpec `json:"spec,omitempty"`
	Status VRDBStatus     `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VRDBConfigList contains a list of VRDBConfig
type VRDBConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VRDBConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VRDBConfig{}, &VRDBConfigList{})
}
