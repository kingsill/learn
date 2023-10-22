package main

import "fmt"

func main() {
	// 切片，底层为数组，但是长度不固定----------------------------------------------------------------------------------
	//最好使用make函数来进行初始化
	c := make([]int, 0) //make([]tyoe,长度，容量)
	//从尾部插入
	c = append(c, 1, 2, 3)   //使用append来进行插入，这里是最常用的从尾部插入
	fmt.Printf("c: %v\n", c) //append为附加，可以理解为连接
	//从头部插入
	c = append([]int{6, 7, 8}, c...) //这里格式类似与连接切片，将{6，7，8}作为数组插入到c的前面再赋给c
	fmt.Printf("c: %v\n", c)
	//连接切片，后面的切片需加...
	a := []int{9, 10}
	c = append(c, a...)
	fmt.Printf("c: %v\n", c)
	//从中间插入
	c = append(c[:2], append([]int{4, 5}, c[2:]...)...) //可以看作是切片{4，5}先和c[2:]连接，这里切片连接需加。。。，
	fmt.Printf("c: %v\n", c)                            //然后将这部分再与c[：2]接，再加。。。
	fmt.Printf("c[1:5]: %v\n", c[1:5])                  //切片的切割遵循左闭右开，c[:2]={6,7}即对应0，1，不含2
	//copy函数，需要数组初始化时赋予足够的长度，铭记切片是引用类型的数据
	b := make([]int, 0)
	fmt.Printf("copy(c, b): %v\n", copy(b, c)) //使用copy是不会自动扩容，copy返回的是复制的元素个数
	fmt.Printf("b: %v\n", b)
	d := make([]int, 10) //这里直接初始化容量
	fmt.Printf("d: %v\n", d)
	fmt.Printf("copy(c, d): %v\n", copy(d, c))
	fmt.Printf("%T\n", copy(d, c)) //可以看到copy返回的类型为整数型，d已经复制完成
	fmt.Printf("d: %v\n", d)

	//使用。。。来充当函数类型中的个数---------------------------------------------------------------------
	D(1, 2, 3, 4)

}
func D(args ...int) { //名称使用args，这里使用。。。来忽略参数个数，相当于创建切片，后面将值传给该函数，切片自动扩容
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
	fmt.Printf("args: %T\n", args)
}
