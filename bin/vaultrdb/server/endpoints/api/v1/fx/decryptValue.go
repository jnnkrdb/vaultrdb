package api_v1_fx

import (
	"net/http"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/configstore"
	"github.com/jnnkrdb/vaultrdb/pkg/cryptography"
	"github.com/jnnkrdb/vaultrdb/pkg/http/helpers"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// receive a specific value via POST and send an encrypted version of it
// as a response
func DecryptValue(w http.ResponseWriter, r *http.Request) {

	var body helpers.StringValueJson
	if body.Receive(w, r.Body) != nil {
		return
	}

	res, err := cryptography.Decrypt(configstore.GetConfig(configstore.EncryptionPassphrase), body.Value)
	if err != nil {
		logging.SLog.Warn("couldn't decrypt body",
			"response-code", http.StatusBadRequest,
			"error", err.Error(),
		)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	helpers.StringValueJson{Value: res}.Send(w)
}
