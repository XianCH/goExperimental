package main

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/x14n/goExperimental/gin-crud-api/v1/config"
	"github.com/x14n/goExperimental/gin-crud-api/v1/controller"
	"github.com/x14n/goExperimental/gin-crud-api/v1/helper"
	"github.com/x14n/goExperimental/gin-crud-api/v1/repository"
	"github.com/x14n/goExperimental/gin-crud-api/v1/router"
	"github.com/x14n/goExperimental/gin-crud-api/v1/service"
)

func main() {

	//Database
	db, _ := config.DatabaseConnection()
	validate := validator.New()

	//Init Repository
	tagRepository := repository.NewTagsRepositoryImpl(db)

	//Init Service
	tagService := service.NewTagServiceImpl(tagRepository, validate)

	//Init controller
	tagController := controller.NewTagController(tagService)

	//Router
	routes := router.NewRouter(tagController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
