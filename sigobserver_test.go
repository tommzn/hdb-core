package core

import (
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type SigObserverTestSuite struct {
	suite.Suite
}

func TestSigObserverTestSuite(t *testing.T) {
	suite.Run(t, new(SigObserverTestSuite))
}

func (suite *SigObserverTestSuite) TestObserveSignals() {

	testChan := make(chan os.Signal)
	successChan := make(chan bool, 1)

	observer := newSigObserverWithChan(testChan)

	go notifyOsSignal(testChan)
	go waitForOsSignal(observer, successChan)

	time.Sleep(4 * time.Second)
	suite.Len(successChan, 1)
}

func notifyOsSignal(osSignalChan chan os.Signal) {
	time.Sleep(3 * time.Second)
	osSignalChan <- syscall.SIGTERM
}

func waitForOsSignal(observer *sigObserver, successChan chan bool) {
	observer.observe()
	successChan <- true
}
