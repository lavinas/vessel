package cline

import (
	"github.com/lavinas/vessel/internal/dto"
	"github.com/lavinas/vessel/internal/port"
	"github.com/lavinas/vessel/internal/core/service"
)

// AssetCreateCmd represents the asset create command
type AssetCreate struct {
	Name        string `arg:"-n,--name" help:"The name of the asset"`
	Description string `arg:"-d,--description" help:"The description of the asset"`
	ClassName   string `arg:"-c,--class" help:"The class name of the asset"`
}

// Run is a method that runs the handler
func (c *AssetCreate) Run(repo port.Repository, logger port.Logger, config port.Config) port.Response {
	dto := &dto.AssetCreateRequest{
		Name:        c.Name,
		Description: c.Description,
		ClassName:   c.ClassName,
	}
	rn := service.NewAssetCreate(repo, logger, config)
	return rn.Run(dto)
}

