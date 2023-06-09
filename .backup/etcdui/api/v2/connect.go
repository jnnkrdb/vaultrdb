package v2

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/jnnkrdb/vaultrdb/svc/etcdui/config"
)

var rootUsersV2 = make(map[string]*config.UserInfo)

func Connect(w http.ResponseWriter, r *http.Request) {
	config.Mu.Lock()
	defer config.Mu.Unlock()

	sess := config.Sessionmgr.SessionStart(w, r)
	host := strings.TrimSpace(r.FormValue("host"))
	uname := r.FormValue("uname")
	passwd := r.FormValue("passwd")
	if !strings.HasPrefix(host, "http") {
		host = "http://" + host
	}

	if config.UseAuth {
		_, ok := rootUsersV2[host]
		if !ok && uname != "root" {
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

	if uinfo, ok := sess.Get("uinfov2").(*config.UserInfo); ok {
		if host == uinfo.Host && uname == uinfo.Username && passwd == uinfo.Password {
			info := GetInfo(host)
			b, _ := json.Marshal(map[string]interface{}{"status": "running", "info": info})
			io.WriteString(w, string(b))
			return
		}
	}

	uinfo := &config.UserInfo{Host: host, Username: uname, Password: passwd}
	_, err := NewClient(uinfo)
	if err != nil {
		log.Println(r.Method, "v2", "connect fail.")
		b, _ := json.Marshal(map[string]interface{}{"status": "error", "message": err.Error()})
		io.WriteString(w, string(b))
		return
	}
	_ = sess.Set("uinfov2", uinfo)

	if config.UseAuth {
		if uname == "root" {
			rootUsersV2[host] = uinfo
		}
	} else {
		rootUsersV2[host] = uinfo
	}
	log.Println(r.Method, "v2", "connect success.")
	info := GetInfo(host)
	b, _ := json.Marshal(map[string]interface{}{"status": "running", "info": info})
	io.WriteString(w, string(b))
}
