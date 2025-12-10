package controllers

import (
	"gin-1/config"
	"gin-1/models"

	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Nama  string `json:"nama" binding:"required"`
	Stok  int    `json:"stok" binding:"required,min=1"`
	Harga int    `json:"harga" binding:"required,min=1000"`
}

func CreateProduct(c *gin.Context) {
	var req CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		Nama:  req.Nama,
		Stok:  req.Stok,
		Harga: req.Harga,
	}

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(500, gin.H{"error": "gagal menyimpan produk"})
		return
	}

	c.JSON(200, gin.H{
		"message": "produk berhasil ditambahkan",
		"data":    product,
	})
}
