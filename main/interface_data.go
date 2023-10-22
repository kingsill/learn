// 值接收者和指针接收者实现接口的区别
package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

//func (stu *Student) Speak(think string) (talk string) {
func (stu Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}

}

func main() {
	var peo People = &Student{} //值接收者可以直接输入指针类型，反之则不行
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
