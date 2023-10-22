// mutex 互斥锁
// 保证只有一个goroutine可以访问共享资源
package main

import (
	"fmt"
	"sync"
)

var i int = 100
var wg sync.WaitGroup
var lock sync.Mutex

func sub() {
	defer wg.Done()
	lock.Lock() //从当前lock开始对i加锁，防止其他协程读写
	i--
	fmt.Printf("i--: %v\n", i)
	// time.Sleep(time.Microsecond * 10)
	lock.Unlock() //从当前开始解锁
}
func add() {
	defer wg.Done()
	lock.Lock()
	i++
	fmt.Printf("i++: %v\n", i)
	// time.Sleep(time.Microsecond * 10)
	lock.Unlock()
}

// 不适用mutex顺序不能确定
func main() {
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go add()
		go sub()

	}
	wg.Wait()
	fmt.Println("end")
	//当前只能保证结果为100？

}
