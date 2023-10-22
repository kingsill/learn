// open-closed principle 对扩展允许，但不进行修改
// 方法里调用别的接口，完成相关功能
package main

import "fmt"

type Pet interface { //定义Pet接口
	eat()
	sleep()
}

type Dog struct {
}
type Cat struct {
}

//Dog实现Pet
func (Dog) sleep() {
	fmt.Println("dog sleep")
}
func (Dog) eat() {
	fmt.Println("dog eat")
}

//Cat实现Pet
func (Cat) sleep() {
	fmt.Println("cat sleep")
}
func (Cat) eat() {
	fmt.Println("cat eat")
}

type Person struct {
	name string
}

func (Person) care(p Pet) { //care函数调用Pet接口
	p.sleep()
	p.eat()
}
func main() {
	wang := Person{"wang"}
	wang.care(Cat{})
	//cat:=Cat{}8
	//wang.care(cat)
}
