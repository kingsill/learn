package main

//普通函数与方法的区别（在接收者分别为值类型和指针类型的时候）

import (
	"fmt"
)

// 1.普通函数
// 接收值类型参数的函数
func valueIntTest(a int) int {
	return a + 10
}

// 接收指针类型参数的函数
func pointerIntTest(a *int) int {
	return *a + 10
}

func structTestValue() {
	a := 2
	fmt.Println("valueIntTest:", valueIntTest(a))
	//函数的参数为值类型，则不能直接将指针作为参数传递
	//fmt.Println("valueIntTest:", valueIntTest(&a))
	//compile error: cannot use &a (type *int) as type int in function argument

	b := 5
	fmt.Println("pointerIntTest:", pointerIntTest(&b))
	//同样，当函数的参数为指针类型时，也不能直接将值类型作为参数传递
	//fmt.Println("pointerIntTest:", pointerIntTest(&b))
	//compile error:cannot use b (type int) as type *int in function argument
}

func main() {
	structTestValue()
}
