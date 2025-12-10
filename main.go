// 3️⃣ Contoh Middleware Auth Paling Dasar
package main

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "secret123" {
			c.JSON(401, gin.H{"error": "unauhorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()

	private := r.Group("/private")
	private.Use(AuthMiddleware())

	private.GET("/dashboard", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "selamat datang"})
	})

	r.Run()
}
