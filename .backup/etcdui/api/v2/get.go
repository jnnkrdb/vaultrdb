package v2

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/jnnkrdb/vaultrdb/svc/etcdui/config"
	"go.etcd.io/etcd/client"
)

func Get(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	data := make(map[string]interface{})
	log.Println("GET", "v2", key)

	var cli client.Client
	sess := config.Sessionmgr.SessionStart(w, r)
	v := sess.Get("uinfov2")
	var uinfo *config.UserInfo
	if v != nil {
		uinfo = v.(*config.UserInfo)
		cli, _ = NewClient(uinfo)
		kapi := client.NewKeysAPI(cli)

		var permissions [][]string
		if r.FormValue("prefix") == "true" {
			var e error
			permissions, e = GetPermissionPrefix(uinfo.Host, uinfo.Username, key)
			if e != nil {
				io.WriteString(w, e.Error())
				return
			}
		} else {
			permissions = [][]string{{key, ""}}
		}

		var (
			min, max int
		)
		if key == config.Separator {
			min = 1
		} else {
			min = len(strings.Split(key, config.Separator))
		}
		max = min
		all := make(map[int][]map[string]interface{})
		if key == config.Separator {
			all[min] = []map[string]interface{}{{"key": key, "value": "", "dir": true, "nodes": make([]map[string]interface{}, 0)}}
		}
		for _, p := range permissions {
			pKey, pRange := p[0], p[1]
			var opt *client.GetOptions
			if pRange != "" {
				if pRange == "c" {
					pKey += config.Separator
				}
				opt = &client.GetOptions{Recursive: true, Sort: true}
			}
			if resp, err := kapi.Get(context.Background(), pKey, opt); err != nil {
				data["errorCode"] = 500
				data["message"] = err.Error()
			} else {
				if resp.Node == nil {
					data["errorCode"] = 500
					data["message"] = "The node does not exist."
				} else {
					max = config.GetNode(resp.Node, key, all, min, max)
				}
			}
		}

		//b, _ := json.MarshalIndent(all, "", "  ")
		//fmt.Println(string(b))

		// parent-child mapping
		for i := max; i > min; i-- {
			for _, a := range all[i] {
				for _, pa := range all[i-1] {
					if i == 2 { // The last is root
						pa["nodes"] = append(pa["nodes"].([]map[string]interface{}), a)
						pa["dir"] = true
					} else {
						if strings.HasPrefix(a["key"].(string), pa["key"].(string)+config.Separator) {
							pa["nodes"] = append(pa["nodes"].([]map[string]interface{}), a)
							pa["dir"] = true
						}
					}
				}
			}
		}

		for _, n := range all[min] {
			if n["key"] == key {
				config.NodesSort(n)
				data["node"] = n
				break
			}
		}
	}
	var dataByte []byte
	var err error
	if dataByte, err = json.Marshal(data); err != nil {
		io.WriteString(w, err.Error())
	} else {
		io.WriteString(w, string(dataByte))
	}
}
