package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
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
	// lock is a global mutex lock to gain control of logrus.SetOutput
	lock = sync.RWMutex{}
)

// Constructor
func GetInstance(logLevelString string, jsonFormat bool) ILogger {
	if instance == nil {
		createLoggerOnce.Do(func() {
			instance = &logger{}

			logrusLogger := logrus.New()
			logrusLogger.SetOutput(os.Stdout)
			if jsonFormat {
				logrusLogger.SetFormatter(&logrus.JSONFormatter{})
			}

			logLevel, err := logrus.ParseLevel(logLevelString)
			if err != nil {
				logrusLogger.Warnf("Cannot set log level, will fall back to default level %s", defaultLogLevel)
				logLevel = defaultLogLevel
			}

			logrusLogger.SetLevel(logLevel)
			instance.Logger = logrusLogger
		})
	} else {
		instance.Warn(nil, "The instance already exists. Will ignore passed log level [%s]. Returning logger instance with logLevel [%s].", logLevelString, logrus.Level(instance.GetLevel()))
	}

	return instance
}

func (l *logger) Info(ctx context.Context, message string, args ...interface{}) {
	if l.IsLevelEnabled(InfoLevel) {
		lock.Lock()
		l.SetOutput(os.Stdout)
		entry, ok := l.createEntry(ctx)
		if ok {
			entry.Infof(message, parseArgs(args...)...)
		}
		lock.Unlock()
	}
}

func (l *logger) Trace(ctx context.Context, message string, args ...interface{}) {
	if l.IsLevelEnabled(TraceLevel) {
		lock.Lock()
		l.SetOutput(os.Stdout)
		entry, ok := l.createEntry(ctx)
		if ok {
			entry.Tracef(message, parseArgs(args...)...)
		}
		lock.Unlock()
	}
}

func (l *logger) Debug(ctx context.Context, message string, args ...interface{}) {
	if l.IsLevelEnabled(DebugLevel) {
		lock.Lock()
		l.SetOutput(os.Stdout)
		entry, ok := l.createEntry(ctx)
		if ok {
			entry.Debugf(message, parseArgs(args...)...)
		}
		lock.Unlock()
	}
}

func (l *logger) Warn(ctx context.Context, message string, args ...interface{}) {
	if l.IsLevelEnabled(WarnLevel) {
		lock.Lock()
		l.SetOutput(os.Stdout)
		entry, ok := l.createEntry(ctx)
		if ok {
			entry.Warnf(message, parseArgs(args...)...)
		}
		lock.Unlock()
	}
}

func (l *logger) Error(ctx context.Context, message string, args ...interface{}) {
	if l.IsLevelEnabled(ErrorLevel) {
		lock.Lock()
		l.SetOutput(os.Stderr)
		entry, ok := l.createEntry(ctx)
		if ok {
			entry.Errorf(message, parseArgs(args...)...)
		}
		lock.Unlock()
	}
}

func (l *logger) Fatal(ctx context.Context, message string, args ...interface{}) {
	if l.IsLevelEnabled(FatalLevel) {
		lock.Lock()
		l.SetOutput(os.Stderr)
		entry, ok := l.createEntry(ctx)
		if ok {
			entry.Fatalf(message, parseArgs(args...)...)
		}
		lock.Unlock()
	}
}

func (l *logger) GetLevel() Level {
	return Level(l.Logger.GetLevel())
}

func (l *logger) IsLevelEnabled(level Level) bool {
	return l.Logger.IsLevelEnabled(logrus.Level(level))
}

func (l *logger) createEntry(ctx context.Context) (*logrus.Entry, bool) {
	entry := logrus.
		NewEntry(l.Logger).
		WithFields(logrus.Fields{
			"frame": getFrameInfo(),
		})
	if ctx != nil {
		requestId, ok := ctx.Value(key).(string)
		if !ok {
			return nil, false
		}
		entry = entry.WithField("requestId", requestId)
	}
	return entry, true
}

type frameInfo struct {
	Function string `json:"function"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}

// Function to retrieve information about the log-calling function
func getFrameInfo() frameInfo {
	// target a the frame just outside the call stack of this logger
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

	return frameInfo{
		Function: frame.Func.Name(),
		File:     frame.File,
		Line:     frame.Line,
	}
}

// parseArgs checks if the args have a value and replaces the value with nil (string)
func parseArgs(args ...interface{}) []interface{} {
	for key, value := range args {
		reflectValue := reflect.ValueOf(value)

		if isPointer(reflectValue) && reflectValue.IsNil() {
			switch value.(type) {
			case *int8, *uint8, *int16, *uint16, *int32, *uint32, *int64, *uint64, *int, *uint, *uintptr, *float32, *float64, *complex64, *complex128:
				args[key] = 0
			default:
				args[key] = "nil"
			}
		}

		args[key] = actualValue(args[key])

	}

	return args
}

// isPointer returns true if the interface is a pointer
func isPointer(value reflect.Value) bool {
	return value.Kind() == reflect.Ptr
}

// actualValue returns the actual value, no pointers
func actualValue(value interface{}) interface{} {
	if isPointer(reflect.ValueOf(value)) {
		switch value.(type) {
		case *string:
			text := value.(*string)
			return *text
		case *int8:
			number := value.(*int8)
			return *number
		case *uint8:
			number := value.(*uint8)
			return *number
		case *int16:
			number := value.(*int16)
			return *number
		case *uint16:
			number := value.(*uint16)
			return *number
		case *int32:
			number := value.(*int32)
			return *number
		case *uint32:
			number := value.(*uint32)
			return *number
		case *int64:
			number := value.(*int64)
			return *number
		case *uint64:
			number := value.(*uint64)
			return *number
		case *int:
			number := value.(*int)
			return *number
		case *uint:
			number := value.(*uint)
			return *number
		case *uintptr:
			number := value.(*uintptr)
			return *number
		case *float32:
			number := value.(*float32)
			return *number
		case *float64:
			number := value.(*float64)
			return *number
		case *complex64:
			number := value.(*complex64)
			return *number
		case *complex128:
			number := value.(*complex128)
			return *number
		case *bool:
			boolean := value.(*bool)
			return *boolean
		}
	}

	return value
}

const requestIdHeader = "x-request-id"

type keyType int

const key keyType = 0

func GetIdFromRequest(req *http.Request) string {
	return req.Header.Get(requestIdHeader)
}

func NewContext(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, key, requestId)
}
