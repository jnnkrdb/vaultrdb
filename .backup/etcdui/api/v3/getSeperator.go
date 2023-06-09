package v3

import (
	"io"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/svc/etcdui/config"
)

func GetSeparator(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, config.Separator)
}
