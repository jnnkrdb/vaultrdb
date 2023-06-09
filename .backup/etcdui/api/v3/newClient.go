package v3

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/jnnkrdb/vaultrdb/svc/etcdui/config"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/pkg/transport"
	"google.golang.org/grpc"
)

func newClient(uinfo *config.UserInfo) (*clientv3.Client, error) {
	endpoints := []string{uinfo.Host}
	var err error

	// use tls if usetls is true
	var tlsConfig *tls.Config
	if config.UseTLS {
		tlsInfo := transport.TLSInfo{
			CertFile:      config.Cert,
			KeyFile:       config.KeyFile,
			TrustedCAFile: config.CACert,
		}
		tlsConfig, err = tlsInfo.ClientConfig()
		if err != nil {
			log.Println(err.Error())
		}
	}

	conf := clientv3.Config{
		Endpoints:          endpoints,
		DialTimeout:        time.Second * time.Duration(config.ConnectTimeout),
		TLS:                tlsConfig,
		DialOptions:        []grpc.DialOption{grpc.WithBlock()},
		MaxCallSendMsgSize: config.SendMessageSize,
	}
	if config.UseAuth {
		conf.Username = uinfo.Username
		conf.Password = uinfo.Password
	}

	var c *clientv3.Client
	c, err = clientv3.New(conf)
	if err != nil {
		return nil, err
	}
	return c, nil
}
