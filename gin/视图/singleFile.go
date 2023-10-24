package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func postFile(c *gin.Context) {

	fh, _ := c.FormFile("file")

	fmt.Printf("fh.Filename: %v\n", fh.Filename)
	fmt.Printf("fh.Size: %v\n", fh.Size)

	trace := "./save/" + fh.Filename //文件的保存路径，./代表从当前项目的根目录进行路径的查询
	// c.SaveUploadedFile(fh, trace)

	f, _ := fh.Open()
	//读取文件内的信息
	b, _ := io.ReadAll(f)
	fmt.Printf("b: %v\n", b)

	f2, _ := os.Create(trace)
	io.Copy(f2, f)

	defer f.Close()
	defer f2.Close()

	c.JSON(200, gin.H{"msg": "upload success"})
}

func main() {
	e := gin.Default()

	//设置内存限制
	//默认单位是字节， <<代表左移位， 8<<20=8*2^20，2^10=1024，这里设置的内存限制即为8MB
	e.MaxMultipartMemory = 8 << 20 //
	e.POST("/file", postFile)

	e.Run()
}
