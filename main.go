// Bagian 4 — Path Parameter

// Contoh:
// /user/10
// /produk/12345
// /blog/2024/01/10

// Path param dipakai saat kamu ingin menangani:

// /user/:id

// /article/:slug

// /order/:orderId
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 1️⃣ Cara membuat path parameter
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"user_id": id,
		})
	})

	// 2️⃣ Bisa lebih dari satu parameter
	r.GET("/blog/:tahun/:bulan/:tanggal", func(c *gin.Context) {
		tahun := c.Param("tahun")
		bulan := c.Param("bulan")
		tgl := c.Param("tanggal")

		c.JSON(200, gin.H{
			"tahun":  tahun,
			"bulan":  bulan,
			"tangal": tgl,
		})
	})

	r.Run()
}
