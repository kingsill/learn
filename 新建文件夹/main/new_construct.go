// 模拟构造函数 构造函数
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person, error) { //Person 的指针
	//产生错误的函数会返回两个变量，一个值和一个错误码；如果后者是 nil 就是成功，非 nil 就是发生错误
	if name == "" {
		return nil, fmt.Errorf("姓名不可以为空") //fmt.errorf使用与printf使用类似，但是用与产生错误对象
	}
	if age < 0 {
		return nil, fmt.Errorf("年龄不可以小于0")
	}
	return &Person{name: name, age: age}, nil //返回Person 的地址，与上述对应

}

func main() {
	Tom, err := NewPerson("Tom", -1)
	if err == nil {
		fmt.Println(Tom)
	} else {
		fmt.Println(err)
	}
}
