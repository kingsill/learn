package main

import "github.com/gin-gonic/gin"

func jwtTokenMiddle(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "1234" { //校验请求头中token是否为1234
		c.Next()
		return //为了当回到响应部分时 停止该函数，避免继续进行
	}
	c.JSON(200, gin.H{"msg": "验证失败"})
	c.Abort() //如果验证失败，则放弃下一步路由的运行
}

func get(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok"})
}

func main() {
	e := gin.Default()

	rg := e.Group("/varify").Use(jwtTokenMiddle) //分组注册路由方式
	{
		rg.GET("/token", get)
	}

	e.Run()
}
