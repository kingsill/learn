// 接口 只关心内部方法 关联方法进行理解
package main

import "fmt"

type USB interface { //接口定义
	read()
	write()
}

type MUSIC interface {
	play()
}
type BUS interface { //接口嵌套
	USB
	MUSIC
}

type computer struct {
	name string
}

type phone struct {
	name string
}

func (c computer) read() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("reading")
}

func (c computer) write() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("writing")
}
func (c computer) play() {
	fmt.Println("播放音乐")
}

func main() { //多态实现方法
	/* 	var usb USB //定义接口
	   	usb = computer{} //具有对应方法的结构体赋予接口
	   	usb.read()	//直接通过接口进行函数的输出
	*/

	var usb USB //结合上面的注释进行理解
	c := computer{"wang"}
	usb = c
	usb.read()
	c.play()
	var bus BUS //嵌套内的都可以进行引用S
	bus = c     //同样适用于具有对应方法的结构体赋予接口
	bus.write()
}
