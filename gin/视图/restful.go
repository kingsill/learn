package main

//由于并非实际情况，函数中本应由数据库传输的内容由函数内给定
//简单演示
import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 定义文章结构体
type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 封装响应内容
type response struct {
	Code int `json:"code"`
	Data any
	Msg  string
}

// 文章列表页面
func getList(c *gin.Context) {
	//自定义文章列表
	articleList := []ArticleModel{
		{"golang", "go语言入门内容"},
		{"python", "python机器学习内容"},
		{"java", "java后端内容"},
	}

	c.JSON(200, response{200, articleList, "文章列表"})
}

// 获取文章详情
func getDetail(c *gin.Context) {
	//获取param中的id
	fmt.Printf("c.Param(\"id\"): %v\n", c.Param("id"))

	article := ArticleModel{"golang", "go语言入门内容"}
	c.JSON(200, response{200, article, "文章详情"})
}

// 创建文章
func create(c *gin.Context) {
	// 接受前端传递过来的json参数

	var article ArticleModel
	c.ShouldBindJSON(&article)

	c.JSON(200, response{200, article, "创建文章"})
}

// 编辑文章
func update(c *gin.Context) {
	//获取要修改的文章的id
	c.Param("id")

	//获取要修改的文章的内容
	var article ArticleModel
	c.ShouldBindJSON(&article)

	c.JSON(200, response{200, article, "修改成功"})
}

// 删除文章
func delete(c *gin.Context) {
	//获取要删除文章的id
	c.Param("id")

	//剩下的操作交给数据库
	c.JSON(200, gin.H{})
}
func main() {
	e := gin.Default()

	e.GET("/articles", getList)       //文章列表
	e.GET("/articles/:id", getDetail) //文章详情
	e.POST("/articles", create)       //创建文章
	e.PUT("articles/:id", update)     //更新文章
	e.DELETE("/articles/:id", delete) //删除文章

	e.Run(":8040")
}
