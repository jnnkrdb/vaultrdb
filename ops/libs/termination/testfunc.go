package termination

import (
	"time"

	"github.com/jnnkrdb/vaultrdb/libs/logging"
)

// this is a testfunc, that does nothing but wait 10 seconds before
// finishing the function
func Testfunc() {

	logging.Log.Info("running SIGTERM shutdown, waiting for 10 seconds")

	time.Sleep(10 * time.Second)

	logging.Log.Info("terminating service...")
}
