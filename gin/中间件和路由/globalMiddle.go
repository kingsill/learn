package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//全局中间件的运行也可以使用next来进行请求和响应的划分

func GM(c *gin.Context) {
	fmt.Println("GLOBAL")
	c.Set("name", "wang")

	c.Next()
	fmt.Println("No")

}

func GM1(c *gin.Context) {
	fmt.Println("GLOBAL！")
	c.Next()
	fmt.Println("No!")

}

func main() {
	e := gin.Default()

	e.Use(GM, GM1) //使用use来进行全局中间件的注册。use后面可以跟多个中间件
	//全局中间件，不管运行哪个路由都会运行全局的中间件

	e.GET("/", func(c *gin.Context) {
		value, _ := c.Get("name")
		fmt.Printf("name: %v\n", value)
		c.JSON(200, gin.H{"msg": "main"})
	})

	e.Run()
}
