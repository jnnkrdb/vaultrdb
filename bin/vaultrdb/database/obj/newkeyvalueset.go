package obj

import (
	"encoding/json"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// the struct 'NewKeyValueSet' is a struct to summarize the information
// which is needed, to create a new KeyValueSet in the database
//
// this object is only used in http request as a json object, which fills
// the acutal KeyValueSet fields with data, to be then integrated into the
// database
type NewKeyValueSet struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

// read the object from the http request as json
func (nkvs *NewKeyValueSet) FromJSON(w http.ResponseWriter, r *http.Request) (err error) {

	if err = json.NewDecoder(r.Body).Decode(nkvs); err != nil {

		logging.Default.Info("error parsing body into struct", "nkvs", nkvs, "err", err)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return
}
