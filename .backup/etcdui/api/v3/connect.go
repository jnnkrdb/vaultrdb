package v3

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/svc/etcdui/config"
)

var rootUsers = make(map[string]*config.UserInfo)

func Connect(w http.ResponseWriter, r *http.Request) {
	config.Mu.Lock()
	defer config.Mu.Unlock()
	sess := config.Sessionmgr.SessionStart(w, r)
	host := r.FormValue("host")
	uname := r.FormValue("uname")
	passwd := r.FormValue("passwd")

	if config.UseAuth {
		if _, ok := rootUsers[host]; !ok && uname != "root" { // no root user
			b, _ := json.Marshal(map[string]interface{}{"status": "root"})
			io.WriteString(w, string(b))
			return
		}
		if uname == "" || passwd == "" {
			b, _ := json.Marshal(map[string]interface{}{"status": "login"})
			io.WriteString(w, string(b))
			return
		}
	}

	if uinfo, ok := sess.Get("uinfo").(*config.UserInfo); ok {
		if host == uinfo.Host && uname == uinfo.Username && passwd == uinfo.Password {
			info := getInfo(host)
			b, _ := json.Marshal(map[string]interface{}{"status": "running", "info": info})
			io.WriteString(w, string(b))
			return
		}
	}

	uinfo := &config.UserInfo{Host: host, Username: uname, Password: passwd}
	c, err := newClient(uinfo)
	if err != nil {
		log.Println(r.Method, "v3", "connect fail.")
		b, _ := json.Marshal(map[string]interface{}{"status": "error", "message": err.Error()})
		io.WriteString(w, string(b))
		return
	}
	defer c.Close()
	_ = sess.Set("uinfo", uinfo)

	if config.UseAuth {
		if uname == "root" {
			rootUsers[host] = uinfo
		}
	} else {
		rootUsers[host] = uinfo
	}
	log.Println(r.Method, "v3", "connect success.")
	info := getInfo(host)
	b, _ := json.Marshal(map[string]interface{}{"status": "running", "info": info})
	io.WriteString(w, string(b))
}
