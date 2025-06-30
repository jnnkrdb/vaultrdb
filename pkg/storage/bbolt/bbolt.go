package bbolt

import (
	"go.etcd.io/bbolt"
)

type BBoltDB struct {
	db      *bbolt.DB
	options *bbolt.Options
}

func NewBBoltDB(path string, options *bbolt.Options) (BBoltDB, error) {
	var (
		db  BBoltDB
		err error = nil
	)
	db.options = options
	db.db, err = bbolt.Open(path, 0600, db.options)
	return db, err
}

// -------------------------------------------------------------- required functions

func (bb BBoltDB) ListBuckets() ([]string, error) {

	// create the transaction
	var (
		result []string = nil
		err    error    = nil
	)

	return result, err
}

func (bb BBoltDB) DeleteBucket(bucket string) error {
	return nil
}

func (bb BBoltDB) ListKeys(bucket string) ([]string, error) {

	// create the transaction
	var (
		result []string = nil
		err    error    = nil
	)

	return result, err
}

func (bb BBoltDB) GetKey(bucket, key string) (string, error)
func (bb BBoltDB) Write(bucket, key, value string) error
func (bb BBoltDB) DeleteKey(bucket, key string) error

func (bb BBoltDB) Close() error {
	return bb.db.Close()
}
