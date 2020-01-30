package logger

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

type LoggerTestSuite struct {
	suite.Suite
}

func TestSuiteForLogger(t *testing.T) {
	suite.Run(t, &LoggerTestSuite{})
}

// This test doesn't assert anything at the moment, it's just a quick way to call these functions
func (suite *LoggerTestSuite) Test() {
	ctx := NewContext(context.Background(), "test")
	logger := GetInstance("trace", false)

	logger.Trace(nil, "testMessage")
	logger.Trace(nil, "test %s", "test")
	logger.Trace(nil, "test %s")
	logger.Trace(nil, "test %d", 1, "extra")
	logger.Trace(nil, "test %d", nil)

	var decimal *int
	decimal = new(int)
	*decimal = 12
	var nul *string
	var boolean *bool
	boolean = new(bool)
	*boolean = true

	logger.Trace(nil, "test %d %s %s %t", decimal, nul, "testvar", boolean)

	logger.Debug(nil, "testMessage")
	logger.Debug(nil, "test %s", "test")
	logger.Debug(nil, "test %s")
	logger.Debug(nil, "test %d", 1, "extra")

	logger.Info(nil, "testMessage")
	logger.Info(nil, "test %s", "test")
	logger.Info(nil, "test %s")
	logger.Info(nil, "test %d", 1, "extra")

	logger.Warn(nil, "testMessage")
	logger.Warn(nil, "test %s", "test")
	logger.Warn(nil, "test %s")
	logger.Warn(nil, "test %d", 1, "extra")

	logger.Error(nil, "testMessage")
	logger.Error(nil, "test %s", "test")
	logger.Error(nil, "test %s")
	logger.Error(nil, "test %d", 1, "extra")

	logger.Trace(ctx, "testMessage")
	logger.Debug(ctx, "testMessage")
	logger.Info(ctx, "testMessage")
	logger.Warn(ctx, "testMessage")
	logger.Error(ctx, "testMessage")

	logger.Trace(context.Background(), "testMessage")
	logger.Debug(context.Background(), "testMessage")
	logger.Info(context.Background(), "testMessage")
	logger.Warn(context.Background(), "testMessage")
	logger.Error(context.Background(), "testMessage")
}
