package handler

import (
	"github.com/lavinas/vessel/internal/dto"
	"github.com/lavinas/vessel/internal/port"
)

// Class represents the class create command
type ClassCmd struct {
	Create *ClassCreateCmd `arg:"subcommand:create" help:"Create a class"`
	Get    *ClassGetCmd    `arg:"subcommand:get" help:"Get a class"`
}

// GetDto is a method that gets the correct DTO based on command line arguments
func (c *ClassCmd) GetDto() port.Request {
	switch {
	case c.Create != nil:
		return c.Create.GetDto()
	case c.Get != nil:
		return c.Get.GetDto()
	}
	return nil
}

// ClassCreateCmd represents the class create command
type ClassCreateCmd struct {
	Name        string `arg:"-n,--name" help:"The name of the class"`
	Description string `arg:"-d,--description" help:"The description of the class"`
}

// GetDto is a method that gets the correct DTO based on command line arguments
func (c *ClassCreateCmd) GetDto() port.Request {
	return &dto.ClassCreateRequest{
		Name:        c.Name,
		Description: c.Description,
	}
}

// ClassGetCmd represents the class get command
type ClassGetCmd struct {
	ID int64 `arg:"-i,--id" help:"The id of the class"`
}

// GetDto is a method that gets the correct DTO based on command line arguments
func (c *ClassGetCmd) GetDto() port.Request {
	return &dto.ClassGetRequest{
		ID: c.ID,
	}
}
