package port

// Config represents the configuration interface for the service
type Config interface {
	Get(key string) string
}
