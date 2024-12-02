package api_v1_configs

import (
	"encoding/json"
	"net/http"
	"slices"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/configstore"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/configstore/buckets"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
	"go.etcd.io/bbolt"
)

// list all keys in the internaldb
//
// parameters will be handled in the future
func List_Keys(w http.ResponseWriter, r *http.Request) {

	// get the requested bucket
	bucket, ok := mux.Vars(r)["bucket"]
	if !ok {
		logging.SLog.Info("bucket in query is missing",
			"response-code", http.StatusBadRequest,
			"query", mux.Vars(r),
		)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// check if the bucket exists
	if !slices.Contains(buckets.DefaultBuckets(), bucket) {
		logging.SLog.Info("bucket does not exist", "response-code", http.StatusBadRequest, "bucket", bucket)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	type kv struct {
		K string `json:"key"`
		V string `json:"value"`
	}

	type bucketSink struct {
		Sink []kv `json:"sink"`
	}

	var result = bucketSink{
		Sink: []kv{},
	}

	// get all keyvalues from the required bucket
	configstore.DB.View(func(tx *bbolt.Tx) error {

		var c = tx.Bucket([]byte(bucket)).Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {

			result.Sink = append(result.Sink, kv{K: string(k), V: string(v)})
		}

		return nil
	})

	// translate into json and ship
	if err := json.NewEncoder(w).Encode(result); err != nil {

		logging.SLog.Info("error parsing result into json",
			"response-code", http.StatusInternalServerError,
			"err", err,
		)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
