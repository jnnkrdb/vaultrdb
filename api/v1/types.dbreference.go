package v1

import (
	"fmt"
	"sync"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
)

// struct which contains the information about the namespace regex
type DataReference struct {

	// The Data field must contain all values, as a base64 encoded version
	//
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +kubebuilder:default={}
	Data EncodedDataField `json:"data,omitempty"`

	// The StringData field must contain all values, as an unencoded version
	//
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +kubebuilder:default={}
	StringData RawDataField `json:"stringData,omitempty"`

	// The UUIDs field must contain all values, as an uuid v4, existing in the mounted database
	//
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +kubebuilder:default={}
	UUIDs UUIDField `json:"uuids,omitempty"`
}

// run datafield validations, for all datafields int the datareference object
func (dr DataReference) RunValidations(_log logr.Logger) error {

	type dataField interface {
		Validate(logr.Logger) error
	}

	var (
		err    error = nil
		eMutex sync.Mutex
		wg     sync.WaitGroup
	)

	wg.Add(3)

	// create 3 goroutines to validate the datafields, since we have 3 datafields in the objects
	// we add 3 goroutines
	for _, df := range []dataField{dr.UUIDs, dr.StringData, dr.Data} {

		_log.Info("validating datafield", "datafield", df)

		go func(datafield dataField) {

			if e := datafield.Validate(_log); e != nil {
				_log.Info("validating datafield failed", "err", e)
				e = fmt.Errorf("error validating datafield: %v - ", e)
				// locking mutex, to write err output to local error
				eMutex.Lock()
				err = fmt.Errorf("%s%s", err.Error(), e.Error())
				eMutex.Unlock()
			}

			wg.Done()
		}(df)
	}

	wg.Wait()

	return err
}

// ============================================== reference fields - EncodedDataField

type EncodedDataField map[string]string

func (edf EncodedDataField) Validate(_log logr.Logger) error {
	// parse through all key-value pairs, to validate the field
	for _key, _encdata := range edf {

		_log = _log.V(5).WithValues("_key", _key)

		_log.Info("validating key/value pair", "_encdata", _encdata)
	}

	return nil
}

// ============================================== reference fields - DataField

type RawDataField map[string]string

func (df RawDataField) Validate(_log logr.Logger) error {
	// parse through all key-value pairs, to validate the field
	for _key, _rawdata := range df {

		_log = _log.V(5).WithValues("_key", _key)

		_log.Info("validating key/value pair", "_rawdata", _rawdata)

	}

	return nil
}

// ============================================== reference fields - UUIDField

type UUIDField map[string]string

func (uf UUIDField) Validate(_log logr.Logger) error {
	// parse through all key-value pairs, to validate the field
	for _key, _uuid := range uf {

		_log = _log.V(5).WithValues("_key", _key)

		_log.Info("validating key/value pair", "_uuid", _uuid)

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
