package handler

import (
	"fmt"

	"github.com/alexflint/go-arg"
	"github.com/lavinas/vessel/internal/port"
)

// CommandLine represents the command line
type CommandLine struct {
	service port.Service
}

// NewCommandLine creates a new CommandLine
func NewCommandLine(service port.Service) *CommandLine {
	return &CommandLine{
		service: service,
	}
}

// Run is a method that runs the command line
func (c *CommandLine) Run() {
	args := Args{}
	arg.MustParse(&args)
	dto := args.GetDto()
	if dto == nil {
		fmt.Println("Invalid Command")
		return
	}
	response := c.service.Run(dto)
	fmt.Println(response.ToLine())
}

// args represents the arguments for the class command
type Args struct {
	Class *ClassCmd `arg:"subcommand:class" help:"Class commands"`
}

// GetDto is a method that gets the correct DTO based on command line arguments
func (a *Args) GetDto() port.Request {
	switch {
	case a.Class != nil:
		return a.Class.GetDto()
	}
	return nil
}
