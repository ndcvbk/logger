package logger

import (
	"github.com/ndcvbk/logger/formatter"
    "fmt"
    "github.com/sirupsen/logrus"
    "runtime"
    "sync"
)

type logger struct {
    *logrus.Logger
}

const (
    cantSetLogLevel       = "Cannot set log level, will fall back to default level %s"
    instanceAlreadyExists = "The instance already exists. Will ignore passed log level [%s]. Returning logger instance with logLevel [%s]."
    frame                 = "frame"
)

var (
    createLoggerOnce sync.Once
    defaultLogLevel  = logrus.WarnLevel
    instance         *logger
)

// Constructor
func GetInstance(logLevelString string) ILogger {
    if instance == nil {
        createLoggerOnce.Do(func() {
            instance = &logger{}

            logrusLogger := logrus.New()
            logrusLogger.SetFormatter(&formatter.Formatter{})

            logLevel, err := logrus.ParseLevel(logLevelString)
            if err != nil {
                logrusLogger.Warnf(cantSetLogLevel, defaultLogLevel)
                logLevel = defaultLogLevel
            }

            logrusLogger.SetLevel(logLevel)
            instance.Logger = logrusLogger
        })
    } else {
        instance.Trace(instanceAlreadyExists, logLevelString, logrus.Level(instance.GetLevel()))
    }

    return instance
}

func (l *logger) Info(message string, args ...interface{}) {
    if l.IsLevelEnabled(InfoLevel) {
        l.createEntry().Infof(message, args...)
    }
}

func (l *logger) Trace(message string, args ...interface{}) {
    if l.IsLevelEnabled(TraceLevel) {
        l.createEntry().Tracef(message, args...)
    }
}

func (l *logger) Debug(message string, args ...interface{}) {
    if l.IsLevelEnabled(DebugLevel) {
        l.createEntry().Debugf(message, args...)
    }
}

func (l *logger) Warn(message string, args ...interface{}) {
    if l.IsLevelEnabled(WarnLevel) {
        l.createEntry().Warnf(message, args...)
    }
}

func (l *logger) Error(message string, args ...interface{}) {
    if l.IsLevelEnabled(ErrorLevel) {
        l.createEntry().Errorf(message, args...)
    }
}

func (l *logger) Fatal(message string, args ...interface{}) {
    if l.IsLevelEnabled(FatalLevel) {
        l.createEntry().Fatalf(message, args...)
    }
}

func (l *logger) GetLevel() Level {
    return Level(l.Logger.GetLevel())
}

func (l *logger) IsLevelEnabled(level Level) bool {
    return l.Logger.IsLevelEnabled(logrus.Level(level))
}

// Generic helper function
func (l *logger) createEntry() *logrus.Entry {
    return logrus.NewEntry(l.Logger).WithFields(logrus.Fields{frame: getFrameInfo()})
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
