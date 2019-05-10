package logger

// ILogger ...
type ILogger interface {
    // Log with loglevel info
    Info(interfaceName string, functionName string, message string)
    // Log with loglevel trace
    Trace(interfaceName string, functionName string, message string)
    // Log with loglevel debug
    Debug(interfaceName string, functionName string, message string)
    // Log with loglevel warn
    Warn(interfaceName string, functionName string, message string)
    // Log with loglevel error
    Error(interfaceName string, functionName string, message string)
    // Log with loglevel fatal
    Fatal(interfaceName string, functionName string, message string)
}