package logger

import (
    "fmt"
    "github.com/logmatic/logmatic-go"
    "github.com/sirupsen/logrus"
    "reflect"
    "runtime"
    "sync"
)

type logger struct {
    *logrus.Logger
}

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
            logrusLogger.SetFormatter(&logmatic.JSONFormatter{})

            logLevel, err := logrus.ParseLevel(logLevelString)
            if err != nil {
                logrusLogger.Warnf("Cannot set log level, will fall back to default level %s", defaultLogLevel)
                logLevel = defaultLogLevel
            }

            logrusLogger.SetLevel(logLevel)
            instance.Logger = logrusLogger
        })
    } else {
        instance.Trace("The instance already exists. Will ignore passed log level [%s]. Returning logger instance with logLevel [%s].", logLevelString, logrus.Level(instance.GetLevel()))
    }

    return instance
}

func (l *logger) Info(message string, args ...interface{}) {
    if l.IsLevelEnabled(InfoLevel) {
        l.createEntry().Infof(message, parseArgs(args)...)
    }
}

func (l *logger) Trace(message string, args ...interface{}) {
    if l.IsLevelEnabled(TraceLevel) {
        l.createEntry().Tracef(message, parseArgs(args...)...)
    }
}

func (l *logger) Debug(message string, args ...interface{}) {
    if l.IsLevelEnabled(DebugLevel) {
        l.createEntry().Debugf(message, parseArgs(args...)...)
    }
}

func (l *logger) Warn(message string, args ...interface{}) {
    if l.IsLevelEnabled(WarnLevel) {
        l.createEntry().Warnf(message, parseArgs(args...)...)
    }
}

func (l *logger) Error(message string, args ...interface{}) {
    if l.IsLevelEnabled(ErrorLevel) {
        l.createEntry().Errorf(message, parseArgs(args...)...)
    }
}

func (l *logger) Fatal(message string, args ...interface{}) {
    if l.IsLevelEnabled(FatalLevel) {
        l.createEntry().Fatalf(message, parseArgs(args...)...)
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
    return logrus.NewEntry(l.Logger).WithFields(logrus.Fields{"frame": getFrameInfo()})
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

// parseArgs checks if the args have a value and replaces the value with nil (string)
func parseArgs(args ...interface{}) []interface{} {
    for key, value := range args {
        reflectValue := reflect.ValueOf(value)

        if isPointer(reflectValue) && reflectValue.IsNil() {
            switch value.(type) {
            case *string:
              args[key] = "nil"
            default:
              args[key] = 0
            }
        }
    }
    return args
}

// isPointer returns true if the interface is a pointer
func isPointer(value reflect.Value) bool {
    return value.Kind() == reflect.Ptr
}

