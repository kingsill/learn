package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段
// 先建立接受（<-chan）再发送(chan<-)
var values = make(chan int)

func main() {
	defer close(values)

	go send()
	fmt.Println("wait...")
	value := <-values //接受
	fmt.Printf("receive: %v\n", value)
	fmt.Println("END")
}

func send() {
	rand.Seed(time.Now().UnixNano()) //rand.seed（） 根据括号中的值，对应会得到不同的序列树，每个相同的值序列树都是相同的
	value := rand.Intn(10)
	// fmt.Printf("time.Now(): %v\n", time.Now())
	fmt.Printf("send: %v\n", value)
	values <- value //发送
}
