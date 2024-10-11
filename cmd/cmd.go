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
	f, err := os.OpenFile("vessel.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger := log.New(f, "vessel: ", log.LstdFlags)
	config := config.NewConfig()
	service := service.NewDtoGeneric(repo, logger, config)
	handler := handler.NewCommandLine(service)
	handler.Run()
}
