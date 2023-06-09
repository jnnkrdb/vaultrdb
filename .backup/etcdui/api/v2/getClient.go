package v2

import (
	"net/http"

	"github.com/jnnkrdb/vaultrdb/svc/etcdui/config"
	"go.etcd.io/etcd/client"
)

func GetClient(w http.ResponseWriter, r *http.Request) client.Client {
	sess := config.Sessionmgr.SessionStart(w, r)
	v := sess.Get("uinfov2")
	if v != nil {
		uinfo := v.(*config.UserInfo)
		c, _ := NewClient(uinfo)
		return c
	}
	return nil
}
