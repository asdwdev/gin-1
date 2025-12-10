package main

import (
	"gin-1/config"
	"gin-1/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Product{})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "database connected",
		})
	})

	r.Run()
}
