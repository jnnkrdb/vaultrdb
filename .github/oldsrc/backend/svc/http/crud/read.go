package crud

import (
	"encoding/json"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/svc/sqlite3"
)

func Read(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("uid") {
		if result, err := sqlite3.SelectAllPairs(); err != nil {
			w.Header().Add("error", err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(result)
		}
	} else {
		if result, err := sqlite3.SelectPairByUID(r.URL.Query().Get("uid")); err != nil {
			w.Header().Add("error", err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(result)
		}
	}
}
