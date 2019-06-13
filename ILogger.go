package logger

type ILogger interface {
    // Log with loglevel info
    Info(message string, args ...interface{})
    
    // Log with loglevel trace
    Trace(message string, args ...interface{})
    
    // Log with loglevel debug
    Debug(message string, args ...interface{})
    
    // Log with loglevel warn
    Warn(message string, args ...interface{})
    
    // Log with loglevel error
    Error(message string, args ...interface{})
    
    // Log with loglevel fatal
    Fatal(message string, args ...interface{})
}
