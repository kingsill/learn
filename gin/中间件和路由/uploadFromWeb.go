package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func upload(c *gin.Context) {
	fh, _ := c.FormFile("file") //html中设定key为file
	log.Println(fh.Filename)

	trace := "./save/" + fh.Filename //将文件存储到save文件夹下

	c.SaveUploadedFile(fh, trace) //第一个是保存的文件，第二个是保存路径

	c.String(http.StatusOK, fh.Filename)
}

func go_upload(c *gin.Context) {
	c.HTML(200, "upload.html", nil) //通过html中设置重定向到post，以上传文件
}

func main() {
	e := gin.Default()
	e.MaxMultipartMemory = 8 << 20
	e.LoadHTMLGlob("templates/*")

	e.GET("/upload", go_upload)
	e.POST("/upload", upload)
	e.Run()
}
