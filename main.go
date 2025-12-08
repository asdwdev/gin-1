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

type RegisterRequest struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
}

func main() {
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		var req RegisterRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "body tidak valid valid"})
			return
		}

		c.JSON(200, gin.H{
			"pesan": "register berhasil",
			"data": gin.H{
				"nama":  req.Nama,
				"email": req.Email,
			},
		})
	})

	r.Run()
}
