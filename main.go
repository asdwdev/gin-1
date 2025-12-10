// ⭐ Bagian 7 — Middleware (Dasar & Latihan)
// 1️⃣ Apa itu Middleware?

// Middleware = fungsi yang dijalankan Sebelum dan/atau Setelah handler endpoint.

// request → middleware → handler → middleware → response

// func MyMiddleware(c *gin.Context) {
//     // sebelum handler
//     ...
//     c.Next() // lanjut ke handler

//	    // setelah handler
//	    ...
//	}
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("ada request masuk:", c.Request.URL.Path)
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(LogRequest())

	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run()
}
