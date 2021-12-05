package core

import (
	"context"
	"sync"
)

// Interface used for executables. Should listen to context cancel
// if graceful shutdown is desired.
type Runable interface {

	// Run is used as an entry point to execute desired busines logic.
	// Should confirm wait group on exit.
	Run(context.Context, *sync.WaitGroup) error
}
