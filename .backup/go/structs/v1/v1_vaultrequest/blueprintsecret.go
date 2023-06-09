package v1_vaultrequest

import (
	"fmt"

	"github.com/jnnkrdb/corerdb/fnc"
)

type BluePrintSecret struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Type      string `json:"type"`
	Immutable bool   `json:"immutable"`
	KeyName   string `json:"keyname"`
}

func (bps BluePrintSecret) Validate() error {

	// checking the secret type
	if !fnc.StringInList(bps.Type, []string{
		"Opaque",
		"kubernetes.io/service-account-token",
		"kubernetes.io/dockercfg",
		"kubernetes.io/dockerconfigjson",
		"kubernetes.io/basic-auth",
		"kubernetes.io/ssh-auth",
		"kubernetes.io/tls",
		"bootstrap.kubernetes.io/token",
	}) {
		return fmt.Errorf("secrettype is not valid")
	}

	return nil
}
