package api_v1_buckets_sink

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/vaultrdbstore"
	"github.com/jnnkrdb/vaultrdb/pkg/http/helpers"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

func Find(w http.ResponseWriter, r *http.Request) {

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

	res, err := vaultrdbstore.DB.ReadKey(bucketpath, key)
	if err != nil {
		logging.Default.Error("error receiving key from bucket",
			"bucket", bucketpath,
			"key", key,
			"response-code", http.StatusInternalServerError,
			"err", err,
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	helpers.StringValueJson{Value: res}.Send(w)
}
