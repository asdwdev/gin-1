// latihan bagian 7
package main

import "github.com/gin-gonic/gin"

func CheckApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-API-KEY")
		if token != "rahasia123" {
			c.JSON(401, gin.H{"error": "API key salah"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	admin := r.Group("/admin")
	admin.Use(CheckApiKey())

	admin.GET("/info", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "akses diterima"})
	})

	r.Run()
}
