package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/greet", func(c *gin.Context) {
		nama := c.Query("nama")
		umur := c.DefaultQuery("umur", "0")

		c.JSON(200, gin.H{
			"nama": nama,
			"umur": umur,
		})
	})

	r.Run()
}
