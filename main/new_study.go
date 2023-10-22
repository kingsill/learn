// 方法 与struct/其他数据类型 关联   特殊类型函数
package main

import "fmt"

type Person struct {
	name     string
	password string
	age      int
}

var w = Person{
	name:     "wang",
	age:      18,
	password: "123456",
}

func main() {
	Person.eat(w)
	var u Person

	fmt.Println("请输入用户名和密码")
	fmt.Printf("用户名：\t")
	fmt.Scanln(&u.name)
	fmt.Printf("密码：\t")
	fmt.Scanln(&u.password)

	if u.login() {
		u.inspect() //调用时，函数名前需加 数据名。
	} else {
		fmt.Println("用户名或密码错误，请重新输入！")
	}
}

func (per Person) login() bool { //检验登录用户名和密码
	if per.name == w.name && per.password == w.password {
		return true
	} else {
		return false
	}
}
func (per Person) inspect() { //验证年龄
	fmt.Println("登录成功！请输入你的年龄：\t")
	fmt.Scanln(&per.age)
	if per.age >= 18 {
		fmt.Println("恭喜你登录成功")
	} else {
		fmt.Println("小孩子不准访问哦")
	}
}
func (per Person) eat() { //与Person结构体关联,per为参数名
	fmt.Printf("%s想吃饭\n", per.name)
}
