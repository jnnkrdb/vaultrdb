package storedb

import (
	"fmt"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
	"go.etcd.io/bbolt"
)

// ######################################################################################################### READ

// read key/value under path
func Read(path, key string) (string, error) {

	// calculating the path
	var bucketList = calculateBucketsFromPath(path)

	// create the transaction
	tx, err := DB.Begin(false)
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	var (
		currentBucket     *bbolt.Bucket = nil
		currentBucketPath string        = bucketList[0]
	)

	if currentBucket = tx.Bucket([]byte(bucketList[0])); currentBucket == nil {
		return "", fmt.Errorf("couldn't find bucket [%s]", currentBucketPath)
	}

	for _, item := range bucketList[1:] {
		currentBucketPath = fmt.Sprintf("%s/%s", currentBucketPath, item)

		// find the bucket
		if currentBucket = currentBucket.Bucket([]byte(item)); currentBucket == nil {
			return "", fmt.Errorf("couldn't find bucket [%s]", currentBucketPath)
		}
	}

	// read the requested key from the bucket
	var result = currentBucket.Get([]byte(key))
	if result == nil {
		return "", fmt.Errorf("key [%s] does not exist", key)
	}

	return string(result), nil
}

// ######################################################################################################### WRITE

// create key/value under path
func Write(path, key, value string) error {

	// calculating the path
	var bucketList = calculateBucketsFromPath(path)

	// create the transaction
	writeMutex.Lock()
	defer writeMutex.Unlock()

	tx, err := DB.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var (
		currentBucket     *bbolt.Bucket = nil
		currentBucketPath string        = bucketList[0]
	)

	if currentBucket, err = tx.CreateBucketIfNotExists([]byte(bucketList[0])); err != nil {
		return fmt.Errorf("couldn't create bucket [%s]: %s", currentBucketPath, err.Error())
	}

	for _, item := range bucketList[1:] {
		currentBucketPath = fmt.Sprintf("%s/%s", currentBucketPath, item)

		// create the bucket if not exist and change to new bucket
		if currentBucket, err = currentBucket.CreateBucketIfNotExists([]byte(item)); err != nil {
			return fmt.Errorf("couldn't create bucket [%s]: %s", currentBucketPath, err.Error())
		}
	}

	// write the key/value pair into the requested bucket
	if err = currentBucket.Put([]byte(key), []byte(value)); err != nil {
		return fmt.Errorf("couldn't write key/value into bucket [%s:%s]: %s", currentBucketPath, key, err.Error())
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

// ######################################################################################################### DELETE

const (
	TypeBucket = "bucket"
	TypeKey    = "key"
)

// remove a key from a bucket
func DeleteKey(path, key string) error {
	return deleteObject(TypeKey, path, key)
}

// remove a bucket from db
func DeleteBucket(path, bucket string) error {
	return deleteObject(TypeBucket, path, bucket)
}

// delete any object from db
func deleteObject(t, path, objectID string) error {

	// calculating the path
	var bucketList = calculateBucketsFromPath(path)

	// create the transaction
	writeMutex.Lock()
	defer writeMutex.Unlock()

	tx, err := DB.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var (
		currentBucket     *bbolt.Bucket = nil
		currentBucketPath string        = ""
	)

	if path == "" &&
		objectID != "" &&
		t == TypeBucket {

		// remove bucket
		if err := tx.DeleteBucket([]byte(objectID)); err != nil {
			logging.SLog.Error("couldn't delete key from bucket", "type", t, "path", path, "objectID", objectID, "error", err.Error())
			return err
		}

	} else {

		currentBucketPath = bucketList[0]

		if currentBucket = tx.Bucket([]byte(bucketList[0])); currentBucket == nil {
			logging.SLog.Info("couldn't find bucket, skipping deletion with no error", "path", currentBucketPath, "requestedPath", path)
			return nil
		}

		for _, item := range bucketList[1:] {
			currentBucketPath = fmt.Sprintf("%s/%s", currentBucketPath, item)

			// find the bucket
			if currentBucket = currentBucket.Bucket([]byte(item)); currentBucket == nil {
				logging.SLog.Info("couldn't find bucket, skipping deletion with no error", "path", currentBucketPath, "requestedPath", path)
				return nil
			}
		}

		switch t {
		case TypeKey:
			// remove the key from the bucket, if exists
			if err := currentBucket.Delete([]byte(objectID)); err != nil {
				logging.SLog.Error("couldn't delete key from bucket", "type", t, "path", path, "objectID", objectID, "error", err.Error())
				return err
			}
		case TypeBucket:
			// remove bucket
			if err := currentBucket.DeleteBucket([]byte(objectID)); err != nil {
				logging.SLog.Error("couldn't delete key from bucket", "type", t, "path", path, "objectID", objectID, "error", err.Error())
				return err
			}
		default:
			return fmt.Errorf("type not supported")
		}
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
