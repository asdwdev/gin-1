package main

import "github.com/gin-gonic/gin"

func main() {
	// bikin router gin
	r := gin.Default()

	// route GET di path "/"
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin!",
		})
	})

	// jalankan server di port 8080
	r.Run() // default :8008
}
