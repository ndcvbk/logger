package logger

import (
    "github.com/sirupsen/logrus"
)

type Level logrus.Level

const (
    TraceLevel = Level(logrus.TraceLevel)
    DebugLevel = Level(logrus.DebugLevel)
    InfoLevel = Level(logrus.InfoLevel)
    WarnLevel = Level(logrus.WarnLevel)
    ErrorLevel = Level(logrus.ErrorLevel)
    FatalLevel = Level(logrus.FatalLevel)
)

type ILogger interface {
    // Log with loglevel trace
    Trace(message string, args ...interface{})

    // Log with loglevel debug
    Debug(message string, args ...interface{})

    // Log with loglevel info
    Info(message string, args ...interface{})
    
    // Log with loglevel warn
    Warn(message string, args ...interface{})
    
    // Log with loglevel error
    Error(message string, args ...interface{})
    
    // Log with loglevel fatal
    Fatal(message string, args ...interface{})

    GetLevel() Level

    IsLevelEnabled(level Level) bool
}
