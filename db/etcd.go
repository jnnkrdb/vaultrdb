package db

import (
	"os"
	"strings"
	"time"

	"github.com/go-logr/logr"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	ETCD     *clientv3.Client = nil
	USE_ETCD bool             = false
)

/*
	Environment Variables:
	- ETCD_ENDPOINTS: "<host>:<port>;<host>:<port>;<host>:<port>"
	- ETCD_USERNAME: "<username>"
	- ETCD_PASSWORD: "<password>"

*/

func InitETCDConnection(_log logr.Logger) error {

	var eps = os.Getenv("ETCD_ENDPOINTS")

	_log.Info("initializing connection to etcd if any", "endpoints", eps)

	if len(eps) > 0 {

		// create the config
		var conf = clientv3.Config{
			Endpoints:   strings.Split(eps, ";"),
			DialTimeout: 10 * time.Second,
		}

		// check multiple settings

		if user, set := os.LookupEnv("ETCD_USERNAME"); set {
			conf.Username = user
		}

		if pass, set := os.LookupEnv("ETCD_PASSWORD"); set {
			conf.Password = pass
		}

		// initialize the etcd connection
		if cli, err := clientv3.New(conf); err != nil {
			return err
		} else {
			ETCD = cli
			USE_ETCD = true
			_log.Info("connection initialized")
		}
	}

	return nil
}
