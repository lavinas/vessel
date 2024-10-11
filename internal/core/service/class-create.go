package service

import (
	"time"

	"github.com/lavinas/vessel/internal/core/entity"
	"github.com/lavinas/vessel/internal/dto"
	"github.com/lavinas/vessel/internal/port"
)

// ClassCreate represents the create class service
type ClassCreate struct {
	Base
}

// NewClass creates a new Class
func NewClassCreate(repo port.Repository, logger port.Logger, config port.Config) *ClassCreate {
	return &ClassCreate{
		Base: Base{
			Repo:   repo,
			Logger: logger,
			Config: config,
		},
	}
}

// Run is a method that runs the service
func (c *ClassCreate) Run(request *dto.ClassCreateRequest) *dto.ClassCreateResponse {
	if err := request.Validate(); err != nil {
		c.LogError(request, err)
		return dto.NewClassCreateResponse(dto.StatusBadRequest, err.Error(), 0, "", "", "")
	}
	tx, err := c.Repo.Begin("assets")
	if err != nil {
		c.LogError(request, err)
		return dto.NewClassCreateResponse(dto.StatusInternalServerError, dto.ErrInternalGeneric, 0, "", "", "")
	}
	defer c.Repo.Rollback(tx)
	class := entity.NewClass(c.Base.Repo)
	if err := class.Create(request.Name, request.Description, tx); err != nil {
		c.LogError(request, err)
		return dto.NewClassCreateResponse(dto.StatusInternalServerError, dto.ErrInternalGeneric, 0, "", "", "")
	}
	if err := c.Repo.Commit(tx); err != nil {
		c.LogError(request, err)
		return dto.NewClassCreateResponse(dto.StatusInternalServerError, dto.ErrInternalGeneric, 0, "", "", "")
	}
	c.LogOk(request)
	return dto.NewClassCreateResponse(dto.StatusSuccess, "ok", class.ID, request.Name, request.Description, time.Now().Format(time.DateTime))
}

