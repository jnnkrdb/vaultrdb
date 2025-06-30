package api_v1_buckets_sink

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/vaultrdb/bin/vaultrdb/conf"
	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

func Write(w http.ResponseWriter, r *http.Request) {
	var log = logging.FromContext(r.Context())

	bucket, ok := mux.Vars(r)["bucket"]
	if !ok {
		log.Warn("bucket in query is missing",
			"response-code", http.StatusBadRequest,
			"query", mux.Vars(r),
		)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	key, ok := mux.Vars(r)["key"]
	if !ok {
		log.Warn("key in query is missing",
			"response-code", http.StatusBadRequest,
			"query", mux.Vars(r),
		)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	if b, err := io.ReadAll(r.Body); err != nil {
		log.Warn("couldn't read content from body",
			"response-code", http.StatusBadRequest,
			"err", err.Error(),
		)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return

	} else {

		if err := conf.Vault.WriteKey(bucket, key, string(b)); err != nil {
			log.Error("error receiving key from bucket",
				"bucket", bucket,
				"key", key,
				"response-code", http.StatusInternalServerError,
				"err", err,
			)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	w.Write([]byte("OK"))
}
