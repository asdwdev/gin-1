// cara bikin project gin & menjalankan server kosong
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default() // bikin router
	r.Run()            // nyalain server (default port 8080)
}
