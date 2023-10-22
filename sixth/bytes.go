// bytes包提供了对字节切片进行读写操作的一系列函数，
// 字节切片处理的函数比较多分为
// 基本处理函数、比较函数、后缀检查函数、索引函数、分割函数、大小写处理函数和子切片处理函数等。
package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 强制转换 包含
func test1() {
	var i int = 100
	b := byte(i)
	fmt.Printf("b: %v\n", b)

	u := "hello world"
	b1 := []byte(u)
	b2 := []byte("Hello world")
	b3 := []byte("hello world")

	fmt.Printf("b2: %v\n", b1)
	fmt.Printf("bytes.Contains(b2, b1): %v\n", bytes.Contains(b2, b1))
	fmt.Printf("bytes.Contains(b3, b1): %v\n", bytes.Contains(b3, b1))
	fmt.Printf("bytes.Contains(b3, []byte(\"h\")): %v\n", bytes.Contains(b3, []byte("h")))

	fmt.Printf("strings.Contains(u, \"h\"): %v\n", strings.Contains(u, "h"))
}

// 计数	count
func test2() {
	b := []byte("helloooooooo")
	b2 := []byte("h")
	b3 := []byte("l")
	b4 := []byte("o")

	s := string(b)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("strings.Count(s, \"o\"): %v\n", strings.Count(s, "o"))

	fmt.Printf("bytes.Count(b, b2): %v\n", bytes.Count(b, b2))
	fmt.Printf("bytes.Count(b, b3): %v\n", bytes.Count(b, b3))
	fmt.Printf("bytes.Count(b, b4): %v\n", bytes.Count(b, b4)) //计数b中b4数量
}

// 重复 repeat
func test3() {
	b := []byte("hi")
	s := string(b)

	fmt.Printf("bytes.Repeat(b, 3): %s\n", bytes.Repeat(b, 3)) //重复输出3次
	fmt.Printf("strings.Repeat(s, 2): %v\n", strings.Repeat(s, 2))
}

// 替换 replace
func test4() {
	b := []byte("hello world")
	old := []byte("o")
	new := []byte("O")

	fmt.Printf("bytes.Replace(s, old, new, 1): %s\n", bytes.Replace(b, old, new, 1))
	//同义写法
	fmt.Printf("bytes.Replace(s, old, new, 1): %v\n", string(bytes.Replace(b, old, new, 1)))
	//如果n小于0，则无限次转换
	fmt.Printf("bytes.Replace(s, old, new, 1): %s\n", bytes.Replace(b, old, new, -1))
}

// runes
func test5() {
	b := []byte("你的世界")
	fmt.Printf("b: %v\n", b)
	r := bytes.Runes(b)
	fmt.Printf("len(b): %v\n", len(b)) //12
	fmt.Printf("len(r): %v\n", len(r)) //4
}

func main() {
	test1()
	// test2()
	// test3()
	// test4()
	// test5()
}
