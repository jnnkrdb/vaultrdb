package api_v1_buckets

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/vaultrdbstore"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

func Delete(w http.ResponseWriter, r *http.Request) {

	bucketpath, ok := mux.Vars(r)["bucketpath"]
	if !ok {
		logging.SLog.Warn("bucketpath in query is missing",
			"response-code", http.StatusBadRequest,
			"query", mux.Vars(r),
		)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	bucket, ok := mux.Vars(r)["bucketpath"]
	if !ok {
		logging.SLog.Warn("bucket in query is missing",
			"response-code", http.StatusBadRequest,
			"query", mux.Vars(r),
		)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := vaultrdbstore.DB.DeleteBucket(bucketpath, bucket); err != nil {
		logging.SLog.Error("deleting bucket failed",
			"response-code", http.StatusInternalServerError,
			"error", err.Error(),
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("OK"))
}
