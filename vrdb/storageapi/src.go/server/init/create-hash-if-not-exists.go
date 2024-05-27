package init

import (
	"strings"
	"vrdb-storage/objects"
	"vrdb-storage/server/serviceconfigs"

	"vrdb.go/cryptography"
	"vrdb.go/logging"
)

// create the hash if neccessary for the database encryption
func CreateHASH() {

	// check if the passphrase was configured already
	// if so, then skip this step

	value, exists, err := objects.GetServiceConfig(serviceconfigs.EncryptionKey)
	switch {
	case err != nil:
		logging.Log.Info("error receiving encryption passphrase from db", "err", err)
		panic(err)
	case exists:
		logging.Log.Info("encryption passphrase already exists in db", "passphrase", strings.Repeat("*", len(value)))
	}

	// insert the passphrase into the database
	if err := objects.SetServiceConfig(serviceconfigs.EncryptionKey, cryptography.GetPassphraseFromCACertHASH()); err != nil {

		logging.Log.Info("error setting encryption passphrase in db", "err", err)
		panic(err)
	}
}
