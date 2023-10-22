// golang模拟OOP 面向对象
// 添加接收者 receiver
// Person类，具有name、age两个属性，sleep、eat、read三种功能,通过方法来实现
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (per Person) sleep() {
	fmt.Printf("%s正在sleep", per.name)
}

func (per Person) eat() {
	fmt.Printf("%s正在eat", per.name)
}

func (per Person) read() {
	fmt.Printf("%s正在read\n", per.name)
}

func main() {
	per := Person{
		name: "wang",
		age:  18,
	}
	per.eat()
	wang := Wang{
		per,
	}
	wang.sleep() //wang.c.sleep
}

//继承 通过结构体嵌套来进行
type Wang struct {
	Person //c Person   这里使用匿名方式的话调用简单
}
