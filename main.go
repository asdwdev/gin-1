// Bagian 5 — POST JSON (Body Request)
// API modern sangat sering menerima JSON, misalnya:

// login (username, password)

// register

// tambah produk

// update profil

// upload data form

// Di Gin, cara ambil JSON:

// var body struct {
//     Username string `json:"username"`
//     Password string `json:"password"`
// }

// if err := c.BindJSON(&body); err != nil {
//     c.JSON(400, gin.H{"error": "json invalid"})
//     return
// }

//	c.JSON(200, gin.H{
//	    "username": body.Username,
//	    "password": body.Password,
//	})
package main

import "github.com/gin-gonic/gin"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		var req LoginRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "body tidak valid"})
			return
		}

		c.JSON(200, gin.H{
			"username": req.Username,
			"status":   "login oke",
		})
	})

	r.Run()
}
