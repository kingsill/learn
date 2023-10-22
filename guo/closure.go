package main

import "fmt"

func main() {
	//匿名函数，在函数中定义某个变量等于函数，直接在定义函数后使用（）赋值-------------------------------------------------
	//闭包 简单理解为函数+引用
	f1 := add()
	fmt.Printf("f1(10): %v\n", f1(10)) //x=0 y=10 >> x=x+y=10
	fmt.Printf("f1(10): %v\n", f1(10)) //x=10,y=10 >> x=x+y=20
	//x的值得到保留，重新引用时x重新回到定义
	f2 := add()
	fmt.Printf("f2(10): %v\n", f2(10))

}
func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y //闭包的关键即在于在内部定义函数中使用在外部函数中定义的变量，使该变量得到保留
		return x
	}
}
