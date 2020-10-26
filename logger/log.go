package logger

// Log is the logging interface
type Log interface {
	Info(string)
}

var registeredLog Log

// Format is the common log format
const Format = "%v: INFO: %s\n"

// RegisterLog is used for registering
func RegisterLog(log Log) {
	registeredLog = log
}

// Info is used to print the message
func Info(message string) {
	registeredLog.Info(message)
}
