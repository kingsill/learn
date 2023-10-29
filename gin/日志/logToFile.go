package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func nihao(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "看文件"})
}

func main() {
	//创建相关文件并留下接口
	f, _ := os.Create("save/log.log")

	//这部分要写在default之前
	// gin.DefaultWriter = f //f os.file 本身也实现了writer方法，如果只需要输出到文件里，可以写这一句，如果要同时输出到控制台，可以用下一句
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) //同时在f和控制台进行输出

	e := gin.Default()

	e.GET("/try", nihao)

	e.Run()
}
