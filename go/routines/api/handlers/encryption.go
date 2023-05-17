package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jnnkrdb/corerdb/crypt"
	"github.com/jnnkrdb/gomw/handlers/httpfnc"
)

func _helperfunc(fnc func(string) (string, error), w http.ResponseWriter, r *http.Request) {
	// encrypt the data
	if r.Body == nil {
		httpfnc.AddErrorToHeaderIfAny(&w, fmt.Errorf("error: request body can't be empty"))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	} else {
		if bBytes, err := io.ReadAll(r.Body); err != nil {
			httpfnc.AddErrorToHeaderIfAny(&w, fmt.Errorf("error: reading body failed: %#v", err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		} else {
			var res string
			if res, err = fnc(string(bBytes)); err != nil {
				httpfnc.AddErrorToHeaderIfAny(&w, fmt.Errorf("error: encrypting body: %#v", err))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			} else {
				w.Write([]byte(res))
			}
		}
	}
}

func Encrypt(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		_helperfunc(crypt.EncryptWithDefault, w, r)

	case http.MethodOptions: // ------------------- default responses
		w.Header().Add("Allow", "POST,OPTIONS")
		w.Write([]byte(http.StatusText(http.StatusOK)))
	default:
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}
}

func Decrypt(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		_helperfunc(crypt.DecryptWithDefault, w, r)

	case http.MethodOptions: // ------------------- default responses
		w.Header().Add("Allow", "POST,OPTIONS")
		w.Write([]byte(http.StatusText(http.StatusOK)))
	default:
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}

}
