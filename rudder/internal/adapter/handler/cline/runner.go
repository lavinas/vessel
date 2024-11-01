package cline

import (
	"fmt"

	"github.com/alexflint/go-arg"
	"github.com/lavinas/vessel/internal/port"
)

// CommandLine represents the command line
type Runner struct {
	repo  port.Repository
	logger port.Logger
	config port.Config
}

// NewCommandLine creates a new CommandLine
func NewRunner(repo port.Repository, logger port.Logger, config port.Config) *Runner {
	return &Runner{
		repo:  repo,
		logger: logger,
		config: config,
	}
}

// Run is a method that runs the command line
func (r *Runner) Run() {
	args := Args{}
	arg.MustParse(&args)
	response := args.Run(r.repo, r.logger, r.config)
	fmt.Println(response.String())
}
