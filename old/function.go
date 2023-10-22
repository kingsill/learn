package main //函数本身为一种数据类型）（指针）

import "fmt"

func main() {
	f1(1, 2)
	var f4 func(a, b int)
	f4 = f1         //可以使用f4:= f1
	fmt.Println(f1) //直接使用f1即输出地址，说明函数本身为一种指针变量
	fmt.Println(f4)
	f4(2, 3)
}
func f1(a, b int) {
	fmt.Println(a, b)
}

/* 结果：
1 2
0x85de60
0x85de60
2 3
*/

//f1、f4所指的内存地址相同，存储的为{}中的代码
