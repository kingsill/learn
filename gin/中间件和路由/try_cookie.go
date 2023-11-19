package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	e.GET("/cookie", func(c *gin.Context) {
		//获取cookie
		s, err := c.Cookie("username")
		if err != nil {
			s = "wzl"
			//设置cookie，这里的域名需要与网址一致，可能与SameSite有关
			c.SetCookie("username", s, 60*60, "/", "127.0.0.1", false, true)
			c.String(200, "cookie")
		}
	})

	e.Run()
}
