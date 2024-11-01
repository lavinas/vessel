package main

import (
	"github.com/lavinas/jib/internal/adapter/handler"
)

// main is the entry point of the application
func main() {
	handler.NewRest().Run()
}
