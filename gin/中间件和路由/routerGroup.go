package main

import "github.com/gin-gonic/gin"

type User struct {
	Name string
	Age  int
}

type ArticleList struct {
	Title   string
	Content string
}

func user(c *gin.Context) {
	var Userlist []User = []User{
		{"wang", 18},
		{"zhang", 28},
		{"li", 20},
	}
	c.JSON(200, Userlist)
}

func article(c *gin.Context) {
	var article []ArticleList = []ArticleList{
		{"go", "easy"},
		{"python", "very easy"},
		{"java", "difficult"},
	}
	c.JSON(200, article)
}

func api_Init(e *gin.Engine) { //封包为
	rg := e.Group("/api")
	{
		rg.GET("/userList", user) //此时该路由的地址为 ... /api/userList

		art_Init(rg)
	}

}

func art_Init(rg *gin.RouterGroup) {
	rg2 := rg.Group("/article") //可以分组嵌套
	{
		rg2.GET("/view", article) //该路由地址为 .../api/article/view
	}
}

func main() {

	e := gin.Default()

	api_Init(e)

	e.Run()

}
