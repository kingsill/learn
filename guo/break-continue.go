package main

import "fmt"

func main() {
LOOP1: //break退出整个循环
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i == 3 {
			break LOOP1
		}
	}
	fmt.Println("start")
	fmt.Println("end")

	//break、loop也可以不用标签
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i == 3 {
			break
		}
	}
	fmt.Println("start")
	fmt.Println("end")

LOOP2: //continue退出单词循环
	for i := 0; i < 10; i++ {

		if i == 3 {
			continue LOOP2
		}
		fmt.Printf("i: %v\n", i)
	}
	fmt.Println("start")
	fmt.Println("end")
}
