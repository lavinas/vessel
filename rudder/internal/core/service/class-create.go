package service

import (
	"errors"
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
		return c.setError(request, err, dto.StatusBadRequest, err.Error())
	}
	tx, err := c.Repo.Begin("")
	if err != nil {
		return c.setError(request, err, dto.StatusInternalServerError, dto.ErrInternalGeneric)
	}
	defer c.Repo.Rollback(tx)
	class, respError := c.create(request, tx)
	if respError != nil {
		return respError
	}
	if err := c.Repo.Commit(tx); err != nil {
		return c.setError(request, err, dto.StatusInternalServerError, dto.ErrInternalGeneric)
	}
	c.LogOk(request)
	return dto.NewClassCreateResponse(dto.StatusSuccess, "ok", class.ID, request.Name, request.Description, time.Now().Format(time.DateTime))
}

// create is a method that creates a class
func (c *ClassCreate) create(request *dto.ClassCreateRequest, tx interface{}) (*entity.Class, *dto.ClassCreateResponse) {
	class := entity.NewClass(c.Base.Repo)
	if dup, err := class.CheckDuplicity(request.Name, tx); err != nil {
		return nil, c.setError(request, err, dto.StatusInternalServerError, dto.ErrInternalGeneric)
	} else if dup {
		err := errors.New(dto.ErrClassCreateRequestDuplicated)
		return nil, c.setError(request, err, dto.StatusBadRequest, err.Error())
	}
	if err := class.Create(request.Name, request.Description, tx); err != nil {
		return nil, c.setError(request, err, dto.StatusInternalServerError, dto.ErrInternalGeneric)
	}
	return class, nil
}

// setError is a method that sets an error
func (c *ClassCreate) setError(request *dto.ClassCreateRequest, err error, status, message string) *dto.ClassCreateResponse {
	c.LogError(request, err)
	return dto.NewClassCreateResponse(status, message, 0, "", "", "")
}
