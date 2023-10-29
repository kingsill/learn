package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	//这里设置调试日志输出格式
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("wang %s %s %s %d \n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	e := gin.Default()

	e.GET("/try")

	//.routes查看路由信息
	for _, info := range e.Routes() {
		fmt.Printf("info.Method: %v\n", info.Method)
		fmt.Printf("info.Handler: %v\n", info.Handler)
	}

	e.Run()
}
