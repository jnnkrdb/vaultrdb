package v2

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.etcd.io/etcd/client"
)

func Put(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	value := r.FormValue("value")
	ttl := r.FormValue("ttl")
	dir := r.FormValue("dir")
	log.Println("PUT", "v2", key)

	kapi := client.NewKeysAPI(GetClient(w, r))

	var isDir bool
	if dir != "" {
		isDir, _ = strconv.ParseBool(dir)
	}
	var err error
	data := make(map[string]interface{})
	if ttl != "" {
		var sec int64
		sec, err = strconv.ParseInt(ttl, 10, 64)
		if err != nil {
			log.Println(err.Error())
		}
		_, err = kapi.Set(context.Background(), key, value, &client.SetOptions{TTL: time.Duration(sec) * time.Second, Dir: isDir})
	} else {
		_, err = kapi.Set(context.Background(), key, value, &client.SetOptions{Dir: isDir})
	}
	if err != nil {
		data["errorCode"] = 500
		data["message"] = err.Error()
	} else {
		if resp, err := kapi.Get(context.Background(), key, &client.GetOptions{Recursive: true, Sort: true}); err != nil {
			data["errorCode"] = err.Error()
		} else {
			if resp.Node != nil {
				node := make(map[string]interface{})
				node["key"] = resp.Node.Key
				node["value"] = resp.Node.Value
				node["dir"] = resp.Node.Dir
				node["ttl"] = resp.Node.TTL
				node["createdIndex"] = resp.Node.CreatedIndex
				node["modifiedIndex"] = resp.Node.ModifiedIndex
				data["node"] = node
			}
		}
	}

	var dataByte []byte
	if dataByte, err = json.Marshal(data); err != nil {
		io.WriteString(w, err.Error())
	} else {
		io.WriteString(w, string(dataByte))
	}
}
