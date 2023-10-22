package main

import "github.com/gin-gonic/gin"

type SignUserInfo struct {
	Name       string   `json:"name" binding:"required"`
	Age        int      `json:"age" binding:"lt=30,gt=18" `
	Sex        string   `binding:"oneof=woman man"` //枚举 只能是woman或man
	Password   string   `json:"password"`
	Repassword string   `json:"repassword" binding:"eqfield=Password"`                      //注意这里绑定的是字段名
	Hobby      []string `json:"like list" binding:"required,dive,required,startswith=like"` //从第dive后面开始就是针对数组内的内容进行校验
	IP         string   `binding:"ip"`
	Time       string   `binding:"datetime=2006-01"` //2006-01-02 15:04:05 是一个参考时间的格式,这里后面所跟一定要从参考时间格式截取
}

//验证
func verify(c *gin.Context) {
	//实例化结构体
	var wang SignUserInfo

	//将内容绑定到结构体
	err := c.ShouldBindJSON(&wang)
	if err != nil {
		c.JSON(200, gin.H{"meg": err.Error()})
		return
	}
	c.JSON(200, "ok")

}

func main() {
	e := gin.Default()

	e.POST("/verify", verify)

	e.Run(":8020")
}
