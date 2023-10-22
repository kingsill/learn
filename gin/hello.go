package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//创建一个默认路由
	router := gin.Default()

	//绑定路由规则和路由函数，访问/index的路由，将由对应的函数（get）去处理
	router.GET("/index", func(context *gin.Context) {
		context.String(200, "hello world")
	})

	//启动监听，gin会把web服务运行在本机的0.0.0.0：8080端口
	router.Run(":8080")
}
