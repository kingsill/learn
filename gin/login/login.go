package main

//实现登录
import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

// 可以参考之前写的响应html文件
func Dologin(c *gin.Context) {
	username := c.PostForm("username")                       //获取表单参数
	c.HTML(200, "welcome.html", gin.H{"username": username}) //将username传入

}

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")

	e.GET("/login", Login)
	e.POST("/login", Dologin)

	e.Run()
}
