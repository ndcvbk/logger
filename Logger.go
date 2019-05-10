package logger

import (
    "fmt"
    "github.com/logmatic/logmatic-go"
    "github.com/sirupsen/logrus"
    "sync"
)

// Logger ...
type Logger struct {
    logEntry *logrus.Entry
    singleton *Logger
}

const (
    defaultLogLevel =   logrus.WarnLevel

    cantSetLogLevel =   "Cannot set log level, will fall back to default level %s"

    interfaze =         "interface"
    function =          "function"
)

var (
    createLoggerOnce sync.Once
)

// Constructor
func New(logLevelAsString string) *Logger {
    this := Logger{}

    createLoggerOnce.Do(func() {
        logger := logrus.New()
        logger.SetFormatter(&logmatic.JSONFormatter{})
        logLevel, err := logrus.ParseLevel(logLevelAsString)
        if err != nil {
            logger.Warn(fmt.Printf(cantSetLogLevel, defaultLogLevel))
            logLevel = defaultLogLevel
        }
        logger.SetLevel(logLevel)
        this.singleton = &Logger{ logEntry: logrus.NewEntry(logger)}
    })
    return this.singleton
}

// Info ...
func (l *Logger) Info(interfaceName string, functionName string, message string) {
    l.createEntry(interfaceName, functionName).Info(message)
}

// Trace ...
func (l *Logger) Trace(interfaceName string, functionName string, message string) {
    l.createEntry(interfaceName, functionName).Trace(message)
}

// Debug ...
func (l *Logger) Debug(interfaceName string, functionName string, message string) {
    l.createEntry(interfaceName, functionName).Debug(message)
}

// Warn ...
func (l *Logger) Warn(interfaceName string, functionName string, message string) {
    l.createEntry(interfaceName, functionName).Warn(message)
}

// Error ...
func (l *Logger) Error(interfaceName string, functionName string, message string) {
    l.createEntry(interfaceName, functionName).Error(message)
}

// Fatal ...
func (l *Logger) Fatal(interfaceName string, functionName string, message string) {
    l.createEntry(interfaceName, functionName).Fatal(message)
}

// Generic
func (l *Logger) createEntry(interfaceName string, functionName string) *logrus.Entry {
    return l.logEntry.WithFields(logrus.Fields{interfaze: interfaceName, function: functionName})
}