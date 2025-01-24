package storedb

import (
	"strings"
	"sync"
	"time"

	"go.etcd.io/bbolt"
)

var (
	DB         *bbolt.DB  = nil
	writeMutex sync.Mutex = sync.Mutex{}
)

// this opens a bolt file, if not existent, creates one and
// serves the db as the internal key/value store
func OpenDB(path string, timeout time.Duration) (err error) {

	var options *bbolt.Options = &bbolt.Options{
		Timeout: timeout,
	}

	if DB == nil {
		DB, err = bbolt.Open(path, 0600, options)
	}

	return
}

// close the db if opened
func CloseDB() error {
	if err := DB.Close(); err != nil {
		return err
	}

	DB = nil

	return nil
}

// calculate the correct bucket walking path
func calculateBucketsFromPath(path string) []string {

	// replacing spaces with ""
	path = strings.ReplaceAll(path, " ", "")

	if path == "" || path == "/" {
		return nil
	}

	// removing prefix or suffix /
	path, _ = strings.CutPrefix(path, "/")
	path, _ = strings.CutSuffix(path, "/")
	return strings.Split(path, "/")
}
