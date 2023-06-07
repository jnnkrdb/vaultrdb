package v3

import (
	"context"

	"go.etcd.io/etcd/clientv3"
)

func getTTL(cli *clientv3.Client, lease int64) int64 {
	resp, err := cli.Lease.TimeToLive(context.Background(), clientv3.LeaseID(lease))
	if err != nil {
		return 0
	}
	if resp.TTL == -1 {
		return 0
	}
	return resp.TTL
}
