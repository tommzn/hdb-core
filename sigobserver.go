package core

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// newSigObserver returns a new observer for OS signals.
func newSigObserver() *sigObserver {
	return &sigObserver{
		osSignalChan: make(chan os.Signal),
	}
}

// newSigObserverWithChan returns a new observer for OS signals
// on passed signal channnel.
func newSigObserverWithChan(osSignalChan chan os.Signal) *sigObserver {
	return &sigObserver{
		osSignalChan: osSignalChan,
	}
}

// observe blocks until one of following OS signals occurs.
// SIGTERM, SIGSTOP, SIGKILL
func (observer *sigObserver) observe() {

	signal.Notify(observer.osSignalChan, syscall.SIGTERM)
	signal.Notify(observer.osSignalChan, syscall.SIGSTOP)
	signal.Notify(observer.osSignalChan, syscall.SIGKILL)

	log.Println("[SigObserver] Observing os signals...")
	sig := <-observer.osSignalChan
	log.Printf("[SigObserver] Caught OS signal: %+v\n", sig)
}
