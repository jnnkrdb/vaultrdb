package v2

import (
	"context"
	"io"
	"log"
	"net/http"
	"strconv"

	"go.etcd.io/etcd/client"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	dir := r.FormValue("dir")
	log.Println("DELETE", "v2", key)

	kapi := client.NewKeysAPI(GetClient(w, r))

	isDir, _ := strconv.ParseBool(dir)
	if isDir {
		if _, err := kapi.Delete(context.Background(), key, &client.DeleteOptions{Recursive: true, Dir: true}); err != nil {
			io.WriteString(w, err.Error())
			return
		}
	} else {
		if _, err := kapi.Delete(context.Background(), key, nil); err != nil {
			io.WriteString(w, err.Error())
			return
		}
	}

	io.WriteString(w, "ok")
}
