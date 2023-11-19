// 如果通道关闭，仍能够读取其中的内容，如果没有则返回0值
// 但如果通道没有关闭，则读取会造成死锁
// 另一种遍历方法直接使用for
package main

import "fmt"

var c = make(chan (int))

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {

		fmt.Printf("v: %v\n", v)
		// if v == 9 {//此处和close（c）需保留一个，防止发生死锁
		// 	break

		// }
	}

}
