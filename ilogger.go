package logger

// ILogger ...
type ILogger interface {
    Init()
    Info(interfaceName string, functionName string, message string)
    Trace(interfaceName string, functionName string, message string)
    Debug(interfaceName string, functionName string, message string)
    Warn(interfaceName string, functionName string, message string)
    Error(interfaceName string, functionName string, message string)
    Fatal(interfaceName string, functionName string, message string)
}