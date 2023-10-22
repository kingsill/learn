package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 查询参数
func query(c *gin.Context) {
	//getquery方法获得传入的数据及是否有对应的关键词
	fmt.Println(c.GetQuery("user"))
	//query方法内调用getquery方法，获得传入的数据
	fmt.Println(c.Query("user"))

	fmt.Println(c.QueryArray("user")) //拿到多个相同的查询参数,返回切片类型

	fmt.Println(c.DefaultQuery("addr", "上海")) //设定默认的输出，没有输入的话则使用这个值
}

// 动态参数
func parameter(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("book_id"))
}

// 表单参数,使用post请求
func form(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Println(name)
	s := c.PostFormArray("name") //获取多个相同的表单参数，以切边输出
	fmt.Printf("s: %v\n", s)
	c.JSON(200, gin.H{"name_array": s, "name": name})

	c.DefaultPostForm("city", "淄博") //如果用户没传，就使用默认值

	fmt.Println(c.MultipartForm()) //接受所有的form表单，包括文件
}

// 原始参数
func raw(c *gin.Context) {
	// b, _ := c.GetRawData()
	// fmt.Printf("%v\n", string(b))
	s := c.GetHeader("Content-Type") //获得header中的内容
	fmt.Printf("s: %v\n", s)
	switch s {
	case "application/json":
		type user struct { //通过增加备注可以使json中格式对应字段名，都一样也可以不备注
			Name string `json:"name"`
			Addr string `json:"city"`
			Age  int
		}
		var User user //结构体实例化
		c.ShouldBindJSON(&User)

		// 当然也可以通过json.Unmarshal方法来进行json到结构体的转化

		fmt.Printf("user: %v\n", User)
	}
}

func main() {
	e := gin.Default()

	e.GET("/quest", query)
	e.GET("/param/:user_id/:book_id", parameter) //注意这里每个字符前都加：

	e.POST("/list", form)
	e.POST("/raw", raw)

	e.Run(":8020")
}
