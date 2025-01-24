package helpers

import (
	"net/http"
	"os"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// reads the given file and sends the content as byte
func FilesContent(w http.ResponseWriter, r *http.Request, file string) {

	if b, err := os.ReadFile(file); err != nil {

		logging.SLog.Error("error reading bytes from file", "file", file, "error", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	} else {

		w.Write(b)
	}
}
