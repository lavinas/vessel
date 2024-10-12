package service

import (
	"errors"
	"time"

	"github.com/lavinas/vessel/internal/core/entity"
	"github.com/lavinas/vessel/internal/dto"
	"github.com/lavinas/vessel/internal/port"
)

// ClassGet represents the get class service
type ClassGet struct {
	Base
}

// NewClassGet creates a new ClassGet
func NewClassGet(repo port.Repository, logger port.Logger, config port.Config) *ClassGet {
	return &ClassGet{
		Base: Base{
			Repo:   repo,
			Logger: logger,
			Config: config,
		},
	}
}

// Run is a method that runs the service
func (c *ClassGet) Run(request *dto.ClassGetRequest) *dto.ClassGetResponse {
	if err := request.Validate(); err != nil {
		c.LogError(request, err)
		return dto.NewClassGetResponse(dto.StatusBadRequest, err.Error(), 0, "", "", "")
	}
	tx, err := c.Repo.Begin("")
	if err != nil {
		c.LogError(request, err)
		return dto.NewClassGetResponse(dto.StatusInternalServerError, dto.ErrInternalGeneric, 0, "", "", "")
	}
	defer c.Repo.Rollback(tx)
	cl := entity.NewClass(c.Base.Repo)
	if err := cl.GetByID(request.ID, tx); err != nil {
		c.LogError(request, err)
		return dto.NewClassGetResponse(dto.StatusInternalServerError, dto.ErrInternalGeneric, 0, "", "", "")
	}
	if !cl.IsLoaded() {
		c.LogError(request, errors.New(dto.ErrClassCreateRequestDuplicated))
		return dto.NewClassGetResponse(dto.StatusNotFound, dto.ErrClassGetRequestNotFound, 0, "", "", "")
	}
	c.LogOk(request)
	return dto.NewClassGetResponse(dto.StatusSuccess, "", cl.ID, cl.Name, cl.Description, cl.CreatedAt.Format(time.DateTime))
}
