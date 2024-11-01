package handler

import (
	"github.com/lavinas/jib/internal/entity"
	"gofr.dev/pkg/gofr"
)

// Rest is a struct that represents the REST handler
type Rest struct {
}

// NewRest is a function that creates a new Rest instance
func NewRest() *Rest {
	return &Rest{}
}

// Run is a method that runs the server
func (g *Rest) Run() {
	app := gofr.New()
	app.GET("/ping", func(ctx *gofr.Context) (interface{}, error) {
		return "pong", nil
	})
	if err := app.AddRESTHandlers(&entity.User{}); err != nil {
	    return
	}
	app.Run()
}
