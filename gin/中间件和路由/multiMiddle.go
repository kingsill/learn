package main

//单个中间件+拦截+时间统计
import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	//这里跟在每个路由的后面的函数可以视为单独的中间件，并按照顺序进行执行
	e.GET("/try", m1, m2, m3)

	e.Run()
}

//c.Next() 在其之前的就是请求中间件， 在其之后的就是响应中间件
//next可以类似栈来理解或者defer函数

//abort可以理解为放弃运行后续的中间件，直接进入当前的中间件的响应并继续响应之前的中间件

func m1(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok"})

	startTime := time.Now() //增加计时功能，可以获得该路由的运行时间
	fmt.Printf("startTime: %v\n", startTime)

	fmt.Println("1 in")
	c.Next()
	c.Abort() //可以看到这里的abort没有作用，是因为等m1响应的时候所有中间件已经运行，无法停止下一个中间件
	fmt.Println("1 out")

	sinceTime := time.Since(startTime) //获取从start之后过去的时间
	fmt.Println(sinceTime)
}

func m2(c *gin.Context) {

	fmt.Println("2 in")
	c.Abort() //这里的abort放弃了m3中间件的运行，因此m3的请求和响应都失效了，直接进入m2的响应
	c.Next()

	fmt.Println("2 out")
	c.JSON(200, gin.H{"hello": "maybe"})
}

func m3(c *gin.Context) {
	fmt.Println("3 in")
	c.Next()
	fmt.Println("3 out")

}
