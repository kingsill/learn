package main

import "github.com/gin-gonic/gin"

func val(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok"})
}

func bef(c *gin.Context) {
	c.JSON(200, gin.H{"news": "niubi"})
}

func main() {
	e := gin.Default()

	rg := e.Group("/group").Use(bef)
	{
		rg.GET("/ok", val)
	}

	e.Run()
}
