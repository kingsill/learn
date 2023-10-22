package main

import "fmt"

func main() {
	//递归 函数自己调用自己，需要定义好退出条件，可能产生大量goroutine，产生栈空间内存溢出
	//init函数  init函数先于main函数自动执行
	//变量初始化》》init函数》》main函数
}
func init() { //自动运行，无需引用
	fmt.Println("init")
}
