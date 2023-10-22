// runtime.gosched 让出当前时间片给其他协程
// runtime.goexit 直接结束当前协程
// runtime.gomaxprocs 默认使用最多的cpu核心数量
// go的协程为抢占式，会争夺运行权
package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Printf("a: %v\n", i)
		runtime.Gosched()
	}
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Printf("b: %v\n", i)
		runtime.Gosched()
	}
}

func show(meg string) {
	for i := 1; i < 10; i++ {
		fmt.Printf("meg: %v,time:%v\n", meg, i)
		if i >= 5 {
			runtime.Goexit() //退出当前协程
		}
	}

}

func main() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU()) //runtime.numcou（）查询电脑核心数（逻辑处理器数量）

	// runtime.GOMAXPROCS(3)

	// go show("go")
	go a()
	go b()
	for i := 0; i < 8; i++ {
		// runtime.Gosched() //schedule 让出当前的cpu时间片
		fmt.Println("让出")
	}
	time.Sleep(time.Second * 3) //等待三秒
	fmt.Println("end")
}
