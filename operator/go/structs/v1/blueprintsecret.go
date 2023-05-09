package v1

import (
	"fmt"

	"github.com/jnnkrdb/corerdb/fnc"
)

var _SecretTypeList = []string{
	"Opaque",
	"kubernetes.io/service-account-token",
	"kubernetes.io/dockercfg",
	"kubernetes.io/dockerconfigjson",
	"kubernetes.io/basic-auth",
	"kubernetes.io/ssh-auth",
	"kubernetes.io/tls",
	"bootstrap.kubernetes.io/token",
}

type BluePrintSecret struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Type      string `json:"type"`
	Immutable bool   `json:"immutable"`
}

func (bps BluePrintSecret) Validate() error {

	// checking the secret type
	if !fnc.StringInList(bps.Type, _SecretTypeList) {
		return fmt.Errorf("secrettype is not valid")
	}

	return nil
}
