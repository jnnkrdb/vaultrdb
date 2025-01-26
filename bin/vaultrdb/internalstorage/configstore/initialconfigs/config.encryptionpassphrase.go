package initialconfigs

import (
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/configstore"
	"github.com/jnnkrdb/vaultrdb/pkg/cryptography"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

func calculateencryptionpassphrase() error {

	logging.SLog.Info("calculating the encryption passphrase from cacert")

	var encKey string = ""
	if encKey = configstore.GetConfig(configstore.EncryptionPassphrase); encKey != "" {
		logging.SLog.Info("encryption passphrase is already configured")
		return nil
	}

	// set encryption passphrase
	return configstore.SetConfig(configstore.EncryptionPassphrase,
		cryptography.GetPassphraseFromCACertHASH())
}
