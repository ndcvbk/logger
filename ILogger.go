package logger

// ILogger ...
type ILogger interface {
    // Log with loglevel info
    Info(message string)
    // Log with loglevel trace
    Trace(message string)
    // Log with loglevel debug
    Debug(message string)
    // Log with loglevel warn
    Warn(message string)
    // Log with loglevel error
    Error(message string)
    // Log with loglevel fatal
    Fatal(message string)
}