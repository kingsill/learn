package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	e.POST("/file", upload)

	e.Run()
}

func upload(c *gin.Context) {
	f, _ := c.MultipartForm()
	fh := f.File["f1"] //这里严格对应form-data中key值

	fmt.Printf("len(fh): %v\n", len(fh)) //获得文件数目

	for _, v := range fh {
		fmt.Printf("v.Filename: %v\n", v.Filename) //输出文件名

		trace := "./save/" + v.Filename //文件的保存路径，./代表从当前项目的根目录进行路径的查询

		c.SaveUploadedFile(v, trace)
	}
	c.JSON(200, gin.H{"msg": "ok"})
}
