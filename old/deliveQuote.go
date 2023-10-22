package main //切片为引用类型的变量，指向的地址类型相同（函数本身也是引用类型的数据）
import "fmt" /*值类型的数据：int、string、bool、float64、array
引用类型的数据：slice、map、chan*/

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Println("初始参数", s)
	update(s)
	fmt.Println("调用后参数", s)
}

func update(x []int) {
	fmt.Println("传递函数", x)
	x[0] = 100
	fmt.Println("修改完传递函数", x)
}
