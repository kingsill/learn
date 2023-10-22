package main //九九乘法表

import "fmt"

func main() {
	for i := 9; i > 0; i-- {
		for j := i; j > 0; j-- {
			x := i * j
			fmt.Printf("%d*%d=%d\t", i, j, x)
		}
		fmt.Println()
	}
}
