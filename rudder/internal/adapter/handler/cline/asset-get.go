package cline

import (
	"github.com/lavinas/vessel/internal/port"
	"github.com/lavinas/vessel/internal/dto"
	"github.com/lavinas/vessel/internal/core/service"
)

// AssetGetCmd represents the asset get command
type AssetGet struct {
	ID int64 `arg:"-i,--id" help:"The id of the asset"`
}

// Run is a method that runs the handler
func (c *AssetGet) Run(repo port.Repository, logger port.Logger, config port.Config) port.Response {
	dto := &dto.AssetGetRequest{
		ID: c.ID,
	}
	rn := service.NewAssetGet(repo, logger, config)
	return rn.Run(dto)
}