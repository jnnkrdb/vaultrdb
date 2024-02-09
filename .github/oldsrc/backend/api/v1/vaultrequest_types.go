/*
Copyright 2023.

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

	"github.com/jnnkrdb/vaultrdb/svc/sqlite3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VaultRequestSpec defines the desired state of VaultRequest
type VaultRequestSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	DataMap map[string]DataMap `json:"datamap"`

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Namespaces NamespacesRegex `json:"namespaces"`
}

// struct which contains the information about the namespace regex

type DataMap struct {

	// The UID field must contain a valid string, existing in the mounted sql database

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	UID string `json:"uid,omitempty"`

	// The Data field must an base64 encoded version of the required value

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Data string `json:"data,omitempty"`

	// The StringData field must an unencoded version of the required value

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	StringData string `json:"stringData,omitempty"`
}

// run datafield validations, for all datafields int the datareference object
func (dm DataMap) GetData() (string, error) {
	if len(dm.UID) > 0 {
		if pair, err := sqlite3.SelectPairByUID(dm.UID); err != nil {
			return "", err
		} else {
			return pair.Value, nil
		}
	}

	if len(dm.Data) > 0 {
		if unenc, err := base64.StdEncoding.DecodeString(dm.Data); err != nil {
			return "", err
		} else {
			return string(unenc), nil
		}
	}

	if len(dm.StringData) > 0 {
		return dm.Data, nil
	} else {
		return "", fmt.Errorf("no valid data found")
	}
}

// VaultRequestStatus defines the observed state of VaultRequest
type VaultRequestStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +operator-sdk:csv:customresourcedefinitions:type=status

	Deployed []DeployedObject `json:"deployed,omitempty"`
}

// +operator-sdk:csv:customresourcedefinitions:type=status
type DeployedObject struct {
	Kind      string `json:"kind"`
	Namespace string `json:"namespace"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// VaultRequest is the Schema for the vaultrequests API
// +kubebuilder:subresource:status
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=vaultrequests,shortName=vr;vrs
type VaultRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VaultRequestSpec   `json:"spec,omitempty"`
	Status VaultRequestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VaultRequestList contains a list of VaultRequest
type VaultRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VaultRequest{}, &VaultRequestList{})
}
