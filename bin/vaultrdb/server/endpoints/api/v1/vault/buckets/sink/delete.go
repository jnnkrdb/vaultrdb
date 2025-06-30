package api_v1_buckets_sink

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/vaultrdbstore"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

func Delete(w http.ResponseWriter, r *http.Request) {

	bucketpath, ok := mux.Vars(r)["bucketpath"]
	if !ok {
		logging.Default.Warn("bucketpath in query is missing",
			"response-code", http.StatusBadRequest,
			"query", mux.Vars(r),
		)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	key, ok := mux.Vars(r)["key"]
	if !ok {
		logging.Default.Warn("key in query is missing",
			"response-code", http.StatusBadRequest,
			"query", mux.Vars(r),
		)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := vaultrdbstore.DB.DeleteKey(bucketpath, key); err != nil {
		logging.Default.Error("error removing key from bucket",
			"bucket", bucketpath,
			"key", key,
			"response-code", http.StatusInternalServerError,
			"err", err,
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("OK"))
}
