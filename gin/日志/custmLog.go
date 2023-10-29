package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.New()

	//gin.LoggerWithConfig()也可以起到相同作用，查看源码可解
	e.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string { //也可以将这里封装为一个函数

		// return "[wang] " + params.ClientIP + " |" + params.TimeStamp.Format("2006-01-02 15:04:05") + " |" + strconv.Itoa(params.StatusCode) + " |" + params.Method + " |" + params.Path
		//strconv.Itoa 将int转变为string
		//上下两种写法输出结果一致，fmt.sprintf 返回拼接的字符串
		return fmt.Sprintf(
			"[wang] %s | %d |%s |%s\n",
			params.TimeStamp.Format("2006-01-02 15:04:05"),
			params.StatusCode,
			params.Method,
			params.Path,
		)
	}))
	fmt.Printf("\033[97;41m 红底白字    \033[0m  正常颜色 ") //vscode 的控制台输出是看不到颜色的
	e.GET("/try")

	e.Run()
}
