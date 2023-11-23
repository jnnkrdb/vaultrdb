package crud

import (
	"net/http"

	"github.com/jnnkrdb/vaultrdb/svc/sqlite3"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("uid") {
		w.Header().Add("error", "uid is missing")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := sqlite3.DeletePair(r.URL.Query().Get("uid")); err != nil {
		w.Header().Add("error", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("OK"))
}
