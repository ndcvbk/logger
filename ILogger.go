package logger

// ILogger ...
type ILogger interface {
    // Log with loglevel info
    Info(message string, variadic ...interface{})
    // Log with loglevel trace
    Trace(message string, variadic ...interface{})
    // Log with loglevel debug
    Debug(message string, variadic ...interface{})
    // Log with loglevel warn
    Warn(message string, variadic ...interface{})
    // Log with loglevel error
    Error(message string, variadic ...interface{})
    // Log with loglevel fatal
    Fatal(message string, variadic ...interface{})
}
