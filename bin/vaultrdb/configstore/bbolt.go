package configstore

import (
	"os"
	"strings"

	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/configstore/buckets"
	"github.com/jnnkrdb/vaultrdb/pkg/cryptography"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
	"go.etcd.io/bbolt"
)

// path of the internal database for configs
const internalDB string = "/opt/vaultrdb/data/internal.db"

// internal config store
// gets persisted in the data directory
var DB *bbolt.DB = nil

// initialize the configstore buckets
func InitConfigStore() {

	// open the database, or create one, if not exists
	logging.SLog.Info("opening configstore", "path", internalDB)
	if _db, err := bbolt.Open(internalDB, 0600, nil); err != nil {

		logging.SLog.Info("error opening internal.db", "err", err)
		os.Exit(1)

	} else {

		DB = _db
	}

	// create the required buckets for initialization
	if DB.Update(func(tx *bbolt.Tx) error {

		logging.SLog.Info("creating buckets in internal.db if neccessary", "buckets", buckets.DefaultBuckets())

		for i := range buckets.DefaultBuckets() {

			if _, err := tx.CreateBucketIfNotExists([]byte(buckets.DefaultBuckets()[i])); err != nil {

				logging.SLog.Info("error creating bucket", "bucket", buckets.DefaultBuckets()[i], "err", err)
				return err
			}
		}
		return nil
	}) != nil {

		os.Exit(1)
	}

	// creating the encryption hash for the store.db
	if DB.Update(func(tx *bbolt.Tx) error {

		// request the encKey from the internal db, to check, if it already exists
		var encKey = string(tx.Bucket([]byte(buckets.Vault)).Get([]byte("vrdb-encryption-key")))

		// if the encKey is unset, then set it
		if encKey == "" {

			// create enckey
			encKey = cryptography.GetPassphraseFromCACertHASH()

			// save the enckey to internal db
			logging.SLog.Info("saving encKey in internaldb", "vrdb-encryption-key", strings.Repeat("*", len(encKey)))

			if err := tx.Bucket([]byte(buckets.Vault)).Put([]byte("vrdb-encryption-key"), []byte(cryptography.GetPassphraseFromCACertHASH())); err != nil {
				logging.SLog.Info("error commiting vrdb-encryption-key to bucket", "err", err)
				return err
			}
		}

		return nil
	}) != nil {

		os.Exit(1)
	}
}

// close the internal database connection
func Close() {
	if err := DB.Close(); err != nil {
		logging.SLog.Info("error closing internal config db", "err", err)
	}
}
