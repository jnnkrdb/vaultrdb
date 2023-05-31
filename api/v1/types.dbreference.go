package v1

import (
	"fmt"

	"github.com/google/uuid"
)

// struct which contains the information about the namespace regex
type DataReference struct {

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +kubebuilder:default={}
	RefData DataField `json:"refdata,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +kubebuilder:default={}
	RefUUID UUIDField `json:"refuuid,omitempty"`
}

// ============================================== reference fields
type DataField map[string]string

type UUIDField map[string]string

func (uf UUIDField) Validate() error {

	if !(len(uf) > 0) {
		return fmt.Errorf("uuidfield is empty")
	}

	// parse through all key-value pairs, to validate the field
	for _key, _uuid := range uf {
		// check if the key field is empty
		if _key == "" {
			return fmt.Errorf("empty key not allowed")
		}

		// check if the uuid field is empty
		if _uuid == "" {
			return fmt.Errorf("empty uuid not allowed")
		}

		// validate if the uuid field is an actual uuidv4
		if _, err := uuid.Parse(_uuid); err != nil {
			return fmt.Errorf("error with key [%s]: %v", _uuid, err)
		}

		// check if the key exists in database

	}

	return nil
}
