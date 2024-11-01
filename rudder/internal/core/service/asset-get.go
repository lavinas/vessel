package service

import (
	"errors"
	"time"

	"github.com/lavinas/vessel/internal/core/entity"
	"github.com/lavinas/vessel/internal/dto"
	"github.com/lavinas/vessel/internal/port"
)

// AssetGet represents the asset get service
type AssetGet struct {
	Base
}

// NewAssetGet creates a new AssetGet
func NewAssetGet(repo port.Repository, logger port.Logger, config port.Config) *AssetGet {
	return &AssetGet{
		Base: Base{
			Repo:   repo,
			Logger: logger,
			Config: config,
		},
	}
}

// Run is a method that runs the service
func (a *AssetGet) Run(request *dto.AssetGetRequest) *dto.AssetGetResponse {
	if err := request.Validate(); err != nil {
		return a.setError(request, err, dto.StatusBadRequest, err.Error())
	}
	tx, err := a.Repo.Begin("")
	if err != nil {
		return a.setError(request, err, dto.StatusInternalServerError, dto.ErrInternalGeneric)
	}
	defer a.Repo.Rollback(tx)
	asset := entity.NewAsset(a.Base.Repo)
	if err := asset.GetByID(request.ID, tx); err != nil {
		return a.setError(request, err, dto.StatusInternalServerError, dto.ErrInternalGeneric)
	}
	if !asset.IsLoaded() {
		err := errors.New(dto.ErrAssetGetRequestNotFound)
		return a.setError(request, err, dto.StatusNotFound, err.Error())
	}
	a.LogOk(request)
	return dto.NewAssetGetResponse(dto.StatusSuccess, "", asset.ID, asset.Name, asset.Description, asset.Class.Name, asset.CreatedAt.Format(time.DateTime))
}

// setError is a method that sets an error
func (a *AssetGet) setError(request *dto.AssetGetRequest, err error, status, message string) *dto.AssetGetResponse {
	a.LogError(request, err)
	return dto.NewAssetGetResponse(status, message, 0, "", "", "", "")
}
