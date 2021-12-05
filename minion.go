package core

import (
	"context"
	"errors"
	"sync"
)

// NewMinion returns a new runable minion with given child.
func NewMinion(child Runable) *Minion {
	return &Minion{
		sigObserver: newSigObserver(),
		wg:          &sync.WaitGroup{},
		errorChan:   make(chan error, 1),
		child:       child,
	}
}

// Run calls run on currentchild and observe os signals and manual stop.
func (minion *Minion) Run(ctx context.Context) error {

	if minion.child == nil {
		return errors.New("Missing executable!")
	}

	minion.withCancel(ctx)
	minion.wg.Add(1)
	go func() {
		if err := minion.child.Run(minion.ctx, minion.wg); err != nil {
			minion.errorChan <- err
		}

	}()

	go minion.observeOsSignals()
	return minion.waitForChild()
}

// WithCancel generates context with concel func.
func (minion *Minion) withCancel(ctx context.Context) {
	minion.ctx, minion.cancelFunc = context.WithCancel(ctx)
}

// WaitForChild will return, with error if occurs, after child has
// finished execution.
func (minion *Minion) waitForChild() error {
	minion.wg.Wait()
	if len(minion.errorChan) > 0 {
		return <-minion.errorChan
	}
	return nil
}

// ObserveOsSignals will use it's internal observer to handle
// OS signals.
func (minion *Minion) observeOsSignals() {
	minion.sigObserver.observe()
	minion.cancelFunc()
	minion.wg.Wait()
}

// Stop can be called to manual stop execution of a minion.
func (minion *Minion) Stop() {
	minion.cancelFunc()
}
