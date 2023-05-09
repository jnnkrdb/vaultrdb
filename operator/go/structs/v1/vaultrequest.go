package v1

import (
	"context"

	"github.com/jnnkrdb/k8s/operator"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type VaultRequestSpec struct {
	Data []struct {
		VaultSetID string `json:"vaultsetid"`
		StoreKey   string `json:"storekey"`
	} `json:"data"`

	Secrets    []BluePrintSecret    `json:"secrets"`
	ConfigMaps []BluePrintConfigMap `json:"configmaps"`
}

// deepcopy
func (in *VaultRequest) DeepCopyInto(out *VaultRequest) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = VaultRequestSpec{}
}

// ----------------------------------------------------
// kubernetes dependencies
type VaultRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VaultRequestSpec `json:"spec"`
}

type VaultRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultRequest `json:"items"`
}

func (in *VaultRequest) DeepCopyObject() runtime.Object {
	out := VaultRequest{}
	in.DeepCopyInto(&out)
	return &out
}

func (in *VaultRequestList) DeepCopyObject() runtime.Object {
	out := VaultRequestList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		out.Items = make([]VaultRequest, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
	return &out
}

// ----------------------------------------------------
// helper functions

const _VR_RESOURCE string = "vaultrequests"

// requests all deployed vaultrequests and returns them as a vaultrequestslist
func GetVaultRequestList() (vrList VaultRequestList, err error) {
	err = operator.CRD().Get().Resource(_VR_RESOURCE).Do(context.TODO()).Into(&vrList)
	return
}
