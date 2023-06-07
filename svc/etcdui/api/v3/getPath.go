package v3

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/jnnkrdb/vaultrdb/svc/etcdui/config"
	"go.etcd.io/etcd/clientv3"
)

func GetPath(w http.ResponseWriter, r *http.Request) {
	originKey := r.FormValue("key")
	log.Println("GET", "v3", originKey)
	var (
		data = make(map[string]interface{})
		/*
			{1:["/"], 2:["/foo", "/foo2"], 3:["/foo/bar", "/foo2/bar"], 4:["/foo/bar/test"]}
		*/
		all = make(map[int][]map[string]interface{})
		min int
		max int
		//prefixKey string
	)

	var cli *clientv3.Client
	sess := config.Sessionmgr.SessionStart(w, r)
	v := sess.Get("uinfo")
	var uinfo *config.UserInfo
	if v != nil {
		uinfo = v.(*config.UserInfo)
		cli, _ = newClient(uinfo)
		defer cli.Close()

		permissions, e := getPermissionPrefix(uinfo.Host, uinfo.Username, originKey)
		if e != nil {
			io.WriteString(w, e.Error())
			return
		}

		// parent
		var (
			presp *clientv3.GetResponse
			err   error
		)
		if originKey != config.Separator {
			presp, err = cli.Get(context.Background(), originKey)
			if err != nil {
				data["errorCode"] = 500
				data["message"] = err.Error()
				dataByte, _ := json.Marshal(data)
				io.WriteString(w, string(dataByte))
				return
			}
		}
		if originKey == config.Separator {
			min = 1
			//prefixKey = seperator
		} else {
			min = len(strings.Split(originKey, config.Separator))
			//prefixKey = originKey
		}
		max = min
		all[min] = []map[string]interface{}{{"key": originKey}}
		if presp != nil && presp.Count != 0 {
			all[min][0]["value"] = string(presp.Kvs[0].Value)
			all[min][0]["ttl"] = getTTL(cli, presp.Kvs[0].Lease)
			all[min][0]["createdIndex"] = presp.Kvs[0].CreateRevision
			all[min][0]["modifiedIndex"] = presp.Kvs[0].ModRevision
		}
		all[min][0]["nodes"] = make([]map[string]interface{}, 0)

		for _, p := range permissions {
			key, rangeEnd := p[0], p[1]
			//child
			var resp *clientv3.GetResponse
			if rangeEnd != "" {
				resp, err = cli.Get(context.Background(), key, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
			} else {
				resp, err = cli.Get(context.Background(), key, clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
			}
			if err != nil {
				data["errorCode"] = 500
				data["message"] = err.Error()
				dataByte, _ := json.Marshal(data)
				io.WriteString(w, string(dataByte))
				return
			}

			for _, kv := range resp.Kvs {
				if string(kv.Key) == config.Separator {
					continue
				}
				keys := strings.Split(string(kv.Key), config.Separator) // /foo/bar
				for i := range keys {                                   // ["", "foo", "bar"]
					k := strings.Join(keys[0:i+1], config.Separator)
					if k == "" {
						continue
					}
					node := map[string]interface{}{"key": k}
					if node["key"].(string) == string(kv.Key) {
						node["value"] = string(kv.Value)
						if key == string(kv.Key) {
							node["ttl"] = getTTL(cli, kv.Lease)
						} else {
							node["ttl"] = 0
						}
						node["createdIndex"] = kv.CreateRevision
						node["modifiedIndex"] = kv.ModRevision
					}
					level := len(strings.Split(k, config.Separator))
					if level > max {
						max = level
					}

					if _, ok := all[level]; !ok {
						all[level] = make([]map[string]interface{}, 0)
					}
					levelNodes := all[level]
					var isExist bool
					for _, n := range levelNodes {
						if n["key"].(string) == k {
							isExist = true
						}
					}
					if !isExist {
						node["nodes"] = make([]map[string]interface{}, 0)
						all[level] = append(all[level], node)
					}
				}
			}
		}

		// parent-child mapping
		for i := max; i > min; i-- {
			for _, a := range all[i] {
				for _, pa := range all[i-1] {
					if i == 2 {
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
	}
	data = all[min][0]
	if dataByte, err := json.Marshal(map[string]interface{}{"node": data}); err != nil {
		io.WriteString(w, err.Error())
	} else {
		io.WriteString(w, string(dataByte))
	}
}
