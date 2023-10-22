// newscanner需要通过reader进行建a立
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	scan()
}
func scan() {
	r := strings.NewReader("123 456 你好！")
	s := bufio.NewScanner(r)

	//Splitner的分割函数
	//本方法必须在scan之前调用
	s.Split(bufio.ScanWords) //以 空 格 作为 分隔符 进行分割
	// s.Split(bufio.ScanBytes) //以 字符 int8 作为 分隔符 进行分割
	// s.Split(bufio.ScanRunes) // 以 字符 int32 作为分割符 进行分割 （可以显示汉字等语言）

	for s.Scan() { //扫描完返回false，没扫描完返回true
		fmt.Printf("s.Text(): %v\n", s.Text())
	}
}
