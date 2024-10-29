package main

import (
	 "gofr.dev/pkg/gofr"
)

// main is the entry point for the application
func main() {
	app := gofr.New()
	app.GET("/hello", func(ctx *gofr.Context) (interface{}, error) {
		return "Hello, World!", nil
	})
	app.Run()
}