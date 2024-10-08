package service

import (
	"time"

	"github.com/lavinas/vessel/internal/core/entity"
	"github.com/lavinas/vessel/internal/dto"
	"github.com/lavinas/vessel/internal/port"
)

// Class represents the class service
type ClassService struct {
	Base
}

// NewClass creates a new Class
func NewClassService(repo port.Repository, logger port.Logger, config port.Config) *ClassService {
	return &ClassService{
		Base: Base{
			Repo:   repo,
			Logger: logger,
			Config: config,
		},
	}
}

// Create is a method that creates a class
func (s *ClassService) Create(request *dto.ClassCreateRequest) *dto.ClassCreateResponse {
	if err := request.Validate(); err != nil {
		s.LogError(request, err)
		return dto.NewClassCreateResponse(dto.StatusBadRequest, err.Error(), 0)
	}
	tx, err := s.Repo.Begin("assets")
	if err != nil {
		s.LogError(request, err)
		return dto.NewClassCreateResponse(dto.StatusInternalServerError, "internal server error", 0)
	}
	defer s.Repo.Rollback(tx)
	class := entity.NewClass(s.Base.Repo)
	if err := class.Create(request.Name, request.Description, tx); err != nil {
		s.LogError(request, err)
		return dto.NewClassCreateResponse(dto.StatusInternalServerError, "internal server error", 0)
	}
	if err := s.Repo.Commit(tx); err != nil {
		s.LogError(request, err)
		return dto.NewClassCreateResponse(dto.StatusInternalServerError, "internal server error", 0)
	}
	s.LogOk(request)
	return dto.NewClassCreateResponse(dto.StatusSuccess, "class created", class.ID)
}

// GetByID is a method that gets a class by ID
func (s *ClassService) Get(request *dto.ClassGetRequest) *dto.ClassGetResponse {
	if err := request.Validate(); err != nil {
		s.LogError(request, err)
		return dto.NewClassGetResponse(dto.StatusBadRequest, err.Error(), 0, "", "", "")
	}
	tx, err := s.Repo.Begin("assets")
	if err != nil {
		s.LogError(request, err)
		return dto.NewClassGetResponse(dto.StatusInternalServerError, "internal server error", 0, "", "", "")
	}
	defer s.Repo.Rollback(tx)
	c := entity.NewClass(s.Base.Repo)
	if err := c.GetByID(request.ID, nil); err != nil {
		s.LogError(request, err)
		return dto.NewClassGetResponse(dto.StatusInternalServerError, "internal server error", 0, "", "", "")
	}
	s.LogOk(request)
	return dto.NewClassGetResponse(dto.StatusSuccess, "ok", c.ID, c.Name, c.Description, c.CreatedAt.Format(time.DateTime))
}
