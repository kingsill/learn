package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// test
func GM(c *gin.Context) {
	fmt.Println("GLOBAL")
}

func main() {
	e := gin.Default()

	e.Use(GM)

	e.GET("/")

	e.Run()
}
