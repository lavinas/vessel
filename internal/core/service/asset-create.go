package service

import (
	"errors"
	"time"

	"github.com/lavinas/vessel/internal/core/entity"
	"github.com/lavinas/vessel/internal/dto"
	"github.com/lavinas/vessel/internal/port"
)

// AssetCreate represents the create asset service
type AssetCreate struct {
	Base
}

// NewAssetCreate creates a new AssetCreate
func NewAssetCreate(repo port.Repository, logger port.Logger, config port.Config) *AssetCreate {
	return &AssetCreate{
		Base: Base{
			Repo:   repo,
			Logger: logger,
			Config: config,
		},
	}
}

// Run is a method that runs the service
func (a *AssetCreate) Run(request *dto.AssetCreateRequest) *dto.AssetCreateResponse {
	if err := request.Validate(); err != nil {
		a.LogError(request, err)
		return dto.NewAssetCreateResponse(dto.StatusBadRequest, err.Error(), 0, "", "", "", "")
	}
	tx, err := a.Repo.Begin("")
	if err != nil {
		a.LogError(request, err)
		return dto.NewAssetCreateResponse(dto.StatusInternalServerError, dto.ErrInternalGeneric, 0, "", "", "", "")
	}
	defer a.Repo.Rollback(tx)
	asset := entity.NewAsset(a.Base.Repo)
	if dup, err := asset.CheckDuplicity(request.Name, tx); err != nil {
		a.LogError(request, err)
		return dto.NewAssetCreateResponse(dto.StatusInternalServerError, dto.ErrInternalGeneric, 0, "", "", "", "")
	} else if dup {
		a.LogError(request, errors.New(dto.ErrAssetCreateRequestDuplicated))
		return dto.NewAssetCreateResponse(dto.StatusBadRequest, dto.ErrAssetCreateRequestDuplicated, 0, "", "", "", "")
	}
	if exist, err := asset.CheckClassExistence(request.ClassName, tx); err != nil {
		a.LogError(request, err)
		return dto.NewAssetCreateResponse(dto.StatusInternalServerError, dto.ErrInternalGeneric, 0, "", "", "", "")
	} else if !exist {
		a.LogError(request, errors.New(dto.ErrAssetCreateRequestClassNotFound))
		return dto.NewAssetCreateResponse(dto.StatusBadRequest, dto.ErrAssetCreateRequestClassNotFound, 0, "", "", "", "")
	}
	if err := asset.Create(request.ClassName, request.Name, request.Description, tx); err != nil {
		a.LogError(request, err)
		return dto.NewAssetCreateResponse(dto.StatusInternalServerError, dto.ErrInternalGeneric, 0, "", "", "", "")
	}
	if err := a.Repo.Commit(tx); err != nil {
		a.LogError(request, err)
		return dto.NewAssetCreateResponse(dto.StatusInternalServerError, dto.ErrInternalGeneric, 0, "", "", "", "")
	}
	a.LogOk(request)
	return dto.NewAssetCreateResponse(dto.StatusSuccess, "ok", asset.ID, request.Name, request.Description, request.ClassName, time.Now().Format(time.DateTime))
}
