package v1

type BluePrintConfigMap struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Immutable bool   `json:"immutable"`
}

func (bpcm BluePrintConfigMap) Validate() error {

	return nil
}
