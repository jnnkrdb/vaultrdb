package storedb

import (
	"fmt"

	"go.etcd.io/bbolt"
)

// ######################################################################################################### READ

func ReadBuckets(path string) ([]string, error) {
	return db.ReadBuckets(path)
}

func (ldb *LocalDB) ReadBuckets(path string) ([]string, error) {

	// calculating the path
	var bucketList = calculateBucketsFromPath(path)

	// create the transaction
	var (
		result []string = nil
		err    error    = nil
	)

	// if bucketList is nil, then the toplevel buckets are requested
	if bucketList == nil {
		err = ldb.db.View(func(tx *bbolt.Tx) error {
			return tx.ForEach(func(name []byte, b *bbolt.Bucket) error {
				result = append(result, string(name))
				return nil
			})
		})
		return result, err
	}

	// calculate lower level bucket
	err = ldb.db.View(func(tx *bbolt.Tx) error {

		// read the correct bucket
		currentBucket, _, e := readBuckets(tx, bucketList)
		if e != nil {
			return e
		}

		// cursor over the whole bucket
		return currentBucket.ForEachBucket(func(k []byte) error {
			result = append(result, string(k))
			return nil
		})
	})

	return result, err
}

func ReadKeys(path string) ([]string, error) {
	return db.ReadKeys(path)
}

func (ldb *LocalDB) ReadKeys(path string) ([]string, error) {

	// calculating the path
	var bucketList = calculateBucketsFromPath(path)

	// create the transaction
	var (
		result []string = nil
		err    error    = nil
	)

	err = ldb.db.View(func(tx *bbolt.Tx) error {

		// read the correct bucket
		currentBucket, _, e := readBuckets(tx, bucketList)
		if e != nil {
			return e
		}

		// cursor over the whole bucket
		c := currentBucket.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			result = append(result, string(k))
		}

		return nil
	})

	return result, err
}

func ReadKey(path, key string) (string, error) {
	return db.ReadKey(path, key)
}

func (ldb *LocalDB) ReadKey(path, key string) (string, error) {

	// calculating the path
	var bucketList = calculateBucketsFromPath(path)

	// create the transaction
	var (
		result []byte = nil
		err    error  = nil
	)

	err = ldb.db.View(func(tx *bbolt.Tx) error {

		// read the correct bucket
		currentBucket, currentBucketPath, e := readBuckets(tx, bucketList)
		if e != nil {
			return e
		}

		result = currentBucket.Get([]byte(key))
		if result == nil {
			return fmt.Errorf("key [%s:%s] does not exist", currentBucketPath, key)
		}

		return nil
	})

	return string(result), err
}

// ######################################################################################################### WRITE

func WriteKey(path, key, value string) error {
	return db.WriteKey(path, key, value)
}

func (ldb *LocalDB) WriteKey(path, key, value string) error {

	// calculating the path
	var bucketList = calculateBucketsFromPath(path)

	// create the transaction
	ldb.writeMutex.Lock()
	defer ldb.writeMutex.Unlock()

	return ldb.db.Update(func(tx *bbolt.Tx) error {

		// create the buckets if not exist
		currentBucket, currentBucketPath, e := createBuckets(tx, bucketList)
		if e != nil {
			return e
		}

		// write the key/value pair into the requested bucket
		if e = currentBucket.Put([]byte(key), []byte(value)); e != nil {
			return fmt.Errorf("couldn't write key/value into bucket [%s:%s]: %s", currentBucketPath, key, e.Error())
		}

		return nil
	})
}

func WriteBucket(path string) error {
	return db.WriteBucket(path)
}

func (ldb *LocalDB) WriteBucket(path string) error {

	// calculating the path
	var bucketList = calculateBucketsFromPath(path)

	// create the transaction
	ldb.writeMutex.Lock()
	defer ldb.writeMutex.Unlock()

	return ldb.db.Update(func(tx *bbolt.Tx) error {
		_, _, err := createBuckets(tx, bucketList)
		return err
	})
}

// ######################################################################################################### DELETE

const (
	TypeBucket = "bucket"
	TypeKey    = "key"
)

func DeleteKey(path, key string) error {
	return db.deleteObject(TypeKey, path, key)
}

func DeleteBucket(path, bucket string) error {
	return db.deleteObject(TypeBucket, path, bucket)
}

func (ldb *LocalDB) DeleteKey(path, key string) error {
	return ldb.deleteObject(TypeKey, path, key)
}

func (ldb *LocalDB) DeleteBucket(path, bucket string) error {
	return ldb.deleteObject(TypeBucket, path, bucket)
}

func (ldb *LocalDB) deleteObject(t, path, objectID string) error {

	// calculating the path
	var bucketList = calculateBucketsFromPath(path)

	// create the transaction
	ldb.writeMutex.Lock()
	defer ldb.writeMutex.Unlock()

	return ldb.db.Update(func(tx *bbolt.Tx) error {

		// if the object is a first level bucket
		if path == "" && objectID != "" && t == TypeBucket {
			return tx.DeleteBucket([]byte(objectID))
		}

		// read the correct bucket
		currentBucket, _, e := readBuckets(tx, bucketList)
		if e != nil {
			return e
		}

		// remove the object

		switch t {
		case TypeKey: // remove the key from the bucket, if exists
			return currentBucket.Delete([]byte(objectID))

		case TypeBucket: // remove bucket, if exists
			return currentBucket.DeleteBucket([]byte(objectID))

		default: // do not support other types
			return fmt.Errorf("type not supported")
		}
	})
}
