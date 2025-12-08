// 📌 Bagian 2 — Routing Dasar (GET & POST)

// Routing = cara server menentukan “kalau user akses URL ini, jalankan handler ini.”
// ✔ Routing metode HTTP:

// GET

// POST

// PUT

// DELETE

// ✔ Routing dengan parameter:

// Query param

// Path param

// Tapi sekarang kita bahas GET dan POST dulu.

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 	1️⃣ Routing GET

	// GET dipakai untuk ambil data
	r.GET("/halo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"pesan": "ini endpoint GET",
		})
	})

	// 	2️⃣ Routing POST

	// POST dipakai untuk mengirim data.
	r.POST("/kirim", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "data diterima",
		})
	})

	r.Run()
}
