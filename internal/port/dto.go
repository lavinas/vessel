package port

// Request represents the request interface
type Request interface {
	Action() string
	ToJson() string
}

// Response represents the response interface
type Response interface {
	ToJson() string
	ToLine() string
}
