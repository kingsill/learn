package main

import (
	"fmt"
)

func main() {
	fmt.Println("请输入要比较的两个数")
	var a, b int
	fmt.Scanln(&a, &b)
	fmt.Printf("max=%d", max(a, b))
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
