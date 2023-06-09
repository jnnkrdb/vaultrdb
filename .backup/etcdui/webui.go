package etcdui

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-logr/logr"
	hdnlrs "github.com/jnnkrdb/gomw/handlers"
	"github.com/jnnkrdb/gomw/middlewares"
	"github.com/jnnkrdb/vaultrdb/svc/etcdui/config"
	"github.com/jnnkrdb/vaultrdb/svc/etcdui/session"

	apiv2 "github.com/jnnkrdb/vaultrdb/svc/etcdui/api/v2"
	apiv3 "github.com/jnnkrdb/vaultrdb/svc/etcdui/api/v3"
)

var httpHandlers = hdnlrs.HttpFunctionSet{

	// web ui
	{Pattern: "/", MainHandler: http.FileServer(http.Dir("/app/assets")), Middlewares: middlewares.MiddleWareChain{}},

	// version 2 api
	{Pattern: "/v2/separator", MainHandler: http.HandlerFunc(apiv3.GetSeparator), Middlewares: middlewares.MiddleWareChain{}},
	{Pattern: "/v2/connect", MainHandler: http.HandlerFunc(apiv2.Connect), Middlewares: middlewares.MiddleWareChain{}},
	{Pattern: "/v2/put", MainHandler: http.HandlerFunc(apiv2.Put), Middlewares: middlewares.MiddleWareChain{}},
	{Pattern: "/v2/get", MainHandler: http.HandlerFunc(apiv2.Get), Middlewares: middlewares.MiddleWareChain{}},
	{Pattern: "/v2/delete", MainHandler: http.HandlerFunc(apiv2.Delete), Middlewares: middlewares.MiddleWareChain{}},

	// version 3 api
	{Pattern: "/v3/separator", MainHandler: http.HandlerFunc(apiv3.GetSeparator), Middlewares: middlewares.MiddleWareChain{}},
	{Pattern: "/v3/getpath", MainHandler: http.HandlerFunc(apiv3.GetPath), Middlewares: middlewares.MiddleWareChain{}},
	{Pattern: "/v3/connect", MainHandler: http.HandlerFunc(apiv3.Connect), Middlewares: middlewares.MiddleWareChain{}},
	{Pattern: "/v3/put", MainHandler: http.HandlerFunc(apiv3.Put), Middlewares: middlewares.MiddleWareChain{}},
	{Pattern: "/v3/get", MainHandler: http.HandlerFunc(apiv3.Get), Middlewares: middlewares.MiddleWareChain{}},
	{Pattern: "/v3/delete", MainHandler: http.HandlerFunc(apiv3.Delete), Middlewares: middlewares.MiddleWareChain{}},
}

func StartETCDUi(_log logr.Logger) error {

	_log.V(3).Info("use etcd web ui", "state", config.UseETCDWebUi)

	if config.UseETCDWebUi {

		_log.V(0).Info("booting etcd web ui")
		// setting the default seperator
		var err error

		// start session management
		_log.V(3).Info("starting the session manager")
		config.Sessionmgr, err = session.NewManager("memory", "_etcdui_session", 86400)
		if err != nil {
			_log.Error(err, "error starting session manager")
			return err
		}

		// session garbage collection
		_log.V(3).Info("initializing garbage collection for sessions")
		time.AfterFunc(86400*time.Second, func() {

			_log.V(5).Info("starting garbage collection")
			config.Sessionmgr.GC()
		})

		// starting http service for webui hosting and
		// api functions
		go func() {
			if err = (&http.Server{
				Addr:    fmt.Sprintf("%s:%d", config.ETCDUIHost, config.ETCDUIPort),
				Handler: hdnlrs.GetHandler(httpHandlers),
			}).ListenAndServe(); err != nil {
				_log.Error(err, "error keeping up http server")
				os.Exit(1)
			}
		}()
	}

	return nil
}
