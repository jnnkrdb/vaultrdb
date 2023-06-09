package v3

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/jnnkrdb/vaultrdb/svc/etcdui/config"
	"go.etcd.io/etcd/clientv3"
)

func Get(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	key := r.FormValue("key")
	log.Println("GET", "v3", key)

	var cli *clientv3.Client
	sess := config.Sessionmgr.SessionStart(w, r)
	v := sess.Get("uinfo")
	var uinfo *config.UserInfo
	if v != nil {
		uinfo = v.(*config.UserInfo)
		cli, _ = newClient(uinfo)
		defer cli.Close()

		permissions, e := getPermissionPrefix(uinfo.Host, uinfo.Username, key)
		if e != nil {
			io.WriteString(w, e.Error())
			return
		}
		if r.FormValue("prefix") == "true" {
			pnode := make(map[string]interface{})
			pnode["key"] = key
			pnode["nodes"] = make([]map[string]interface{}, 0)
			for _, p := range permissions {
				var (
					resp *clientv3.GetResponse
					err  error
				)
				if p[1] != "" {
					prefixKey := p[0]
					if p[0] == "/" {
						prefixKey = ""
					}
					resp, err = cli.Get(context.Background(), prefixKey, clientv3.WithPrefix())
				} else {
					resp, err = cli.Get(context.Background(), p[0])
				}
				if err != nil {
					data["errorCode"] = 500
					data["message"] = err.Error()
				} else {
					for _, kv := range resp.Kvs {
						node := make(map[string]interface{})
						node["key"] = string(kv.Key)
						node["value"] = string(kv.Value)
						node["dir"] = false
						if key == string(kv.Key) {
							node["ttl"] = getTTL(cli, kv.Lease)
						} else {
							node["ttl"] = 0
						}
						node["createdIndex"] = kv.CreateRevision
						node["modifiedIndex"] = kv.ModRevision
						nodes := pnode["nodes"].([]map[string]interface{})
						pnode["nodes"] = append(nodes, node)
					}
				}
			}
			data["node"] = pnode
		} else {
			if resp, err := cli.Get(context.Background(), key); err != nil {
				data["errorCode"] = 500
				data["message"] = err.Error()
			} else {
				if resp.Count > 0 {
					kv := resp.Kvs[0]
					node := make(map[string]interface{})
					node["key"] = string(kv.Key)
					node["value"] = string(kv.Value)
					node["dir"] = false
					node["ttl"] = getTTL(cli, kv.Lease)
					node["createdIndex"] = kv.CreateRevision
					node["modifiedIndex"] = kv.ModRevision
					data["node"] = node
				} else {
					data["errorCode"] = 500
					data["message"] = "The node does not exist."
				}
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
