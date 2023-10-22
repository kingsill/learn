// 最重要的即为reader和writer两个接口
// os包中read为方法，实现reader接口，以此类推
// strings.reader
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func str() {
	r := strings.NewReader("hello") //strings.NewReader可以床在一个reader接口
	buf := make([]byte, 10)
	r.Read(buf)
	fmt.Printf("string(buf): %v\n", string(buf))
}

func copy() {
	r := strings.NewReader("world\n")
	//io.copy自动实现读写操作
	_, err := io.Copy(os.Stdout, r) //os.stdout自动输出到控制台
	if err != nil {
		log.Fatal(err) //fatal这里等于print
	}
	fmt.Printf("err: %v\n", err)
}

func bufcopy() {
	a := strings.NewReader("3,2,1")
	b := strings.NewReader("ok!")
	buf := make([]byte, 8)

	io.CopyBuffer(os.Stdout, a, buf)
	// fmt.Printf("buf: %v\n", buf)
	io.CopyBuffer(os.Stdout, b, buf) //buf无需重复调用
}
func main() {
	copy()
	str()
	bufcopy()
}
