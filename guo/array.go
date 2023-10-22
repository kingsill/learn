package main

import "fmt"

func main() {
	//数组----------------------------------------------------------------------------------
	var a1 [3]int
	fmt.Printf("%T", a1)
	var a2 []int
	fmt.Printf("a2: %T\n", a2)
	var a3 = [...]int{1, 2, 3} //数组使用。。。来表示容量的时候必须使用=
	fmt.Printf("c: %v\n", len(a3))
	fmt.Printf("a3[1]: %v\n", a3[1])
	fmt.Println(a3)
	for k, v := range a3 { //遍历数组
		fmt.Printf("a3[%d]: %v\n", k, v)
	}

}
