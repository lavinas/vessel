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
		return a.setError(request, err, dto.StatusBadRequest, err.Error())
	}
	tx, err := a.Repo.Begin("")
	if err != nil {
		return a.setError(request, err, dto.StatusInternalServerError, dto.ErrInternalGeneric)
	}
	defer a.Repo.Rollback(tx)
	asset, respError := a.create(request, tx)
	if respError != nil {
		return respError
	}
	if err := a.Repo.Commit(tx); err != nil {
		return a.setError(request, err, dto.StatusInternalServerError, dto.ErrInternalGeneric)
	}
	a.LogOk(request)
	return dto.NewAssetCreateResponse(dto.StatusSuccess, "ok", asset.ID, request.Name, request.Description, request.ClassName, time.Now().Format(time.DateTime))
}

// create is a method that creates an asset
func (a *AssetCreate) create(request *dto.AssetCreateRequest, tx interface{}) (*entity.Asset, *dto.AssetCreateResponse) {
	asset := entity.NewAsset(a.Base.Repo)
	if dup, err := asset.CheckDuplicity(request.Name, tx); err != nil {
		return nil, a.setError(request, err, dto.StatusInternalServerError, dto.ErrInternalGeneric)
	} else if dup {
		err := errors.New(dto.ErrAssetCreateRequestDuplicated)
		return nil, a.setError(request, err, dto.StatusBadRequest, err.Error())
	}
	if exist, err := asset.CheckClassExistence(request.ClassName, tx); err != nil {
		return nil, a.setError(request, err, dto.StatusInternalServerError, dto.ErrInternalGeneric)
	} else if !exist {
		err := errors.New(dto.ErrAssetCreateRequestClassNotFound)
		return nil, a.setError(request, err, dto.StatusBadRequest, err.Error())
	}
	if err := asset.Create(request.ClassName, request.Name, request.Description, tx); err != nil {
		return nil, a.setError(request, err, dto.StatusInternalServerError, dto.ErrInternalGeneric)
	}
	return asset, nil
}

// setError is a method that sets the error message
func (a *AssetCreate) setError(request *dto.AssetCreateRequest, err error, status, desc string) *dto.AssetCreateResponse {
	a.LogError(request, err)
	return dto.NewAssetCreateResponse(status, desc, 0, "", "", "", "")
}
