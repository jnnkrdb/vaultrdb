package api_v1_fx

import (
	"net/http"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/fx/crypt"
	"github.com/jnnkrdb/vaultrdb/pkg/http/helpers"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// receive a specific value via POST and send an encrypted version of it
// as a response
func EncryptValue(w http.ResponseWriter, r *http.Request) {

	var body helpers.StringValueJson
	if body.Receive(w, r) != nil {
		return
	}

	res, err := crypt.Decrypt(body.Value)
	if err != nil {
		logging.SLog.Info("couldn't decrypt body", "response-code", http.StatusBadRequest)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	helpers.StringValueJson{Value: res}.Send(w)
}
