package api_v1_buckets

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/internalstorage/vaultrdbstore"
	"github.com/jnnkrdb/vaultrdb/pkg/http/helpers"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

func List(w http.ResponseWriter, r *http.Request) {

	bucketpath, ok := mux.Vars(r)["bucketpath"]
	if !ok {
		logging.SLog.Warn("bucketpath in query is missing",
			"response-code", http.StatusBadRequest,
			"query", mux.Vars(r),
		)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	logging.SLog.Warn("bucketpath in query is the following",
		"bucketpath", bucketpath,
		"query", mux.Vars(r),
	)

	res, err := vaultrdbstore.DB.ReadBuckets(bucketpath)
	if err != nil {
		logging.SLog.Error("error reading bucket from path",
			"bucket", bucketpath,
			"response-code", http.StatusInternalServerError,
			"error", err.Error(),
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	helpers.StringValueJson{Values: res}.Send(w)
}
