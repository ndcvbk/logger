package logger

import (
    "context"
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
    Trace(ctx context.Context, message string, args ...interface{})

    // Log with loglevel debug
    Debug(ctx context.Context, message string, args ...interface{})

    // Log with loglevel info
    Info(ctx context.Context, message string, args ...interface{})
    
    // Log with loglevel warn
    Warn(ctx context.Context, message string, args ...interface{})
    
    // Log with loglevel error
    Error(ctx context.Context, message string, args ...interface{})
    
    // Log with loglevel fatal
    Fatal(ctx context.Context, message string, args ...interface{})

    GetLevel() Level

    IsLevelEnabled(level Level) bool
}
