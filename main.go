// latihan bagian 4
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/user/:nama", func(c *gin.Context) {
		nama := c.Param("nama")
		c.JSON(200, gin.H{
			"nama": nama,
		})
	})

	r.GET("/produk/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id": id,
		})
	})

	r.Run()
}
