package service

import (
	"github.com/lavinas/vessel/internal/dto"
	"github.com/lavinas/vessel/internal/port"
)

// DtoGeneric represents the service that processes the DTO
type DtoGeneric struct {
	Base
}

// NewDtoService creates a new DtoService
func NewDtoGeneric(repo port.Repository, logger port.Logger, config port.Config) *DtoGeneric {
	return &DtoGeneric{
		Base: Base{
			Repo:   repo,
			Logger: logger,
			Config: config,
		},
	}
}

// Run is a method that runs the service
func (s *DtoGeneric) Run(request port.Request) port.Response {
	switch request.Action() {
	case "ClassCreate": // Create a class
		rn := NewClassCreate(s.Repo, s.Logger, s.Config)
		return rn.Run(request.(*dto.ClassCreateRequest))
	case "ClassGet": // Get a class
		rn := NewClassGet(s.Repo, s.Logger, s.Config)
		return rn.Run(request.(*dto.ClassGetRequest))
	case "AssetCreate": // Create an asset
		rn := NewAssetCreate(s.Repo, s.Logger, s.Config)
		return rn.Run(request.(*dto.AssetCreateRequest))
	case "AssetGet": // Get an asset
		rn := NewAssetGet(s.Repo, s.Logger, s.Config)
		return rn.Run(request.(*dto.AssetGetRequest))
	default:
		return dto.NewBaseResponse(dto.StatusBadRequest, "invalid action")
	}
}
