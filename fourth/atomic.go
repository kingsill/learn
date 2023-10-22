//atomic与mutex锁相同，为了防止多个协程对某个变量同时更改，
//原子操作可以理解为变量级别的互斥锁
//性能更高，保护的从代码块变为变量

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

var i int32 = 100
var cb int32 = 1

func main() {
	for c := 0; c < 200; c++ {

		ok := atomic.CompareAndSwapInt32(&i, 100, 0) //如果i==100，换为0，返回true；反之值不变，返回false
		if ok {
			//第一次运行时为ok，运行此部分代码
			wg.Add(2)

			go add()
			atomic.StoreInt32(&cb, i) //存储，将i存储到cb
			fmt.Printf("cb: %v\n", cb)

			go sub()
			wg.Wait()
			fmt.Printf("i: %v\n", i)

		} else {
			//第二次运行时，i的值为0，false，运行else部分代码
			fmt.Println("err")
			fmt.Printf("i: %v\n", i)
			break

		}
	}
}

// cas compare and swap       old  new
func add() {
	defer wg.Done()
	atomic.AddInt32(&i, 1) //自加
}

func sub() {
	defer wg.Done()
	atomic.AddInt32(&i, -1) //自减
}
