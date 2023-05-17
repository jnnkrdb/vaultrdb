package v1

type VaultRequestData struct {
	VaultSetID string               `json:"vaultsetid"`
	Secrets    []BluePrintSecret    `json:"secrets"`
	ConfigMaps []BluePrintConfigMap `json:"configmaps"`
}

func (vrd VaultRequestData) Validate() error {

	return nil
}
