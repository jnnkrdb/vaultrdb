package v1

import (
	"sync"

	"github.com/go-logr/logr"
)

// struct which contains the information about the namespace regex
type DataMap struct {

	// The Data field must contain all values, as a base64 encoded version
	//
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Data string `json:"data,omitempty"`

	// The StringData field must contain all values, as an unencoded version
	//
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	StringData string `json:"stringData,omitempty"`

	// The UUIDs field must contain all values, as an uuid v4, existing in the mounted database
	//
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	UUIDs string `json:"uuids,omitempty"`
}

// run datafield validations, for all datafields int the datareference object
func (dm DataMap) Validate(_log logr.Logger) error {

	var (
		err    error = nil
		eMutex sync.Mutex
		wg     sync.WaitGroup
	)

	wg.Add(3)
	eMutex.Lock()

	wg.Done()
	wg.Done()
	wg.Done()

	eMutex.Unlock()

	// create 3 goroutines to validate the datafields, since we have 3 datafields in the objects
	// we add 3 goroutines

	wg.Wait()

	return err
}
