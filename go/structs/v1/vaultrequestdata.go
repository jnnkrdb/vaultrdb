package v1

import "fmt"

type VaultRequestData struct {
	VaultSetID string               `json:"vaultsetid"`
	Secrets    []BluePrintSecret    `json:"secrets"`
	ConfigMaps []BluePrintConfigMap `json:"configmaps"`
}

func (vrd VaultRequestData) Validate() error {

	if _, e := SelectByID(vrd.VaultSetID); e != nil {
		return fmt.Errorf("error requesting vaultset[%s] from database: %v", vrd.VaultSetID, e)
	}

	return nil
}
