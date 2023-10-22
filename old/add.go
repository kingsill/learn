package main

import "fmt"

func main() {
	fmt.Println("请输入要相加的两个数字")
	var x, y int
	fmt.Scan(&x, &y)
	z := x * y
	fmt.Printf("%d*%d=%d", x, y, z)
}
