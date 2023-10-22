package main

import (
	"net/http"

	"github.com/gin-gonic/gin" //引入框架
)

func getUser(c *gin.Context) {
	username := c.PostForm("username")            //获得username
	userage := c.DefaultPostForm("userage", "18") //如果没有的话，就给一个默认值18
	usergender := c.PostForm("usergender")        //获得usergender

	c.JSON(http.StatusOK, gin.H{"username": username, "userage": userage, "usergender": usergender}) //返回json数据
} //返回json数据
func main() {
	router := gin.Default() //设置路由

	router.GET("/user", getUser)

	router.Run() //设置运行接口
}
