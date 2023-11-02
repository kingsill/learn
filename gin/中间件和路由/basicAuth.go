package main

import "github.com/gin-gonic/gin"

//basicAuth是简答的验证功能

//模拟存储私人信息到这里
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

//BasicAuth好像不可以通过额外路由
// func test_BasicAuth(c *gin.Context) {
// 	//BasicAuth返回一个基本http中间件，接受一个map[string]string作为参数
// 	gin.BasicAuth(gin.Accounts{
// 		"foo":    "nihao",
// 		"austin": "1234",
// 		"lena":   "hello2",
// 		"manu":   "4321",
// 	})
// }

func test_BasicAuthsecrets(c *gin.Context) {
	//获取用户名，它是由BaiscAuth中间件设置的
	user := c.MustGet(gin.AuthUserKey).(string)
	if secret, ok := secrets[user]; ok {
		c.JSON(200, gin.H{
			"user":   user,
			"secret": secret,
		})
	}
}

func main() {
	e := gin.Default()

	//BasicAuth只能单独在这里进行注册
	rg := e.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "nihao", //这里设置的为用户名和密码
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	{
		rg.GET("/secrets", test_BasicAuthsecrets)
	}

	e.Run()

}
