package main

import "fmt"

func main() {
	s := "Hello,World!"
	for i, v := range s { //i为数组中每个元素的序号，v为每个元素的值
		fmt.Printf("%d\t", i) //i的类型为int，需要用%d来进行输出
		fmt.Printf("%c\t", v) //v的类型为int32，但是为ASCII编码，需要%c的符号来进行输出
	}
}
