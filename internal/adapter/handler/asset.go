package handler

import (
	"github.com/lavinas/vessel/internal/dto"
	"github.com/lavinas/vessel/internal/port"
)

// AssetCmd represents the asset create command
type AssetCmd struct {
	Create *AssetCreateCmd `arg:"subcommand:create" help:"Create an asset"`
	Get    *AssetGetCmd    `arg:"subcommand:get" help:"Get an asset"`
}

// GetDto is a method that gets the correct DTO based on command line arguments
func (c *AssetCmd) GetDto() port.Request {
	switch {
	case c.Create != nil:
		return c.Create.GetDto()
	case c.Get != nil:
		return c.Get.GetDto()
	}
	return nil
}

// AssetCreateCmd represents the asset create command
type AssetCreateCmd struct {
	Name        string `arg:"-n,--name" help:"The name of the asset"`
	Description string `arg:"-d,--description" help:"The description of the asset"`
	ClassName   string `arg:"-c,--class" help:"The class name of the asset"`
}

// GetDto is a method that gets the correct DTO based on command line arguments
func (c *AssetCreateCmd) GetDto() port.Request {
	return &dto.AssetCreateRequest{
		Name:        c.Name,
		Description: c.Description,
		ClassName:   c.ClassName,
	}
}

// AssetGetCmd represents the asset get command
type AssetGetCmd struct {
	ID int64 `arg:"-i,--id" help:"The id of the asset"`
}

// GetDto is a method that gets the correct DTO based on command line arguments
func (c *AssetGetCmd) GetDto() port.Request {
	return &dto.AssetGetRequest{
		ID: c.ID,
	}
}
