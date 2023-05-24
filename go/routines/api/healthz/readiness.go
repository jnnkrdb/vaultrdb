package healthz

import (
	"fmt"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/settings"
)

// healthz check -> readyness
func Readiness(w http.ResponseWriter, r *http.Request) {

	var errors = []error{}

	// check sql connection
	if err := settings.PSQL.Ping(); err != nil {

		errors = append(errors, err)
	}

	// converting error messages to a connected string
	if len(errors) > 0 {

		var errormsg string

		for _, e := range errors {

			errormsg += fmt.Sprintf("%v", e)
		}

		http.Error(w, errormsg, http.StatusInternalServerError)

	} else {

		w.Write([]byte("OK"))
	}
}
