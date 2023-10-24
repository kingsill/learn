package main

import "github.com/gin-gonic/gin"

//这里可以回顾之前的静态文件响应
func main() {
	e := gin.Default()

	e.GET("/download", download) //为了方便，使用get

	e.Run()
}

func download(c *gin.Context) {

	//响应头设置
	c.Header("Content-Type", "application/octet-stream")              //用来表示是文件流，唤起浏览器下载，一般这个设置一定要跟下面的文件名设置
	c.Header("Content-Disposition", "attachment; filename="+"牛逼.png") // 用来指定下载下来的文件名
	// c.Header("Content-Transfer-Encoding", "binary")                   // 表示传输过程中的编码形式，乱码问题可能就是因为它

	c.File("save/Default.jpg") //这里因为是图片文件格式，所以浏览器自动预览
}
