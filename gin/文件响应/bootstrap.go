package main

// 集成bootstrap前端框架
import (
	"github.com/gin-gonic/gin"
)

func try(c *gin.Context) {
	c.HTML(200, "test.html", nil)
}

func main() {
	e := gin.Default()

	e.LoadHTMLGlob("templates/*") //前面是url路径，后面是实际路径

	e.GET("/try", try)

	e.Run()
}
