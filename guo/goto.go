package main

import "fmt"

func main() {
	//go to 直接跳转、代码可读性降低、可能留下隐患
	//goto不会打破循环，如果标签放在前面会重复当次循环
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {

			if i == 3 && j == 3 {
				goto END
			}
			fmt.Printf("i: %v\n", i)
			fmt.Printf("j: %v\n", j)
		}
	}
END:
	fmt.Println("end")

}
