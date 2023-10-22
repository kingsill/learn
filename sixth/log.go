//内置log包，实现简单的日志服务
/* print系列 单纯打印日志
panic系列 打印日志，抛出panic异常
fatal系列 打印日志，强制结束程序（os.Exit（1）），defer函数不会执行 */

package main

import (
	"fmt"
	"log"
	"os"
)

// print	printf	println
func test1() {
	log.Print("nihao")
	log.Printf("%d", 12)
	name := "wang"
	log.Println(name, "", 2)
}

// panic 第一个panic之后的部分不再执行，但是defer会执行
func test2() {
	defer log.Print("end")
	log.Panic(321)
	log.Panic("panic")
}

// 与panic相比不再执行defer
func test3() {
	defer log.Print("end")
	log.Fatal(321)
	log.Fatal("fatal")
}

//---------------------------------------------------------------------------------

// flag	标准日志配置
func test4() {
	log.SetFlags(11) //设置标准输出配置，这里将日期，时间，文件路径控制输出日志
	// log.SetFlags(log.Ldate | log.Ltime | log.Llongfile) 	等同于上一句，具体可查看源码，此处为位掩码形式

	i := log.Flags()
	fmt.Printf("i: %v\n", i) //这里可以查看当前的输出配置，或者初始化配置

	log.Print(123)
}

// flag 日志前缀配置
func test5() {
	log.SetPrefix("error:")
	log.Print("nihao")
}

// 日志输出位置配置,将日志输出到文件
func test6() {
	f, err := os.OpenFile("a.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 777) //"a.txt"也可以
	defer f.Close()
	if err != nil {
		log.Panic("open failed")
	}
	log.SetOutput(f)
	log.Print("my log:")
}

// 自定义logger
// 将标准日志配置、日志前缀、日志输出位置整合到一个函数中
func test7() {
	f, err := os.OpenFile("b.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 777)
	defer f.Close()
	if err != nil {
		log.Panic("open err")
	}

	l := log.New(f, "new log", 11)
	l.Print("new world")
}

func main() {
	test1()
	test2()
	test3()

	// test4()
	// test5()
	// test6()
	// test7()
}
