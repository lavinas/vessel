package cline

import (
	"github.com/lavinas/vessel/internal/port"
)

// AssetCmd represents the asset create command
type Asset struct {
	Create *AssetCreate `arg:"subcommand:create" help:"Create an asset"`
	Get    *AssetGet    `arg:"subcommand:get" help:"Get an asset"`
}

// Run is a method that runs the handler
func (c *Asset) Run(repo port.Repository, logger port.Logger, config port.Config) port.Response {
	switch {
	case c.Create != nil:
		return c.Create.Run(repo, logger, config)
	case c.Get != nil:
		return c.Get.Run(repo, logger, config)
	}
	return nil
}