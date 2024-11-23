package metadata

import (
	"net/http"

	"github.com/jnnkrdb/vaultrdb/pkg/http/helpers"
)

const (
	URI_Metadata_Version string = "/meta/version"
	VERSION_FILE         string = "/opt/vaultrdb/config/VERSION"
)

func version(w http.ResponseWriter, r *http.Request) {
	helpers.FilesContent(w, r, VERSION_FILE)
}
