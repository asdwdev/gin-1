// 📌 Bagian 3 — Query Parameter
// (Contoh: /search?q=golang)

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 1️⃣ Apa itu Query Param?

	// Format:

	// /endpoint?key=value

	// Contoh:

	// /hello?nama=nil

	// Cara ambilnya:

	// nama := c.Query("nama")

	// Jika tidak ada query param:

	// c.DefaultQuery("nama", "guest")

	r.GET("/hello", func(c *gin.Context) {
		// nama := c.Query("nama")
		nama := c.DefaultQuery("nama", "nil")
		c.JSON(200, gin.H{
			"halo": nama,
		})
	})

	r.Run()

}
