package v2

import (
	"strings"

	"github.com/jnnkrdb/vaultrdb/svc/etcdui/config"
	"go.etcd.io/etcd/client"
)

func nodesSort(node map[string]interface{}) {
	if v, ok := node["nodes"]; ok && v != nil {
		a := v.([]map[string]interface{})
		if len(a) != 0 {
			for i := 0; i < len(a)-1; i++ {
				nodesSort(a[i])
				for j := i + 1; j < len(a); j++ {
					if a[j]["key"].(string) < a[i]["key"].(string) {
						a[i], a[j] = a[j], a[i]
					}
				}
			}
			nodesSort(a[len(a)-1])
		}
	}
}

func getNode(node *client.Node, selKey string, all map[int][]map[string]interface{}, min, max int) int {
	keys := strings.Split(node.Key, config.Separator) // /foo/bar
	if len(keys) < min && strings.HasPrefix(node.Key, selKey) {
		return max
	}
	for i := range keys { // ["", "foo", "bar"]
		k := strings.Join(keys[0:i+1], config.Separator)
		if k == "" {
			continue
		}
		nodeMap := map[string]interface{}{"key": k, "dir": true, "nodes": make([]map[string]interface{}, 0)}
		if k == node.Key {
			nodeMap["value"] = node.Value
			nodeMap["dir"] = node.Dir
			nodeMap["ttl"] = node.TTL
			nodeMap["createdIndex"] = node.CreatedIndex
			nodeMap["modifiedIndex"] = node.ModifiedIndex
		}
		keylevel := len(strings.Split(k, config.Separator))
		if keylevel > max {
			max = keylevel
		}

		if _, ok := all[keylevel]; !ok {
			all[keylevel] = make([]map[string]interface{}, 0)
		}
		var isExist bool
		for _, n := range all[keylevel] {
			if n["key"].(string) == k {
				isExist = true
			}
		}
		if !isExist {
			all[keylevel] = append(all[keylevel], nodeMap)
		}
	}

	if len(node.Nodes) != 0 {
		for _, n := range node.Nodes {
			max = getNode(n, selKey, all, min, max)
		}
	}
	return max
}
