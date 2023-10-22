package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

// 获得请求头的各种方法
func header(c *gin.Context) {
	//不区分大小写，单词中间用 - 隔开
	fmt.Printf("c.GetHeader(\"user-agent\"): %v\n", c.GetHeader("user-agent")) //getheader方法只能获取一个请求头

	//header是一个MAP[string][]string类型数据
	c.JSON(200, gin.H{"header": c.Request.Header}) // 获取header中的全部内容

	//也可以用c.request.header.get()来获取单个值，自动匹配大小写
	///通过c.request.header[xxx]，即通过使用map查询的话。需要大小写完全一致
}

// 区别对待用户和爬虫的方法
func crawerDetect(c *gin.Context) {
	s := c.GetHeader("user-agent")
	//第一种方法是通过正则表达式进行匹配
	//第二种方式是通过strings包的包含来进行匹配(正则校验)

	if strings.Contains(s, "Postman") { //由于这里用的是postman，这里使用postman代替常用的python
		//s中包含python字符，说明是爬虫

		//返回空值
		c.JSON(200, gin.H{"data": "爬虫给爷爬！"})
		return
	}

	//正则校验的第一种方式
	matched, _ := regexp.MatchString(`python`, "python") //regexp.matchstring第一个参数为正则表达式，第二个为校验的字符
	if matched {
	}

	//正则校验的第二种方式
	r := regexp.MustCompile("python") //首先通过正则表达式匹配构造regexp.ExpReg
	if r.MatchString("python") {      //通过上述的结构体的方法进行校验
	}

	//欢迎用户
	c.JSON(200, gin.H{"data": "欢迎阿龙"})
	regexp.MustCompile("python")
}

// 响应头设置
func responseHead(c *gin.Context) {
	//用c.header以设置响应头
	c.Header("Token", "jhgeu%hsg845jUIF83jh")
	c.Header("Content-Type", "application/text; charset=utf-8")
	c.JSON(0, gin.H{"data": "看看响应头"})
}

func main() {
	e := gin.Default()

	e.GET("/header", header)
	e.GET("/crawer", crawerDetect)
	e.GET("/response", responseHead)

	e.Run(":8071")
}
