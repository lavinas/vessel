package service

import (
	"github.com/lavinas/vessel/internal/port"
)

const (
	okMessage  = "%s OK: %s\n"
	errMessage = "%s ERROR %s: %s\n"
)

// Base represents the base service
type Base struct {
	Repo   port.Repository
	Logger port.Logger
	Config port.Config
}

// LogOk is a method that logs an OK message
func (s *Base) LogOk(request port.Request) {
	s.Logger.Printf(okMessage, request.Action, request.ToJson())
}

// LogError is a method that logs an error message
func (s *Base) LogError(request port.Request, err error) {
	s.Logger.Printf(errMessage, request.Action, request.Action(), err.Error())
}
