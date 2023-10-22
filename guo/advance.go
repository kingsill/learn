package main

import "fmt"

func main() {
	//高阶函数，返回函数----------------------------------------------------------------------------
	d := cal("+")
	e := d(1, 2)
	fmt.Printf("%d", e)

}
func sub(a int, b int) int {
	return a - b
}
func add(a int, b int) int {
	return a + b
}
func cal(c string) func(int, int) int { //这里返回的即为函数类型，即func (int ,int )int 类型的函数
	switch c {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}
