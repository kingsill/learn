package main

//单个中间件+拦截
import (
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	//这里跟在每个路由的后面的函数可以视为单独的中间件，并按照顺序进行执行
	e.GET("/try", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "ok"})
		s := c.Request.Header.Get("user-agent")

		//c.Abort() 可以阻止调用挂起的处理程序
		if matched, _ := regexp.MatchString(`Postman`, s); matched {
			//这里举得例子是如果检查到是postman发送的请求，则阻塞在这一步
			c.Abort()
		}

	}, func(c *gin.Context) {
		fmt.Println("1")
		c.JSON(200, gin.H{"hello": "maybe"})
	}, func(c *gin.Context) {
		fmt.Println("2")
	})

	e.Run()
}
