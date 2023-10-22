package main

import "fmt"

func main() {
	//指针------------------------------------------------------------------
	var ip *int
	fmt.Printf("ip:%T", ip)
	c := 1
	ip = &c
	fmt.Printf("ip: %v\n", ip)
	d := *ip
	fmt.Printf("d: %v\n", d)
	//指向数组的指针   var name [MAX]8int  max指数组内元素的个数
	a := [3]int{1, 2, 3}
	var pp [3]*int
	fmt.Printf("pp: %v\n", pp)
	for k, _ := range a {
		pp[k] = &a[k]
	}
	fmt.Printf("pp: %v\n", pp)
	*pp[0] = 10
	fmt.Printf("a: %v\n", a)
}
