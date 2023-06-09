package v2

import (
	"time"

	"github.com/jnnkrdb/vaultrdb/svc/etcdui/config"
	"go.etcd.io/etcd/client"
	"go.etcd.io/etcd/pkg/transport"
)

func NewClient(uinfo *config.UserInfo) (client.Client, error) {
	var transportConf client.CancelableTransport = nil
	if config.UseTLS {
		tlsInfo := transport.TLSInfo{
			CertFile:      config.Cert,
			KeyFile:       config.KeyFile,
			TrustedCAFile: config.CACert,
		}
		conf, err := transport.NewTransport(tlsInfo, time.Second*time.Duration(config.ConnectTimeout))
		if err == nil {
			transportConf = conf
		}
	}
	cfg := client.Config{
		Endpoints:               []string{uinfo.Host},
		HeaderTimeoutPerRequest: time.Second * time.Duration(config.ConnectTimeout),
		Transport:               transportConf,
	}

	if config.UseAuth {
		cfg.Username = uinfo.Username
		cfg.Password = uinfo.Password
	}

	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	return c, nil
}
