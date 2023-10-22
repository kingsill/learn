package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 以string进行响应，返回string
func res_string(c *gin.Context) {
	c.String(200, "ok")
}

// 以json进行响应，返回json
func res_json(c *gin.Context) {
	//json响应结构体，同理也可以响应map，注意map是无序的，其内容先后顺序随机
	type userjson struct {
		Name     string `json:"姓名"` //用备注来进行显示
		Id       int
		Age      int
		Password string `json:"-"` //"-"代表这段不进行渲染，也可以使用password，不再包外进行使用来防止显示
	}
	jsonUser := userjson{"wang", 1, 12, "wqer"}
	//响应后面的C.json去了。
	c.JSON(200, jsonUser)

	//从map中延申出去，这里即响应map，可以通过多次嵌套来方便json的书写
	c.JSON(200, gin.H{"username": "fengfeng", "age": 18})
	// gin.H是map[string]any即map[string]interface{}的缩写。
}

// 响应xml
func res_xml(c *gin.Context) {
	c.XML(200, gin.H{"name": "li", "major": gin.H{"grade": "fresh", "title": "machine"}})
}

// 响应yaml
func res_yaml(c *gin.Context) {
	c.YAML(200, gin.H{"name": "li", "major": gin.H{"grade": "fresh", "title": "machine"}})
}

// 响应html
func res_html(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"username": "wang2"}) //注意这里的字段名需要与html模板文件中{{.XXX}}一致，否则无法传入
}

// 重定向
func redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")

}

func main() {
	e := gin.Default()

	e.LoadHTMLGlob("templates/*") //html响应时，需要使用loadhtmlglob或loadhtmlfiles方法来加载模板文件

	//文件响应

	//配置单独的一个静态文件
	// e.StaticFile("/static/Default.jpg", "/golang/gin/static/Default.jpg")//单独“/”代表从根目录进行路径识别，这里是从盘符开始
	e.StaticFile("/static/Default.jpg", "./static/Default.jpg") //“./”即从当前目录开始寻径
	// e.StaticFile("/static/Default.jpg", "static/Default.jpg")//什么都不写等同于“./”
	//golang中没有相对文件的路径，只有相对项目的路径，即无论你在什么位置写代码，都是以项目为根目录

	//网页请求某个静态目录
	e.StaticFS("/newstatic", http.Dir("./static/static"))

	e.GET("/text", res_string)
	e.GET("/json", res_json)
	e.GET("/xml", res_xml)
	e.GET("/yaml", res_yaml)
	e.GET("/html", res_html)
	e.GET("/baidu", redirect)

	e.Run(":8000")
}
