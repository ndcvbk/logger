package logger

import (
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
    logger := GetInstance("trace")

    logger.Trace("testMessage")
    logger.Trace("test %s", "test")
    logger.Trace("test %s")
    logger.Trace("test %d", 1, "extra")
    logger.Trace("test %d", nil)

    var decimal *int
    decimal = new(int)
    *decimal = 12
    var nul *string
    var boolean *bool
    boolean = new(bool)
    *boolean = true

    logger.Trace("test %d %s %s %t", decimal, nul, "testvar", boolean)

    logger.Debug("testMessage")
    logger.Debug("test %s", "test")
    logger.Debug("test %s")
    logger.Debug("test %d", 1, "extra")

    logger.Info("testMessage")
    logger.Info("test %s", "test")
    logger.Info("test %s")
    logger.Info("test %d", 1, "extra")

    logger.Warn("testMessage")
    logger.Warn("test %s", "test")
    logger.Warn("test %s")
    logger.Warn("test %d", 1, "extra")

    logger.Error("testMessage")
    logger.Error("test %s", "test")
    logger.Error("test %s")
    logger.Error("test %d", 1, "extra")
}
