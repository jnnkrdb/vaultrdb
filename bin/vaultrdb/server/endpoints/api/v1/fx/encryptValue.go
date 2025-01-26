package api_v1_fx

import (
	"net/http"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/configstore"
	"github.com/jnnkrdb/vaultrdb/pkg/cryptography"
	"github.com/jnnkrdb/vaultrdb/pkg/http/helpers"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// receive a specific value via POST and send a decrypted version of it
// as a response
func DecryptValue(w http.ResponseWriter, r *http.Request) {

	var body helpers.StringValueJson
	if body.Receive(w, r) != nil {
		return
	}

	res, err := cryptography.Encrypt(configstore.GetConfig(configstore.EncryptionPassphrase), body.Value)
	if err != nil {
		logging.SLog.Info("couldn't encrypt body", "response-code", http.StatusBadRequest)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	helpers.StringValueJson{Value: res}.Send(w)
}
