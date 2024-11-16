package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/libs/logging"
)

type StringValueJson struct {
	Value string `json:"value"`
}

// receive a struct from a http request
func (svj *StringValueJson) Receive(w http.ResponseWriter, r *http.Request) error {

	if err := json.NewDecoder(r.Body).Decode(svj); err != nil {

		logging.Log.WithValues(
			"response-code", http.StatusBadRequest,
			"string-value-json", *svj,
		).Info("couldn't parse body into struct")

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return err
	}

	return nil
}

// send the struct via http
func (svj StringValueJson) Send(w http.ResponseWriter) {

	if err := json.NewEncoder(w).Encode(svj); err != nil {

		logging.Log.WithValues(
			"response-code", http.StatusBadRequest,
		).Info("couldn't parse to json")

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
