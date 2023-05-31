package v1

import (
	"github.com/jnnkrdb/vaultrdb/structs/v1/v1_vaultrequest"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	GroupName    = "vaultrdb.jnnkrdb.de"
	GroupVersion = "v1"
)

var (
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: GroupVersion}
	SchemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme        = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&v1_vaultrequest.VaultRequestList{},
		&v1_vaultrequest.VaultRequest{},
	)

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
