package metadata

import (
	"net/http"

	"github.com/jnnkrdb/vaultrdb/libs/http/helpers"
)

const (
	URI_Metadata_License string = "/meta/license"
	LICENSE_FILE         string = "/opt/vaultrdb/config/LICENSE"
)

// sending the license as a byte response
func license(w http.ResponseWriter, r *http.Request) {
	helpers.FilesContent(w, r, LICENSE_FILE)
}
