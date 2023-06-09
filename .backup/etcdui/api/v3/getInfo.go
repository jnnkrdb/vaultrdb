package v3

import (
	"context"
	"fmt"
	"log"
)

func getInfo(host string) map[string]string {
	info := make(map[string]string)
	uinfo := rootUsers[host]
	rootClient, err := newClient(uinfo)
	if err != nil {
		log.Println(err)
		return info
	}
	defer rootClient.Close()

	status, err := rootClient.Status(context.Background(), host)
	if err != nil {
		log.Fatal(err)
	}
	mems, err := rootClient.MemberList(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	kb := 1024
	mb := kb * 1024
	gb := mb * 1024
	var sizeStr string
	for _, m := range mems.Members {
		if m.ID == status.Leader {
			info["version"] = status.Version
			gn, rem1 := size(int(status.DbSize), gb)
			mn, rem2 := size(rem1, mb)
			kn, bn := size(rem2, kb)
			if sizeStr != "" {
				sizeStr += " "
			}
			if gn > 0 {
				info["size"] = fmt.Sprintf("%dG", gn)
			} else {
				if mn > 0 {
					info["size"] = fmt.Sprintf("%dM", mn)
				} else {
					if kn > 0 {
						info["size"] = fmt.Sprintf("%dK", kn)
					} else {
						info["size"] = fmt.Sprintf("%dByte", bn)
					}
				}
			}
			info["name"] = m.GetName()
			break
		}
	}
	return info
}
