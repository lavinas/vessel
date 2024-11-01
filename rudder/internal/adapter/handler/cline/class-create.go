package cline

import (
	"github.com/lavinas/vessel/internal/port"
	"github.com/lavinas/vessel/internal/dto"
	"github.com/lavinas/vessel/internal/core/service"
)

// ClassCreateCmd represents the class create command
type ClassCreate struct {
	Name        string `arg:"-n,--name" help:"The name of the class"`
	Description string `arg:"-d,--description" help:"The description of the class"`
}

// Run is a method that runs the handler
func (c *ClassCreate) Run(repo port.Repository, logger port.Logger, config port.Config) port.Response {
	dto := &dto.ClassCreateRequest{
		Name:        c.Name,
		Description: c.Description,
	}
	rn := service.NewClassCreate(repo, logger, config)
	return rn.Run(dto)
}