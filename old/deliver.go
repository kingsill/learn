package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4} //数组定义方式
	deliver(arr)
	fmt.Println(arr)
}

func deliver(arr2 [4]int) {
	fmt.Println(arr2)
	arr2[2] = 100
	fmt.Println(arr2)
}
