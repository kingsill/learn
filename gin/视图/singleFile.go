package main

import "github.com/gin-gonic/gin"

func postFile(c *gin.Context) {

}

func main() {
	e := gin.Default()
	e.POST("/file", postFile)

	e.Run()
}
