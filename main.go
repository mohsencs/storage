package main

import (
	"fmt"
	"log"
	"storage/controller"
	"storage/model"
	"storage/repo"
	"storage/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := gin.Default()
	basePath := "/api"
	api := r.Group(basePath)
	pGroup := api.Group("promotion")

	mongoStorage, err := model.NewMongoStorage()
	if err != nil {
		log.Fatal(err)
	}
	pRepository := repo.NewPromotionRepository(mongoStorage)
	pService := service.NewPromotionService(pRepository)
	controller.NewPromotionController(pGroup, pService)

	r.Run()

	fmt.Println("finished.")
}
