package termination

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

var (
	sigChan       = make(chan os.Signal, 2)
	handlers      = []interface{ Close() error }{}
	started  bool = false
)

type TerminationHandler struct {
	handlers []interface{ Close() error }
	sigChan  chan os.Signal
	started  bool
}

// adds a new handler with the required interface to the handlers
func AddHandlers(t ...interface{ Close() error }) {
	handlers = append(handlers, t...)
}

// this is a terst struct, to test the close caller termination
type Test struct{}

// this is a testfunc, that does nothing but wait 10 seconds before
// finishing the function
func (t *Test) Close() error {
	logging.Default.Debug("running termination shutdown, waiting for 10 seconds")
	time.Sleep(10 * time.Second)
	logging.Default.Debug("terminating service...")
	return nil
}

// configured TerminationHandlers will be executed before shutdown, when
// signal comes in once. when comes in twice, and shutdown is not
// finished, the service will be killed anyways and throw an error
func AsyncHandle(ctx context.Context) {
	if started {
		return
	}

	sigChan = make(chan os.Signal, 2)

	// register the listener, listening to the following sigterms
	signal.Notify(sigChan, []os.Signal{
		os.Interrupt,
		syscall.SIGTERM,
	}...)

	runTerminationHandlers := func() {

		var termWaitGroup sync.WaitGroup
		termWaitGroup.Add(len(handlers))

		for _, handler := range handlers {
			go func() {
				if err := handler.Close(); err != nil {
					logging.Default.Error("error terminating the process", "err", err.Error())
				}
				termWaitGroup.Done()
			}()
		}

		termWaitGroup.Wait()
		logging.Default.Info("finished termination processes")
	}

	go func() {
		select {
		case <-ctx.Done():
			logging.Default.Info("caught termination request, running termination processes", "reason", ctx.Err())
			go runTerminationHandlers()
		case sig := <-sigChan:
			logging.Default.Info("caught termination request, running termination processes", "reason", sig.String())
			go runTerminationHandlers()
		}
	}()

	started = true

	// receiving another
	<-sigChan
	os.Exit(1) // second signal. Exit directly with error.
}
