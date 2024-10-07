package handler

import (
	"fmt"

	"github.com/alexflint/go-arg"
	"github.com/lavinas/vessel/internal/dto"
	"github.com/lavinas/vessel/internal/port"
)

// args represents the arguments for the class command
type args struct {
	Action         string          `arg:"-a,--action" help:"Action to perform"`
	Object         string          `arg:"-o,--object" help:"The object of the class"`
	ClassCreateCmd *ClassCreateCmd `arg:"subcommand:create" help:"Create a class"`
	ClassGetCmd    *ClassGetCmd    `arg:"subcommand:get" help:"Get a class"`
}

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
	args := args{}
	arg.MustParse(&args)
	dto := c.getDto(args)
	if dto == nil {
		fmt.Println("Invalid command")
		return
	}
	response := c.service.Run(dto)
	fmt.Println(response.ToLine())
}

// getDto is a method that gets the correct DTO based on command line arguments
func (c *CommandLine) getDto(arg args) port.Request {
	switch {
	case arg.Action == "create" && arg.Object == "Class":
		return &dto.ClassCreateRequest{
			Name:        arg.ClassCreateCmd.Name,
			Description: arg.ClassCreateCmd.Description,
		}
	case arg.Action == "get" && arg.Object == "Class":
		return &dto.ClassGetRequest{
			ID: arg.ClassGetCmd.ID,
		}
	}
	return nil
}
