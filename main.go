package main

import (
	"gin-1/config"
	"gin-1/controllers"
	"gin-1/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Product{})

	r.POST("/product", controllers.CreateProduct)

	r.Run()
}
