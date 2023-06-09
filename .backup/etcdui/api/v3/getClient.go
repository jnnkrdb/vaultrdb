package v3

import (
	"net/http"

	"github.com/jnnkrdb/vaultrdb/svc/etcdui/config"
	"go.etcd.io/etcd/clientv3"
)

func getClient(w http.ResponseWriter, r *http.Request) *clientv3.Client {
	sess := config.Sessionmgr.SessionStart(w, r)
	v := sess.Get("uinfo")
	if v != nil {
		uinfo := v.(*config.UserInfo)
		c, _ := newClient(uinfo)
		return c
	}
	return nil
}
