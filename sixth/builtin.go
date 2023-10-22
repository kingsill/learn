// 此包不需要导入，可以直接使用
package main

import (
	"fmt"
	"log"
)

func main() {
	test1()
	test2()
	// test3()
	test4()
}

// append    附加    连接切片
func test1() {
	s := []int{1, 2, 3}
	i := append(s, 100)
	fmt.Printf("i: %v\n", i)
	i2 := []int{4, 5, 6}
	i3 := append(i, i2...)
	fmt.Printf("i3: %v\n", i3)
}

// len   返回 数组、切片、字符串、通道 的长度
func test2() {
	u := "hello"
	fmt.Printf("len(u): %v\n", len(u))

	i := []int{1, 2, 3, 10}
	fmt.Printf("len(i): %v\n", len(i))

	c := make(chan int, 10)
	c <- 12
	c <- 23
	fmt.Printf("len(c): %v\n", len(c))
}

// print、println

// panic
func test3() {
	defer fmt.Println("followed by panic") //表面是先输出这句话，但其实是panic后进行的
	log.Panic(123)
	fmt.Println("end?")
}

//new、make
/* make只能用来分配及 初始化 类型为 slice map chan 的数据， new可以分配任意类型的数据
new分配返回的是指针，即类型*T;make返回引用，即本身，T
new分配的空间被清零，make分配后，会进行初始化 */
func test4() {
	i := new([]int) //这里不用赋予容量，内容为空
	fmt.Printf("i: %T\n", i)
	fmt.Printf("i: %v\n", *i)

	//习惯用法
	i2 := make([]int, 10) //这里初始化容量为10，并且进行清零
	fmt.Printf("i2: %T\n", i2)
	fmt.Printf("i2: %v\n", i2)
}
