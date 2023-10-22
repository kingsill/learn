package main

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name   string `binding:"required,sign" msg:"请输入用户名"`
	Age    int    `binding:"gte=18" msg:"未成年人禁止访问"`
	Gender string `binding:"oneof=woman man" msg:"请从woman和man中选择输入"`
}

func custm(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(200, gin.H{"msg": GetValidMsg(err, &user)})
		// return
	}
	c.JSON(200, gin.H{"msg": user})
}

func GetValidMsg(err error, obj any) string {
	// 使用的时候，需要传obj的指针
	getObj := reflect.TypeOf(obj)

	var tt string

	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错字段名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				tt = tt + msg + "  "

			}
		}
		return tt
	}
	return err.Error()
}

func signValid(fl validator.FieldLevel) bool {
	var nameList []string = []string{"fengfeng", "wang", "zhang3"}
	name := fl.Field().Interface().(string)

	for _, v := range nameList {
		if name == v {
			return false
		}
	}

	return true
}

func main() {
	e := gin.Default()

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("sign", signValid)
	}
	e.GET("/custm", custm)

	e.Run(":8010")

}
