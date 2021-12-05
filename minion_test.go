package core

import (
	"context"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type MinionTestSuite struct {
	suite.Suite
}

func TestMinionTestSuite(t *testing.T) {
	suite.Run(t, new(MinionTestSuite))
}

func (suite *MinionTestSuite) TestFinishNormal() {

	mock := childForTest()
	minion := NewMinion(mock)
	ctx := context.Background()

	suite.Nil(minion.Run(ctx))
	suite.True(mock.finished)
	suite.False(mock.stopped)
}

func (suite *MinionTestSuite) TestStopProcessing() {

	mock := childForTest()
	minion := NewMinion(mock)
	ctx := context.Background()

	go minion.Run(ctx)
	time.Sleep(2 * time.Second)
	minion.Stop()

	time.Sleep(1 * time.Second)
	suite.True(mock.counter > 1)
	suite.False(mock.finished)
	suite.True(mock.stopped)
}

func (suite *MinionTestSuite) TestWithoutChild() {

	minion := NewMinion(nil)
	ctx := context.Background()

	suite.NotNil(minion.Run(ctx))
}

func (suite *MinionTestSuite) TestProcessingError() {

	mock := childForTest()
	mock.shouldEndWithError = true
	minion := NewMinion(mock)
	ctx := context.Background()

	suite.NotNil(minion.Run(ctx))
	suite.True(mock.finished)
	suite.False(mock.stopped)
}

func (suite *MinionTestSuite) TestStopProcessingOnOsSignals() {

	mock := childForTest()
	osSignalChan := make(chan os.Signal)
	minion := NewMinion(mock)
	minion.sigObserver = newSigObserverWithChan(osSignalChan)
	ctx := context.Background()

	go minion.Run(ctx)
	time.Sleep(2 * time.Second)
	osSignalChan <- syscall.SIGTERM

	time.Sleep(1 * time.Second)
	suite.True(mock.counter > 1)
	suite.False(mock.finished)
	suite.True(mock.stopped)
}
