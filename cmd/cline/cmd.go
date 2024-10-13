package main

import (
	"log"
	"os"

	"github.com/lavinas/vessel/internal/adapter/config"
	"github.com/lavinas/vessel/internal/adapter/handler/cline"
	"github.com/lavinas/vessel/internal/adapter/repository"
)

// main is the entry point for the application
func main() {
	repo := getRepository()
	defer repo.Close()
	logger := getLogger()
	config := config.NewConfig()
	handler := cline.NewRunner(repo, logger, config)
	handler.Run()
}

// getRepository is a helper function that returns a new repository
func getRepository() *repository.MySql {
	repo, err := repository.NewMySql(os.Getenv("MYSQL_DNS"), os.Getenv("MYSQL_SSH"))
	if err != nil {
		panic(err)
	}
	return repo
}

// getLogger is a helper function that returns a new logger
func getLogger() *log.Logger {
	f, err := os.OpenFile("vessel.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return log.New(f, "vessel: ", log.LstdFlags)
}