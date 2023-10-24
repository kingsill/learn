package main

//绑定参数
import "github.com/gin-gonic/gin"

//定义结构体信息
type Userinfo struct { //json用于json绑定。form标记用于绑定query参数。uri绑定动态参数
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}

//绑定json
func bindJson(c *gin.Context) {
	//实例化结构体
	var wang Userinfo

	//将前端传来的json与wang绑定
	c.ShouldBindJSON(&wang)

	//以json响应输出wang以验证
	c.JSON(200, gin.H{"json": wang})
}

//绑定查询参数
func bindQuery(c *gin.Context) {
	//实例化结构体
	var wang Userinfo

	//将前端传来的query与wang绑定
	c.ShouldBindQuery(&wang)

	//响应输出wang以验证
	c.JSON(200, gin.H{"json": wang})
}

//uri uniform resouce identifier 同一资源标识符，通常用于动态参数
func bindUri(c *gin.Context) {
	//实例化结构体
	var wang Userinfo

	c.ShouldBindUri(&wang)

	//响应输出wang以验证
	c.JSON(200, gin.H{"json": wang})
}

//根据content-type进行匹配,例如form-data，tag用form（默认tag）
func bindAll(c *gin.Context) {
	//实例化结构体
	var wang Userinfo

	c.ShouldBind(&wang)

	//响应输出wang以验证
	c.JSON(200, gin.H{"json": wang})
}

func main() {
	e := gin.Default()

	e.POST("/json", bindJson)
	e.POST("/query", bindQuery)
	e.POST("/uri/:name/:age/:sex", bindUri)
	e.POST("/all", bindAll)

	e.Run(":8072")

}
