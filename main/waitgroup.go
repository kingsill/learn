// 等待组内部有一个计数器
// 调用wait的相关函数会对其进行修改
// wp.done()   计数-1
// wp.add(int)  计数+int
// wp.wait()   直到计数为0前都会阻塞
package main

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

func show(x int) {
	defer wp.Done() //每次协程完成时计数-1
	fmt.Printf("x: %v\n", x)
}

func main() {

	for i := 0; i < 10; i++ {
		go show(i)
		wp.Add(1) //每次创建协程都将计数+1
	}
	wp.Wait() //当计数为0时不再阻塞main函数，等待所有协程完成
	fmt.Println("end")
}
