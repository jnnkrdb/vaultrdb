package v2

import (
	"context"
	"log"
	"strings"

	"go.etcd.io/etcd/client"
)

func GetInfo(host string) map[string]string {
	if !strings.HasPrefix(host, "http://") {
		host = "http://" + host
	}
	info := make(map[string]string)
	uinfo, ok := rootUsersV2[host]
	if ok {
		rootClient, err := NewClient(uinfo)
		if err != nil {
			log.Println(err)
			return info
		}
		ver, err := rootClient.GetVersion(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		memberKapi := client.NewMembersAPI(rootClient)
		member, err := memberKapi.Leader(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		info["version"] = ver.Server
		info["name"] = member.Name
		info["size"] = "unknow" // FIXME: How get?
	}
	return info
}
