package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//指定具体的存储类型，这里为cookie
	store := cookie.NewStore([]byte("secret")) //设置用于身份验证或加密的密钥
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		//获取session
		if session.Get("hello") != "world" {
			session.Set("hello", "world")

			session.Save() //保存session数据，即使session持久化，这里是保存到cookie
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run(":8000")
}
