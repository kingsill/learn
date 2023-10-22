package main

import "fmt"

func add(nums ...int) int {
	var num int
	for i := 0; i < len(nums); i++ {
		num = num + nums[i]
	}
	return num
}
func main() {
	fmt.Println(add(1, 2, 3, 4))
}
