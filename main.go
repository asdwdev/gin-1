// membuat route get pertama
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// ini route GET pertama kita
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello dari gin!")
	})

	// versi json
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "halo dunia JSON",
			"success": true,
		})
	})

	r.Run() // port 8080
}
