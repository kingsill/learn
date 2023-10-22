// 一个通道只对一个协程开放
// 发送、接受等操作在完成前会阻塞
// 在接受者接受数据之前，数据发送不会结束，双方是互相阻塞的
package main

//死锁 通道要输出，但是内部没数据；通道要写入，但是里面有数据
import (
	"fmt"
	"time"
)

var ch chan int

// var ch = make(chan int)//也可以这么说

func main() {
	ch = make(chan int, 10)
	go in(ch)
	go out(ch)
	//通道只能在函数内部建立，但是可以在外部进行声明
	//channel<-data通道接受数据

	time.Sleep(100 * time.Millisecond)
}

func out(in chan int) {
	for {
		fmt.Printf("%v", <-in)
	}
	/* 	defer func() {

		for i := 0; i < 10; i++ {
			fmt.Printf("%v", <-in)
		}
	}() */

}
func in(in chan int) {
	for i := 0; i < 10; i++ {
		in <- i
	}
}
