package v2

import (
	"context"
	"sort"
	"strings"

	"github.com/jnnkrdb/vaultrdb/svc/etcdui/config"
	"go.etcd.io/etcd/client"
)

func GetPermissionPrefix(host, uname, key string) ([][]string, error) {
	if config.UseAuth {
		return [][]string{{key, "p"}}, nil // No auth return all
	} else {
		if uname == "root" {
			return [][]string{{key, "p"}}, nil
		}

		if !strings.HasPrefix(host, "http://") {
			host = "http://" + host
		}
		rootUser := rootUsersV2[host]
		rootCli, err := NewClient(rootUser)
		if err != nil {
			return nil, err
		}
		rootUserKapi := client.NewAuthUserAPI(rootCli)
		rootRoleKapi := client.NewAuthRoleAPI(rootCli)

		if users, err := rootUserKapi.ListUsers(context.Background()); err != nil {
			return nil, err
		} else {
			// Find user permissions
			set := make(map[string]string)
			for _, u := range users {
				if u == uname {
					user, err := rootUserKapi.GetUser(context.Background(), u)
					if err != nil {
						return nil, err
					}
					for _, r := range user.Roles {
						role, err := rootRoleKapi.GetRole(context.Background(), r)
						if err != nil {
							return nil, err
						}
						for _, ks := range role.Permissions.KV.Read {
							var k string
							if strings.HasSuffix(ks, "*") {
								k = ks[:len(ks)-1]
								set[k] = "p"
							} else if strings.HasSuffix(ks, "/*") {
								k = ks[:len(ks)-2]
								set[k] = "c"
							} else {
								if _, ok := set[ks]; !ok {
									set[ks] = ""
								}
							}
						}
					}
					break
				}
			}
			var pers [][]string
			var ks []string
			for k := range set {
				ks = append(ks, k)
			}
			sort.Strings(ks)
			for _, k := range ks {
				pers = append(pers, []string{k, set[k]})
			}
			return pers, nil
		}
	}
}
