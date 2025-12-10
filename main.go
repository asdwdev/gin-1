// Bagian 6 — Validator (VERY IMPORTANT)
package main

import "github.com/gin-gonic/gin"

type RegisterRequest struct {
	Nama  string `json:"nama" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func main() {
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		var req RegisterRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"pesan": "register berhasil",
			"data":  req,
		})
	})

	r.Run()
}
