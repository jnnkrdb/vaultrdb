package configs

import (
	"os"
	"strings"

	"go.etcd.io/bbolt"
	"vrdb.go/cryptography"
	"vrdb.go/logging"
)

const (
	internalDB string = "/opt/vaultrdb/data/internal.db"
)

var (
	DB *bbolt.DB
)

// initialize the internal database for key/value storage of the
// vaultrdb service
//
// the internal db only contains service related configs
func InitInternalDB() {

	// open the database, or create one, if not exists
	logging.Log.Info("opening internal.db")
	if _db, err := bbolt.Open(internalDB, 0600, nil); err != nil {

		logging.Log.Info("error opening internal.db", "err", err)

		os.Exit(1)

	} else {

		DB = _db
	}

	// create the required buckets for initialization
	DB.Update(func(tx *bbolt.Tx) error {

		logging.Log.Info("creating buckets in internal.db if neccessary", "buckets", Buckets)

		for i := range Buckets {

			if _, err := tx.CreateBucketIfNotExists([]byte(Buckets[i])); err != nil {

				logging.Log.Info("error creating bucket", "bucket", Buckets[i], "err", err)
				os.Exit(1)
			}
		}
		return nil
	})

	// creating the encryption hash for the store.db
	if err := DB.Update(func(tx *bbolt.Tx) error {

		// request the encKey from the internal db, to check, if it already exists
		var encKey = string(tx.Bucket([]byte(BucketVault)).Get([]byte("vrdb-encryption-key")))

		// if the encKey is unset, then set it
		if encKey == "" {

			// create enckey
			encKey = cryptography.GetPassphraseFromCACertHASH()

			// save the enckey to internal db
			logging.Log.WithValues(
				"vrdb-encryption-key", strings.Repeat("*", len(encKey)),
			).Info("saving encKey in internaldb")

			if err := tx.Bucket([]byte(BucketVault)).Put([]byte("vrdb-encryption-key"), []byte(cryptography.GetPassphraseFromCACertHASH())); err != nil {
				logging.Log.Info("error commiting vrdb-encryption-key to bucket", "err", err)
				return err
			}
		}

		return nil
	}); err != nil {
		os.Exit(1)
	}
}
