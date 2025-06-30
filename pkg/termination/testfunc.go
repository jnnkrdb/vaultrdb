package termination

import (
	"time"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// this is a testfunc, that does nothing but wait 10 seconds before
// finishing the function
func Testfunc() {

	logging.Default.Info("running termination shutdown, waiting for 10 seconds")

	time.Sleep(10 * time.Second)

	logging.Default.Info("terminating service...")
}
