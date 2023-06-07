package v3

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/svc/etcdui/config"
	"go.etcd.io/etcd/clientv3"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	cli := getClient(w, r)
	defer cli.Close()
	key := r.FormValue("key")
	dir := r.FormValue("dir")
	log.Println("DELETE", "v3", key)

	if _, err := cli.Delete(context.Background(), key); err != nil {
		io.WriteString(w, err.Error())
		return
	}

	if dir == "true" {
		if _, err := cli.Delete(context.Background(), key+config.Separator, clientv3.WithPrefix()); err != nil {
			io.WriteString(w, err.Error())
			return
		}
	}
	io.WriteString(w, "ok")
}
