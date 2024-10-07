package port

// Logger represents the logger interface for the service
type Logger interface {
	Printf(format string, v ...interface{})
}
