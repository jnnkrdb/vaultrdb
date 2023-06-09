package config

import (
	"sync"

	"github.com/jnnkrdb/vaultrdb/svc/etcdui/session"
)

var (
	UseETCDWebUi bool

	ETCDUIHost string
	ETCDUIPort int

	Separator string

	UseTLS  bool
	CACert  string
	Cert    string
	KeyFile string

	UseAuth         bool
	ConnectTimeout  int
	SendMessageSize int

	Sessionmgr *session.Manager
	Mu         sync.Mutex
)
