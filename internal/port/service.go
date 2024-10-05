package port

// Service represents the service port
type Service interface {
	Run(request Request) Response
}
