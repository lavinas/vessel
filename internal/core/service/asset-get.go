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
		a.LogError(request, err)
		return dto.NewAssetGetResponse(dto.StatusBadRequest, err.Error(), 0, "", "", "", "")
	}
	tx, err := a.Repo.Begin("")
	if err != nil {
		a.LogError(request, err)
		return dto.NewAssetGetResponse(dto.StatusInternalServerError, dto.ErrInternalGeneric, 0, "", "", "", "")
	}
	defer a.Repo.Rollback(tx)
	asset := entity.NewAsset(a.Base.Repo)
	if err := asset.GetByID(request.ID, tx); err != nil {
		a.LogError(request, err)
		return dto.NewAssetGetResponse(dto.StatusInternalServerError, dto.ErrInternalGeneric, 0, "", "", "", "")
	}
	if !asset.IsLoaded() {
		a.LogError(request, errors.New(dto.ErrAssetGetRequestNotFound))
		return dto.NewAssetGetResponse(dto.StatusNotFound, dto.ErrAssetGetRequestNotFound, 0, "", "", "", "")
	}
	a.LogOk(request)
	return dto.NewAssetGetResponse(dto.StatusSuccess, "", asset.ID, asset.Name, asset.Description, asset.Class.Name, asset.CreatedAt.Format(time.DateTime))
}
