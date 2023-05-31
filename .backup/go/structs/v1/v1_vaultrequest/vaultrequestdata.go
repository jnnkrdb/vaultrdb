package v1_vaultrequest

import (
	"fmt"

	"github.com/jnnkrdb/vaultrdb/structs/v1/v1_vaultset"
)

type VaultRequestData struct {
	VaultSetID string               `json:"vaultsetid"`
	Secrets    []BluePrintSecret    `json:"secrets"`
	ConfigMaps []BluePrintConfigMap `json:"configmaps"`
}

func (vrd VaultRequestData) Validate() error {

	if _, e := v1_vaultset.SelectByID(vrd.VaultSetID); e != nil {
		return fmt.Errorf("error requesting vaultset[%s] from database: %v", vrd.VaultSetID, e)
	}

	return nil
}
