package main

import (
	"fmt"
	"strings"
)

// 第二题
func devide() {
	fmt.Println("%d", 153%10)
	fmt.Printf("%v\n", 153/10) //golang中使用除法，都是整数则返回整数
	fmt.Printf("%f", 153.0/10) //若有一方为小数，则返回小数
}

func main() {
	s := "pwwkew"
	n := reapeat(s)
	fmt.Printf("n: %v\n", n)
}

// 第三题 烂
func reapeat(s string) int {
	n := 0
	np := 0
	b := []byte(s)
	var key []byte

	for _, v := range b {
		n++
		s2 := string(key)
		if strings.Contains(s2, string(v)) {

			var keyp []byte
			for _, v := range b[n-1:] {
				np++
				s3 := string(keyp)

				if strings.Contains(s3, string(v)) {
					return max(n-1, np-1)
				}
				fmt.Printf("v: %v\n", v)
				fmt.Printf("np: %v\n", np)
				keyp = append(keyp, string(v)...)
				fmt.Printf("s3: %v\n", s3)
			}

		}
		key = append(key, string(v)...)
		fmt.Printf("key: %v\n", s2)
	}
	return n

}
func max(num1, num2 int) int {

	if num1 > num2 {
		return num1
	} else {
		return num2
	}

}
