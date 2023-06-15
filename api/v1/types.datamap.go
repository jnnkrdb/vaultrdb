package v1

import (
	"encoding/base64"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/jnnkrdb/vaultrdb/svc/postgres"
)

// struct which contains the information about the namespace regex

type DataMap struct {

	// The PSQLID field must contain a valid id, existing in the mounted postgres database

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	PSQLID string `json:"psqlid,omitempty"`

	// The Data field must an base64 encoded version of the required value

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Data string `json:"data,omitempty"`

	// The StringData field must an unencoded version of the required value

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	StringData string `json:"stringData,omitempty"`
}

// run datafield validations, for all datafields int the datareference object
func (dm DataMap) GetData(_log logr.Logger) (string, error) {

	_log.V(5).Info("parsing data", "psqlid", dm.PSQLID, "data", dm.Data, "stringData", dm.StringData)

	var errors = [3]error{nil, nil, nil}

	if len(dm.PSQLID) > 0 && postgres.USEPOSTGRES {
		var data string
		if err := postgres.PSQL.QueryRow("SELECT data FROM public.vault WHERE psqlid=$1;", dm.PSQLID).Scan(&data); err != nil {
			_log.V(5).Error(err, "couldn't get data from postgres")
			errors[0] = err
		} else {

			return data, nil
		}
	}

	if len(dm.Data) > 0 {
		if unenc, err := base64.StdEncoding.DecodeString(dm.Data); err != nil {
			_log.V(5).Error(err, "couldn't decode base64 data from datafield", "data", dm.Data)
			errors[1] = err
		} else {

			return string(unenc), nil
		}
	}

	if len(dm.StringData) > 0 {
		return dm.Data, nil
	} else {
		errors[2] = fmt.Errorf("stringData field shouldn't be empty, when data and psqlid are not in use")
	}

	return "", fmt.Errorf("errors[0]: %v | errors[1]: %v | errors[2]: %v", errors[0], errors[1], errors[2])
}
