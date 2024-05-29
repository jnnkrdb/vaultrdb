package internaldb

import (
	"encoding/json"
	"net/http"
	"slices"
	"vrdb-storage/server/configs"

	"github.com/gorilla/mux"
	"go.etcd.io/bbolt"
	"vrdb.go/logging"
)

// list all keys in the internaldb
//
// parameters will be handled in the future
func List_Keys(w http.ResponseWriter, r *http.Request) {

	// get the requested bucket
	bucket, ok := mux.Vars(r)["bucket"]
	if !ok {
		logging.Log.Info("bucket in query is missing", "response-code", http.StatusBadRequest, "query", mux.Vars(r))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// check if the bucket exists
	if !slices.Contains(configs.Buckets, bucket) {
		logging.Log.Info("bucket does not exist", "response-code", http.StatusBadRequest, "bucket", bucket)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	type res struct {
		K string `json:"key"`
		V string `json:"value"`
	}

	var result = []res{}

	// get all keyvalues from the required bucket
	configs.DB.View(func(tx *bbolt.Tx) error {

		var c = tx.Bucket([]byte(bucket)).Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {

			result = append(result, res{K: string(k), V: string(v)})
		}

		return nil
	})

	// translate into json and ship
	if err := json.NewEncoder(w).Encode(result); err != nil {

		logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "err", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
