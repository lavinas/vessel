package cline

import (
	"github.com/lavinas/vessel/internal/port"
)

// Class represents the class create command
type Class struct {
	Create *ClassCreate `arg:"subcommand:create" help:"Create a class"`
	Get    *ClassGet    `arg:"subcommand:get" help:"Get a class"`
}


// Run is a method that runs the handler
func (c *Class) Run(repo port.Repository, logger port.Logger, config port.Config) port.Response {
	switch {
	case c.Create != nil:
		return c.Create.Run(repo, logger, config)
	case c.Get != nil:
		return c.Get.Run(repo, logger, config)
	}
	return nil
}
