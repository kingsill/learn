package main

func main() {

	/* 	a := 10----------------------------------------------------------------------------------
	   	b := 6
	   	c := a % b
	   	fmt.Printf("c: %v\n", c)
	   	var x, y, z int
	   	fmt.Scan(&x, &y, &z)
	   	fmt.Printf("x: %v\n", x)
	   	fmt.Printf("y: %v\n", y)
	   	fmt.Printf("z: %v\n", z)
	   	if x == 1 && y == 2 && z == 3 {
	   		fmt.Println("ok!")
	   	} else {
	   		fmt.Println("wrong!")
	   	} */

	//Grade----------------------------------------------------------------------------------
	/* 	fmt.Println("Please input your mark:")
	   	var mark int
	   	fmt.Scan(&mark)
	   	if mark >= 90 {
	   		fmt.Println("Your grade is: A")
	   	} else if mark >= 80 {
	   		fmt.Println("Your grade is: B")
	   	} else if mark >= 70 {
	   		fmt.Println("Your grade is: C")
	   	} else if mark >= 60 {
	   		fmt.Println("Your grade is: D")
	   	} else {
	   		fmt.Println("Sorry,you are fail! ")
	   	} */

	//Compare----------------------------------------------------------------------------------
	/* 	fmt.Println("Please input 3 numbers, I'll out put the biggest one")
	   	var x, y, z int
	   	fmt.Scanln(&x, &y, &z)
	   	if x >= y { //x>y?z
	   		if x > z { //x>y>z
	   			fmt.Printf("MAX: %v\n", x)
	   		}
	   	} else if y > z { //x<y>z
	   		fmt.Printf("MAX: %v\n", y)
	   	} else { //x<y<z
	   		fmt.Printf("MAX: %v\n", z)
	   	} */

	//for range----------------------------------------------------------------------------------
	/* 	var c = []int{1, 2, 3, 4, 5, 6}
	   	for _, b := range c {
	   		fmt.Printf("b: %v\n", b)
	   	}
	   	var a = [...]int{1, 2, 3, 4}
	   	for _, v := range a {
	   		fmt.Println(v)
	   	} */

	//map
	/* 	m := map[string]string{
	   		"wang": "handsome",
	   		"jian": "silly",
	   	}
	   	fmt.Println(m["wang"])//map中索引
	   	for c, v := range m {
	   		fmt.Printf("c: %v\n", c)
	   		fmt.Printf("v: %v\n", v)
	   	} */

	/* LOOP1: //break退出整个循环
	   	for i := 0; i < 10; i++ {
	   		fmt.Printf("i: %v\n", i)
	   		if i == 3 {
	   			break LOOP1
	   		}
	   	}
	   	fmt.Println("start")
	   	fmt.Println("end")

	   	//break、loop也可以不用标签
	   	for i := 0; i < 10; i++ {
	   		fmt.Printf("i: %v\n", i)
	   		if i == 3 {
	   			break
	   		}
	   	}
	   	fmt.Println("start")
	   	fmt.Println("end")

	   LOOP2: //continue退出单词循环
	   	for i := 0; i < 10; i++ {

	   		if i == 3 {
	   			continue LOOP2
	   		}
	   		fmt.Printf("i: %v\n", i)
	   	}
	   	fmt.Println("start")
	   	fmt.Println("end") */

	//go to 直接跳转、代码可读性降低、可能留下隐患
	//goto不会打破循环，如果标签放在前面会重复当次循环
	/* 	for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {

				if i == 3 && j == 3 {
					goto END
				}
				fmt.Printf("i: %v\n", i)
				fmt.Printf("j: %v\n", j)
			}
		}
	END:
		fmt.Println("end") */

	//数组----------------------------------------------------------------------------------
	/* 	var a1 [3]int
	   	fmt.Printf("%T", a1)
	   	var a2 []int
	   	fmt.Printf("a2: %T\n", a2)
	   	var a3 = [...]int{1, 2, 3} //数组使用。。。来表示容量的时候必须使用=
	   	fmt.Printf("c: %v\n", len(a3))
	   	fmt.Printf("a3[1]: %v\n", a3[1])
	   	fmt.Println(a3)
	   	for k, v := range a3 { //遍历数组
	   		fmt.Printf("a3[%d]: %v\n", k, v)
	   	} */

	// 切片，底层为数组，但是长度不固定----------------------------------------------------------------------------------
	/* 	//最好使用make函数来进行初始化
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
	   	fmt.Printf("d: %v\n", d) */

	// go语言特性，-----------------------------------------------------------------
	// 函数分为三种，普通函数、 匿名函数、 方法，函数本身也是一种数据
	/* 	a := 3
	   	fmt.Printf("a: %v\n", a)//注意在函数中是否使用引用类型的数据
	   	d(2) */
	/* 	func d(a int) int {//func中为形参，不影响外部参数
		a = 1
		fmt.Printf("a: %v\n", a)
		return a
	} */

	//使用。。。来充当函数类型中的个数---------------------------------------------------------------------
	/* 	D(1, 2, 3, 4)
	   	func D(args ...int) { //名称使用args，这里使用。。。来忽略参数个数，相当于创建切片，后面将值传给该函数，切片自动扩容
	   		for _, v := range args {
	   			fmt.Printf("v: %v\n", v)
	   		}
	   		fmt.Printf("args: %T\n", args)
	   	} */

	//高阶函数，返回函数----------------------------------------------------------------------------
	/* 	d := cal("+")
	   	e := d(1, 2)
	   	fmt.Printf("%d", e)

	   	func sub(a int, b int) int {
	   		return a - b
	   	}
	   	func add(a int, b int) int {
	   		return a + b
	   	}
	   	func cal(c string) func(int, int) int { //这里返回的即为函数类型，即func (int ,int )int 类型的函数
	   		switch c {
	   		case "+":
	   			return add
	   		case "-":
	   			return sub
	   		default:
	   			return nil
	   		}
	   	} */

	//匿名函数，在函数中定义某个变量等于函数，直接在定义函数后使用（）赋值-------------------------------------------------
	//闭包 简单理解为函数+引用
	/* 	f1 := add()
	   	fmt.Printf("f1(10): %v\n", f1(10)) //x=0 y=10 >> x=x+y=10
	   	fmt.Printf("f1(10): %v\n", f1(10)) //x=10,y=10 >> x=x+y=20
	   	//x的值得到保留，重新引用时x重新回到定义
	   	f2 := add()
	   	fmt.Printf("f2(10): %v\n", f2(10))

	   	func add() func(y int) int {
	   		var x int
	   		return func(y int) int {
	   			x += y//闭包的关键即在于在内部定义函数中使用在外部函数中定义的变量，使该变量得到保留
	   			return x
	   		}
	   	} */

	//递归 函数自己调用自己，需要定义好退出条件，可能产生大量goroutine，产生栈空间内存溢出

	//init函数  init函数先于main函数自动执行
	//变量初始化》》init函数》》main函数
	/* 	func init() { //自动运行，无需引用
		fmt.Println("init")
	} */

	//指针------------------------------------------------------------------
	/* 	var ip *int
	   	fmt.Printf("ip:%T", ip)
	   	c := 1
	   	ip = &c
	   	fmt.Printf("ip: %v\n", ip)
	   	d := *ip
	   	fmt.Printf("d: %v\n", d)
	   	//指向数组的指针   var name [MAX]8int  max指数组内元素的个数
	   	a := [3]int{1, 2, 3}
	   	var pp [3]*int
	   	fmt.Printf("pp: %v\n", pp)
	   	for k, _ := range a {
	   		pp[k] = &a[k]
	   	}
	   	fmt.Printf("pp: %v\n", pp)
	   	*pp[0] = 10
	   	fmt.Printf("a: %v\n", a) */

	//
}
