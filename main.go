// menerima body json (post request)
package main

import "github.com/gin-gonic/gin"

type UserInput struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var input UserInput

		// baca JSON dari body ke variabel input
		if err := c.BindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// kalau sukses
		c.JSON(200, gin.H{
			"message": "data diterima!",
			"data":    input,
		})
	})

	r.Run()
}
