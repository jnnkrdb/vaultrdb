package crud

import (
	"encoding/json"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/svc/sqlite3"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var p sqlite3.Pair
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.Header().Add("error", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if res, err := sqlite3.UpdatePair(p); err != nil {
		w.Header().Add("error", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	} else {
		json.NewEncoder(w).Encode(res)
	}
}
