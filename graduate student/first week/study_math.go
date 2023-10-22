// math包包含一些常量和一些有用的数学计算函数，例如，三角函数、随机数、绝对值、平方根等
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 常量Π、float32的最大值
func test1() {
	fmt.Printf("math.Pi: %v\n", math.Pi)

	fmt.Printf("math.Pi: %.3v\n", math.Pi) //对于%v来说，精度为所有数字的总数
	fmt.Printf("math.Pi: %.3f\n", math.Pi) //对于%f来说，精度为小数点后面数字的总数

	fmt.Printf("math.MaxFloat32: %v\n", math.MaxFloat32)
}

// 常用函数
func test2() {

	//math.abs()取绝对值
	fmt.Printf("math.MinInt16: %v\n", math.MinInt16)
	fmt.Printf("math.Abs(math.MinInt16): %v\n", math.Abs(math.MinInt16))

	//取x的y次方 math.pow（x，y）		power function 幂函数
	fmt.Printf("3的2次方: %v\n", math.Pow(3, 2))

	//开平方 sqrt；开立方	cbrt
	fmt.Printf("math.Sqrt(9): %v\n", math.Sqrt(9))
	fmt.Printf("math.Cbrt(8): %v\n", math.Cbrt(8))

	//向上取整 ceil	；向下取整 floor；四舍五入 round
	fmt.Printf("math.Round(5.5): %v\n", math.Round(5.5))

	//分别取整数和小数部分
	int2, frac := math.Modf(3.14)
	fmt.Printf("小数部分为：%v,整数部分为：%v\n", frac, int2)
	fmt.Printf("小数部分为：%.f,整数部分为：%.f\n", frac, int2) //%.f直接将后面舍去

	//取余
	f := math.Mod(4, 3) //4/3...1
	fmt.Printf("f: %v\n", f)
}

func init() {
	// rand.Seed(1) //若使用该句，则所有随机都是基于1产生的，会产生相同的随机数
	rand.Seed(time.Now().UnixNano()) //使用时间作为依据以实现伪随机
}

func test3() {
	for i := 0; i < 10; i++ {
		fmt.Printf("rand.Int: %d\n", rand.Int())
	}

	//从0-100进行随机整数
	for i := 0; i < 10; i++ {
		fmt.Printf("rand.Intn(100): %v\n", rand.Intn(100))
	}

	//从0-1之间随机小数
	for i := 0; i < 10; i++ {
		fmt.Printf("rand.Float32(): %v\n", rand.Float32())
	}
}

func main() {
	// test1()
	// test2()
	test3()
}
