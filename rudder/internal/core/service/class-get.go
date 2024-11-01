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
		return c.setError(request, err, dto.StatusBadRequest, err.Error())
	}
	tx, err := c.Repo.Begin("")
	if err != nil {
		return c.setError(request, err, dto.StatusInternalServerError, dto.ErrInternalGeneric)
	}
	defer c.Repo.Rollback(tx)
	cl := entity.NewClass(c.Base.Repo)
	if err := cl.GetByID(request.ID, tx); err != nil {
		return c.setError(request, err, dto.StatusInternalServerError, dto.ErrInternalGeneric)
	}
	if !cl.IsLoaded() {
		err := errors.New(dto.ErrClassGetRequestNotFound)
		return c.setError(request, err, dto.StatusNotFound, err.Error())
	}
	c.LogOk(request)
	return dto.NewClassGetResponse(dto.StatusSuccess, "", cl.ID, cl.Name, cl.Description, cl.CreatedAt.Format(time.DateTime))
}

// setError is a method that sets an error
func (c *ClassGet) setError(request *dto.ClassGetRequest, err error, status, message string) *dto.ClassGetResponse {
	c.LogError(request, err)
	return dto.NewClassGetResponse(status, message, 0, "", "", "")
}
