// 并发 创建协程即在函数前面+go
// go 协程意味着并行，通过通道来通信
// 无缓冲chan只有在双方都就绪后才能通信，
// 无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段
// 所以需要先创建协程接受数据，再在main函数中发送一个数据
package main

import (
	"fmt"
	"time"
)

func show(x string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("x: %v\n", x)
		time.Sleep(time.Millisecond * 100)
		// 为了避免第一个协程执行过快，观察不到并发的效果，加个休眠
	}
}

func main() {
	/* 	show("小明")    //协程随着main函数的结束而结束，这样写的话什么都不会输出
	   	go show("你好") */
	go show("你好")
	show("小明")
	time.Sleep(time.Second) //使main阻塞

}
