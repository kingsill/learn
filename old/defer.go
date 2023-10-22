package main

import "fmt"

func main() { //序号代表执行次序
	pr(1)       //一
	defer pr(2) //七
	pr(3)       //二
	defer pr(4) //六
	pr(5)       //三
	x := 6
	defer pr(x) //五，此时6值已经传递进入
	x++
	fmt.Println(x) //四
}

func pr(s int) {
	fmt.Println(s)
}
