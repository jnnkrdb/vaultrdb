package helpers

import (
	"net/http"
	"os"

	"github.com/jnnkrdb/vaultrdb/libs/logging"
)

// reads the given file and sends the content as byte
func FilesContent(w http.ResponseWriter, r *http.Request, file string) {

	if b, err := os.ReadFile(file); err != nil {

		logging.Log.Error(err, "error reading bytes from file", "file", file)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	} else {

		w.Write(b)
	}
}
