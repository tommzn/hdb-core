package core

import (
	"context"
	"os"
	"sync"
)

// Minion is a worker which observes execution of a child and
// OS signals for graceful shutdowns.
type Minion struct {

	// sigObserver listens for different IS signals.
	sigObserver *sigObserver

	// Wg is used to recognize if child task has finished.
	wg *sync.WaitGroup

	// Ctx, current context.
	ctx context.Context

	// CancelFunc for current context.
	cancelFunc func()

	// ErrorChan is used to fetch errors from child tasks.
	errorChan chan error

	// Child which should be executed
	child Runable
}

// SigObserver listen to OS signals.
type sigObserver struct {

	// OsSignalChan is used to observe signals.
	osSignalChan chan os.Signal
}
