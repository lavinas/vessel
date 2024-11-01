package cline

import (
	"github.com/lavinas/vessel/internal/dto"
	"github.com/lavinas/vessel/internal/port"
	"github.com/lavinas/vessel/internal/core/service"
)

// ClassGetCmd represents the class get command
type ClassGet struct {
	ID int64 `arg:"-i,--id" help:"The id of the class"`
}

// Run is a method that runs the handler
func (c *ClassGet) Run(repo port.Repository, logger port.Logger, config port.Config) port.Response {
	dto := &dto.ClassGetRequest{
		ID: c.ID,
	}
	rn := service.NewClassGet(repo, logger, config)
	return rn.Run(dto)
}