package port

// Request represents the request interface
type Request interface {
	Action() string
	ToJson() string
	Validate() error
}

// Response represents the response interface
type Response interface {
	ToJson() string
	String() string
}
