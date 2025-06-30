package helpers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

type StringValueJson struct {
	Value  string   `json:"value"`
	Values []string `json:"values"`
}

// receive a struct from a http request
func (svj *StringValueJson) Receive(w http.ResponseWriter, body io.ReadCloser) error {

	if err := json.NewDecoder(body).Decode(svj); err != nil {

		if w != nil {

			logging.Default.Info("couldn't parse body into struct",
				"response-code", http.StatusBadRequest,
				"string-value-json", *svj,
			)

			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}

		return err
	}

	return nil
}

// send the struct via http
func (svj StringValueJson) Send(w http.ResponseWriter) {

	if err := json.NewEncoder(w).Encode(svj); err != nil {

		logging.Default.Info("couldn't parse to json",
			"response-code", http.StatusBadRequest,
		)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
