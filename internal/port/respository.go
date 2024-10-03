package port

// Repository represents the repository port
type Repository interface {
	Close() error
	Ping() error
	Reconnect() error
	Check() error
	QueryRow(query string, args ...interface{}) interface{}
	Scan(row interface{}, args ...any) error
}
