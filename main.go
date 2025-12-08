// latihan bagian 2
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "login berhasil",
		})
	})

	r.Run()
}
