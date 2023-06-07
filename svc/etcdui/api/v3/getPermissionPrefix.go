package v3

import (
	"context"

	"github.com/jnnkrdb/vaultrdb/svc/etcdui/config"
)

func getPermissionPrefix(host, uname, key string) ([][]string, error) {
	if !config.UseAuth {
		return [][]string{{key, "p"}}, nil // No auth return all
	} else {
		if uname == "root" {
			return [][]string{{key, "p"}}, nil
		}
		rootUser := rootUsers[host]
		rootCli, err := newClient(rootUser)
		if err != nil {
			return nil, err
		}
		defer rootCli.Close()

		if resp, err := rootCli.UserList(context.Background()); err != nil {
			return nil, err
		} else {
			// Find user permissions
			set := make(map[string]string)
			for _, u := range resp.Users {
				if u == uname {
					ur, err := rootCli.UserGet(context.Background(), u)
					if err != nil {
						return nil, err
					}
					for _, r := range ur.Roles {
						rr, err := rootCli.RoleGet(context.Background(), r)
						if err != nil {
							return nil, err
						}
						for _, p := range rr.Perm {
							set[string(p.Key)] = string(p.RangeEnd)
						}
					}
					break
				}
			}
			var pers [][]string
			for k, v := range set {
				pers = append(pers, []string{k, v})
			}
			return pers, nil
		}
	}
}
