package core

import (
	"context"
	"errors"
	"sync"
	"time"
)

type testChild struct {
	shouldEndWithError bool
	finished, stopped  bool
	counter            int
	stopChan           chan bool
}

func childForTest() *testChild {
	return &testChild{
		shouldEndWithError: false,
		finished:           false,
		stopped:            false,
		counter:            0,
		stopChan:           make(chan bool, 1),
	}
}

func (mock *testChild) Run(ctx context.Context, wg *sync.WaitGroup) error {

	defer wg.Done()

	timeout := time.After(3 * time.Second)
	tick := time.Tick(400 * time.Millisecond)
	for {
		select {
		case <-timeout:
			mock.finished = true
			return mock.getReturnError()
		case <-tick:
			mock.counter++
		case <-ctx.Done():
			mock.stopped = true
			return mock.getReturnError()
		}
	}
}

func (mock *testChild) getReturnError() error {
	if mock.shouldEndWithError {
		return errors.New("Error occured!")
	}
	return nil
}
