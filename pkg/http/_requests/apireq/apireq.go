package apireq

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

var possibleMethods = []string{
	"list",
	"find",
	"write",
	"delete",
}

type APIRequest struct {
	Method string
	Type   string
}

func (ar *APIRequest) FromBody(body io.ReadCloser) error {
	if err := json.NewDecoder(body).Decode(ar); err != nil {

		logging.SLog.Warn("couldn't receive object from json-body",
			"error", err.Error(),
		)

		return err

	}

	// validate the apireq method
	var validMethod bool = false
	for _, m := range possibleMethods {
		if ar.Method == m {
			validMethod = true
			break
		}
	}

	if !validMethod {
		return fmt.Errorf("method of the apireq is not valid")
	}

	return nil
}

func (ar APIRequest) Send(w http.ResponseWriter) error {
	if err := json.NewEncoder(w).Encode(ar); err != nil {

		logging.SLog.Warn("couldn't parse object to json",
			"response-code", http.StatusBadRequest,
			"error", err.Error(),
		)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	return nil
}
