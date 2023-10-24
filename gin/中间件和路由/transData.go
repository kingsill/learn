package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 为了展示泛型定义一结构体
type Person struct {
	Name string
	Age  int
}

func GM1(c *gin.Context) {
	c.Set("user", Person{"WANG", 18}) //set的后面一个传入参数为any。即泛型
}

func GM2(c *gin.Context) {
	fmt.Println("hello")
}

func main() {
	e := gin.Default()

	e.Use(GM1, GM2)

	e.GET("/", func(c *gin.Context) {
		value, _ := c.Get("user")

		c.JSON(200, gin.H{"msg": value}) //由于泛型的使用，整体可以不再断言

		//但是当单独使用其中数值时，需要进行断言
		p := value.(Person)
		fmt.Println(p.Name)
	})

	e.Run()
}
