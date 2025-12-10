// latihan 6
package main

import "github.com/gin-gonic/gin"

type ProductRequest struct {
	Nama  string `json:"nama" binding:"required"`
	Stok  int    `json:"stok" binding:"required,min=1"`
	Harga int    `json:"harga" binding:"required,min=1000"`
}

func main() {
	r := gin.Default()

	r.POST("/produk/tambah", func(c *gin.Context) {
		var req ProductRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"pesan": "produk ditambahkan",
			"data":  req,
		})
	})

	r.Run()
}
