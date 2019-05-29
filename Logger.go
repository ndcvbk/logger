package logger

import (
    "fmt"
    "github.com/logmatic/logmatic-go"
    "github.com/sirupsen/logrus"
    "runtime"
    "sync"
)

// Logger ...
type Logger struct {
    logEntry *logrus.Entry
    logLevel logrus.Level
}

const (
    cantSetLogLevel       = "Cannot set log level, will fall back to default level %s"
    instanceAlreadyExists = "The instance already exists. Will ignore passed log level [%s]. Returning logger instance with logLevel [%s]."
    frame                 = "frame"
)

var (
    createLoggerOnce sync.Once
    defaultLogLevel  = logrus.WarnLevel
    instance         *Logger
)

// Constructor
func GetInstance(logLevelAsString string) *Logger {

    if instance == nil {
        createLoggerOnce.Do(func() {
            instance = &Logger{}

            logger := logrus.New()
            logger.SetFormatter(&logmatic.JSONFormatter{})
            logLevel, err := logrus.ParseLevel(logLevelAsString)

            if err != nil {
                logger.Warn(cantSetLogLevel, defaultLogLevel)
                logLevel = defaultLogLevel
            }

            logger.SetLevel(logLevel)
            entry := logrus.NewEntry(logger)
            entry.Level = logLevel

            instance.logLevel = logLevel
            instance.logEntry = entry
        })
    } else {
        instance.Trace(instanceAlreadyExists, logLevelAsString, instance.logLevel)
    }
    return instance
}

// Info ...
func (l *Logger) Info(message string, variadic ...interface{}) {
    if l.isLogLevelEnabled(logrus.InfoLevel) {
        message = formatMessage(message, variadic...)
        l.createEntry().Info(message)
    }
}

// Trace ...
func (l *Logger) Trace(message string, variadic ...interface{}) {
    if l.isLogLevelEnabled(logrus.TraceLevel) {
        message = formatMessage(message, variadic...)
        l.createEntry().Trace(message)
    }
}

// Debug ...
func (l *Logger) Debug(message string, variadic ...interface{}) {
    if l.isLogLevelEnabled(logrus.DebugLevel) {
        message = formatMessage(message, variadic...)
        l.createEntry().Debug(message)
    }
}

// Warn ...
func (l *Logger) Warn(message string, variadic ...interface{}) {
    if l.isLogLevelEnabled(logrus.WarnLevel) {
        message = formatMessage(message, variadic...)
        l.createEntry().Warn(message)
    }
}

// Error ...
func (l *Logger) Error(message string, variadic ...interface{}) {
    if l.isLogLevelEnabled(logrus.ErrorLevel) {
        message = formatMessage(message, variadic...)
        l.createEntry().Error(message)
    }
}

// Fatal ...
func (l *Logger) Fatal(message string, variadic ...interface{}) {
    if l.isLogLevelEnabled(logrus.FatalLevel) {
        message = formatMessage(message, variadic...)
        l.createEntry().Fatal(message)
    }
}

// Generic helper function
func (l *Logger) createEntry() *logrus.Entry {
    frameInfo := getFrameInfo()
    return l.logEntry.WithFields(logrus.Fields{frame: frameInfo})
}

// Check if logging at this level is enabled.
func (l *Logger) isLogLevelEnabled(level logrus.Level) bool {
    shouldLogAtLevel := logrus.AllLevels[l.logLevel]
    entryLevel := logrus.AllLevels[level]
    return shouldLogAtLevel >= entryLevel
}

// Function to retrieve information about the log-calling function
func getFrameInfo() string {
    // We need the frame at index 3, since we never want runtime.Callers, getFrame, createEntry and Info|Debug|Trace|etc
    targetFrameIndex := 4

    programCounters := make([]uintptr, targetFrameIndex+1)
    n := runtime.Callers(0, programCounters)

    frame := runtime.Frame{Function: "unknown"}
    if n > 0 {
        frames := runtime.CallersFrames(programCounters[:n])
        for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
            var frameCandidate runtime.Frame
            frameCandidate, more = frames.Next()
            if frameIndex == targetFrameIndex {
                frame = frameCandidate
            }
        }
    }

    message := fmt.Sprintf("%s, %s #%v", frame.Func.Name(), frame.File, frame.Line)

    return message
}

// Helps formatting the message if multiple vars have been passed
func formatMessage(message string, variadic ...interface{}) string {
    if variadic != nil && len(variadic) > 0 {
        return fmt.Sprintf(message, variadic...)
    } else {
        return message
    }
}
