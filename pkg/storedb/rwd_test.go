package storedb

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func Test_ReadWriteDelete(t *testing.T) {

	var tempDB = fmt.Sprintf("%s/test.db", t.TempDir())
	OpenDB(tempDB, 5*time.Second, false)
	defer CloseDB()
	t.Logf("using db: %s", db.db.Path())

	// prepare a bunch of key/value pairs in different
	var tests = []struct {
		name  string
		path  string
		key   string
		value string
	}{
		{name: "test-1", path: "test", key: "firstKey", value: "First Test Value"},
		{name: "test-2", path: "test/asdf", key: "firstKey", value: "THIS IS REALLY GREAT!"},
		{name: "test-3", path: "test/asdf", key: "secondKey", value: "LOREM IPSUM..."},
		{name: "test-4", path: "test/asdf/1234", key: "vKey", value: "First Test Value"},
		{name: "test-5", path: "abcd/fghj/aaaaaa", key: "puschel", value: "VaultRDB is the best"},
	}

	// insert into the db
	for _, i := range tests {
		t.Run(fmt.Sprintf("create-%s", i.name), func(t *testing.T) {
			if WriteKey(i.path, i.key, i.value) != nil {
				t.Fatalf("error writing testvalue to database")
			}
		})
	}

	// closing the database connection and reopen it
	CloseDB()
	OpenDB(tempDB, 5*time.Second, false)

	for _, i := range tests {
		t.Run(fmt.Sprintf("read-%s", i.name), func(t *testing.T) {
			result, err := ReadKey(i.path, i.key)
			t.Logf("path: [%s:%s] estimated: %s - testResult: %v", i.path, i.key, i.value, result)
			if err != nil {
				t.Fatalf("error reading [%s:%s] from tempDB: %s", i.path, i.key, err.Error())
			}
			if result != i.value {
				t.Fatalf("error: estimated [%s] does not equal the result [%s]", i.value, result)
			}
		})
	}

	// closing the database connection and reopen it
	CloseDB()
	OpenDB(tempDB, 5*time.Second, false)

	// remove the created objects
	t.Run("delete-key-1", func(t *testing.T) {

		t.Logf("removing key [%s:%s]", tests[0].path, tests[0].key)

		if err := DeleteKey(tests[0].path, tests[0].key); err != nil {
			t.Fatalf("error removing key[%s:%s] from database: %s", tests[0].path, tests[0].key, err.Error())
		}

		if _, err := ReadKey(tests[0].path, tests[0].key); err == nil {
			t.Fatalf("error: estimated [%s:%s] to be gone", tests[0].path, tests[0].key)
		}
	})

	t.Run("delete-key-2", func(t *testing.T) {

		t.Logf("removing key [%s:%s]", tests[1].path, tests[1].key)

		if err := DeleteKey(tests[1].path, tests[1].key); err != nil {
			t.Fatalf("error removing key[%s:%s] from database: %s", tests[1].path, tests[1].key, err.Error())
		}

		if _, err := ReadKey(tests[1].path, tests[1].key); err == nil {
			t.Fatalf("error: estimated [%s:%s] to be gone", tests[1].path, tests[1].key)
		}
	})

	t.Run("delete-bucket-1", func(t *testing.T) {

		var path, bucket string = strings.Split(tests[4].path, "/")[0], strings.Split(tests[4].path, "/")[1]

		t.Logf("removing bucket [%s:%s]", path, bucket)

		if err := DeleteBucket(path, bucket); err != nil {
			t.Fatalf("error removing bucket[%s:%s] from database: %s", path, bucket, err.Error())
		}

		if _, err := ReadKey(tests[4].path, tests[4].key); err == nil {
			t.Fatalf("error: estimated [%s:%s] to be gone", tests[4].path, tests[4].key)
		}
	})

	t.Run("delete-bucket-2", func(t *testing.T) {

		var path, bucket string = "", "test"

		t.Logf("removing bucket [%s:%s]", path, bucket)

		if err := DeleteBucket(path, bucket); err != nil {
			t.Fatalf("error removing bucket[%s:%s] from database: %s", path, bucket, err.Error())
		}

		if _, err := ReadKey("test/asdf/vKey", "test"); err == nil {
			t.Fatalf("error: estimated [%s:%s] to be gone", tests[3].path, tests[3].key)
		}
	})
}
