package v1

import (
	"encoding/base64"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/jnnkrdb/vaultrdb/svc/sqlite3"
)

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
func (dm DataMap) GetData(_log logr.Logger) (string, error) {

	_log.V(3).Info("parsing data", "uid", dm.UID, "data", dm.Data, "stringData", dm.StringData)

	var errors = [3]error{nil, nil, nil}

	if len(dm.UID) > 0 {
		if pair, err := sqlite3.SelectPairByUID(dm.UID); err != nil {
			_log.V(3).Info("couldn't receive data from sql database", "uid", dm.UID, "error.message", err)
			errors[0] = err
		} else {
			return pair.Value, nil
		}
	}

	if len(dm.Data) > 0 {
		if unenc, err := base64.StdEncoding.DecodeString(dm.Data); err != nil {
			_log.V(3).Info("couldn't decode base64 data from datafield", "data", dm.Data, "error.message", err)
			errors[1] = err
		} else {
			return string(unenc), nil
		}
	}

	if len(dm.StringData) > 0 {
		return dm.Data, nil
	} else {
		_log.V(3).Info("data field is empty", "data", dm.Data)
		errors[2] = fmt.Errorf("stringData field shouldn't be empty, when data and uid are not in use")
	}

	return "", fmt.Errorf("errors[0]: %v | errors[1]: %v | errors[2]: %v", errors[0], errors[1], errors[2])
}
