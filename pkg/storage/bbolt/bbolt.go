package bbolt

import (
	"fmt"

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
	var result []string = nil
	var err = bb.db.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(name []byte, b *bbolt.Bucket) error {
			result = append(result, string(name))
			return nil
		})
	})
	return result, err
}

func (bb BBoltDB) WriteBucket(bucket string) error {
	return bb.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
}

func (bb BBoltDB) DeleteBucket(bucket string) error {
	return bb.db.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(bucket))
	})
}

func (bb BBoltDB) ListKeys(bucket string) ([]string, error) {
	var result []string = nil
	var err = bb.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).ForEach(func(k, v []byte) error {
			result = append(result, string(k))
			return nil
		})
	})
	return result, err
}

func (bb BBoltDB) GetKey(bucket, key string) (string, error) {
	var result string = ""
	var err = bb.db.View(func(tx *bbolt.Tx) error {
		if v := tx.Bucket([]byte(bucket)).Get([]byte(key)); v != nil {
			result = string(v)
			return nil
		} else {
			return fmt.Errorf("the requested key does not exist or is empty")
		}
	})
	return result, err
}

func (bb BBoltDB) WriteKey(bucket, key, value string) error {
	return bb.db.Update(func(tx *bbolt.Tx) error {
		if b, err := tx.CreateBucketIfNotExists([]byte(bucket)); err != nil {
			return err
		} else {
			return b.Put([]byte(key), []byte(value))
		}
	})
}

func (bb BBoltDB) DeleteKey(bucket, key string) error {
	return bb.db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Delete([]byte(key))
	})
}

func (bb BBoltDB) Close() error {
	return bb.db.Close()
}
