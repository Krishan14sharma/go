package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	//k := User{a:"vardaan", b:21}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, time.Now().Unix())
	})
	r.Run(":2000") // listen and server on 0.0.0.0:8080
}

func add(a, b int) (int, i`nt) {
	return a, b
}

type User struct {
	a string `json:"name"`
	b int    `json:"age"`
}

