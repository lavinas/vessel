package main

import (
	"log"
	"os"

	"github.com/lavinas/vessel/internal/adapter/config"
	"github.com/lavinas/vessel/internal/adapter/handler"
	"github.com/lavinas/vessel/internal/adapter/repository"
	"github.com/lavinas/vessel/internal/core/service"
)

// main is the entry point for the application
func main() {
	repo, err := repository.NewMySql(os.Getenv("MYSQL_DNS"), os.Getenv("MYSQL_SSH"))
	if err != nil {
		panic(err)
	}
	defer repo.Close()
	logger := log.New(os.Stdout, "vessel: ", log.LstdFlags)
	config := config.NewConfig()
	service := service.NewDtoService(repo, logger, config)
	handler := handler.NewCommandLine(service)
	handler.Run()
}