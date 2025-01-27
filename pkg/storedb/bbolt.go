package storedb

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"go.etcd.io/bbolt"
)

var db *LocalDB = &LocalDB{}

type LocalDB struct {
	db         *bbolt.DB
	writeMutex sync.Mutex
	options    *bbolt.Options
}

// create a new localdb instance
func (localdb *LocalDB) OpenDB(path string, options *bbolt.Options) (err error) {

	if localdb.db == nil {
		localdb.db, err = bbolt.Open(path, 0600, options)

		if localdb.options == nil {
			localdb.options = options
		}
	}

	return
}

// this opens a bolt file, if not existent, creates one and
// serves the db as the internal key/value store
func OpenDB(path string, timeout time.Duration, readonly bool) (err error) {
	return db.OpenDB(path, &bbolt.Options{
		Timeout:  timeout,
		ReadOnly: readonly,
	})
}

// close a specific instance
func (localdb *LocalDB) CloseDB() error {
	if err := localdb.db.Close(); err != nil {
		return err
	}

	localdb.db = nil

	return nil
}

// close the db if opened
func CloseDB() error {
	return db.CloseDB()
}

// calculate the correct bucket walking path
func calculateBucketsFromPath(path string) []string {

	// if the pathes first char is not @ then the path is wrong
	// @ marks the root
	if 

	// replacing spaces with ""
	path = strings.ReplaceAll(path, " ", "")

	// replace // with /, as long as there are //
	for replace_doubleslash := true; replace_doubleslash; replace_doubleslash = strings.Contains(path, "..") {
		path = strings.ReplaceAll(path, "..", ".")
	}

	// if path is empty or single / return nil
	if path == "" || path == "." {
		return nil
	}

	// removing prefix or suffix /
	path, _ = strings.CutPrefix(path, ".")
	path, _ = strings.CutSuffix(path, ".")

	return strings.Split(path, ".")
}

// read buckets
func readBuckets(tx *bbolt.Tx, bucketList []string) (*bbolt.Bucket, string, error) {

	if bucketList == nil {
		return nil, "", fmt.Errorf("bucketList cannot be empty")
	}

	var (
		currentBucket     *bbolt.Bucket = nil
		currentBucketPath string        = bucketList[0]
	)

	if currentBucket = tx.Bucket([]byte(bucketList[0])); currentBucket == nil {
		return nil, currentBucketPath, fmt.Errorf("couldn't find bucket [%s]", currentBucketPath)
	}

	for _, item := range bucketList[1:] {
		currentBucketPath = fmt.Sprintf("%s/%s", currentBucketPath, item)

		// find the bucket
		if currentBucket = currentBucket.Bucket([]byte(item)); currentBucket == nil {
			return nil, currentBucketPath, fmt.Errorf("couldn't find bucket [%s]", currentBucketPath)
		}
	}

	return currentBucket, currentBucketPath, nil
}

// create buckets if not exists
func createBuckets(tx *bbolt.Tx, bucketList []string) (*bbolt.Bucket, string, error) {

	if bucketList == nil {
		return nil, "", fmt.Errorf("bucketList cannot be empty")
	}

	var (
		currentBucket     *bbolt.Bucket = nil
		currentBucketPath string        = bucketList[0]
		err               error         = nil
	)

	if currentBucket, err = tx.CreateBucketIfNotExists([]byte(bucketList[0])); err != nil {
		return nil, currentBucketPath, fmt.Errorf("couldn't create bucket [%s]: %s", currentBucketPath, err.Error())
	}

	for _, item := range bucketList[1:] {
		currentBucketPath = fmt.Sprintf("%s/%s", currentBucketPath, item)

		// create the bucket if not exist and change to new bucket
		if currentBucket, err = currentBucket.CreateBucketIfNotExists([]byte(item)); err != nil {
			return nil, currentBucketPath, fmt.Errorf("couldn't create bucket [%s]: %s", currentBucketPath, err.Error())
		}
	}

	return currentBucket, currentBucketPath, err
}
