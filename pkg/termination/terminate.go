package termination

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// channel used to receive the SIGs from the notifier
var _SIGCHAN = make(chan os.Signal, 2)

// listening to the following sigterms
var shutdownSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
}

// won't be used, but will throw a panic when closed twice
var onlyOneSignalHandler = make(chan struct{})

// collection of termination functions
var terminationFunctions []func()

// configured functions will be executed before shutdown, when
// signal comes in once. when comes in twice, and shutdown is not
// finished, the service will be killed anyways and throw an error
func HandleTermination(terminationFncs ...func()) context.Context {
	terminationFunctions = terminationFncs

	close(onlyOneSignalHandler)

	ctx, cancel := context.WithCancel(context.Background())

	// register the listener
	signal.Notify(_SIGCHAN, shutdownSignals...)

	go func() {
		sig := <-_SIGCHAN

		go func() {

			logging.SLog.Info("caught termination request, running termination processes", "signal", sig.String())

			var termWaitGroup sync.WaitGroup
			termWaitGroup.Add(len(terminationFunctions))

			for i := range terminationFunctions {

				go func() {
					terminationFunctions[i]()
					termWaitGroup.Done()
				}()
			}

			termWaitGroup.Wait()
			logging.SLog.Info("finished termination processes")
			cancel()

		}()

		// receiving another
		sig = <-_SIGCHAN
		os.Exit(1) // second signal. Exit directly with error.

	}()

	return ctx
}

// add handler to termination functions
func AddHandlers(fncs ...func()) {
	terminationFunctions = append(terminationFunctions, fncs...)
}

// internally shut down the service with an sigterm
func Shutdown() {
	_SIGCHAN <- syscall.SIGTERM
}
