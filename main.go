package main

import (
	"memetpaten/go-rest/app"
	"memetpaten/go-rest/controller"
	"memetpaten/go-rest/repository"
	"memetpaten/go-rest/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	// declare db and validate
	db := app.NewDB()
	validate := validator.New()

	// declare category repository
	categoryRepository := repository.NewCategoryRepository()
	// declare category service
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	// declare category controller
	categoryController := controller.NewCategoryController(categoryService) 

	// declare router
	router := httprouter.New()

	// implement category routes
	router.POST("/api/v1/categories", categoryController.Create)
	router.GET("/api/v1/categories", categoryController.FindAll)
	router.GET("/api/v1/categories/:id", categoryController.FindById)
	router.PUT("/api/v1/categories/:id", categoryController.Update)
	router.DELETE("/api/v1/categories/:id", categoryController.Delete)
}