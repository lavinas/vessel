package service

import (
	"github.com/lavinas/vessel/internal/dto"
	"github.com/lavinas/vessel/internal/port"
)

// DtoService represents the service that processes the DTO
type DtoService struct {
	Base
}

// NewDtoService creates a new DtoService
func NewDtoService(repo port.Repository, logger port.Logger, config port.Config) *DtoService {
	return &DtoService{
		Base: Base{
			Repo:   repo,
			Logger: logger,
			Config: config,
		},
	}
}

// Run is a method that runs the service
func (s *DtoService) Run(request port.Request) port.Response {
	switch request.Action() {

	case "ClassCreate": // Create a class
		rn := NewClassService(s.Repo, s.Logger, s.Config)
		return rn.Create(request.(*dto.ClassCreateRequest))
	case "ClassGet": // Get a class
		rn := NewClassService(s.Repo, s.Logger, s.Config)
		return rn.Get(request.(*dto.ClassGetRequest))
	default:
		return dto.NewBaseResponse(dto.StatusBadRequest, "invalid action")
	}
}
