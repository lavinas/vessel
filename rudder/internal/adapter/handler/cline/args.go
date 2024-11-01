package cline

import (
	"github.com/lavinas/vessel/internal/port"
)

// args represents the arguments for the class command
type Args struct {
	Class *Class `arg:"subcommand:class" help:"Class commands"`
	Asset *Asset `arg:"subcommand:asset" help:"Asset commands"`
}

// Run is a method that runs the handler
func (a *Args) Run(repo port.Repository, logger port.Logger, config port.Config) port.Response {
	switch {
	case a.Class != nil:
		return a.Class.Run(repo, logger, config)
	case a.Asset != nil:
		return a.Asset.Run(repo, logger, config)
	}
	return nil
}
