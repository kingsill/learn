// bufio_read中最后一部分也是写入
// 建立有缓存writer	1.先创建一个buffer（可以实现reader）	2.通过buffer建立writer	3.写入	4.fluh or writeTo
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	wrt()
	reset()
	RDfrom()
	RDWT()
}

// 带缓存的写入
func wrt() {
	w, err := os.OpenFile("D:/goLang/os/尝试.txt", os.O_RDWR|os.O_APPEND, 777) //打开文件
	fmt.Printf("err: %v\n", err)
	f := bufio.NewWriter(w) //创建writer接口
	f.WriteString("hello")
	f.Flush() //刷新缓冲区，不写这个的话写入不进去
	w.Close()

}

// 清除缓冲中数据,并充值导向，可以参考read中的reset
func reset() {
	b := bytes.NewBuffer(make([]byte, 0)) //创建了一个buffer，可以实现writer接口
	w := bufio.NewWriter(b)               //将writer接口转换成带有缓冲的接口
	w.WriteString("Hnadsome")             //将handsome写入缓存

	w.Flush() //将handsome写入b

	b2 := bytes.NewBuffer(make([]byte, 0))
	w.Reset(b2) //将接口缓存清空，并将b的下层writer改为b2

	w.WriteString("agly") //将agly写入缓存

	w.Flush() //刷新接口缓冲区，将缓冲中数据写入下层的io.writer接口

	fmt.Printf("b: %v\n", b)
	fmt.Printf("b2: %v\n", b2)
}

func RDfrom() {
	b := bytes.NewBuffer(make([]byte, 0))
	w := bufio.NewWriter(b)
	fmt.Printf("w.Available(): %v\n", w.Available()) //默认情况下可利用的缓冲区大小
	fmt.Printf("w.Buffered(): %v\n", w.Buffered())   //当前已经利用的缓冲区大小

	w.WriteString("wang")
	fmt.Printf("w.Available(): %v\n", w.Available())
	fmt.Printf("w.Buffered(): %v\n", w.Buffered()) //可以得到这里接口已经利用的缓存大小为4
	fmt.Printf("b: %v\n", b)

	w.Flush() //清空w的缓存，并写入到b

	fmt.Printf("b: %v\n", b)
	fmt.Printf("w.Available(): %v\n", w.Available())
	fmt.Printf("w.Buffered(): %v\n", w.Buffered()) //已经利用的缓存重新回到0
	// ----------------------------------------------------------------------------------
	b2 := bytes.NewBuffer(make([]byte, 0))
	w2 := bufio.NewWriter(b2)

	r := strings.NewReader("gogogo")

	w2.ReadFrom(r) //使用readfrom可以无需flush，自动写入
	fmt.Printf("b2: %v\n", b2)
}

//writerune 等同int32 、writebyte等同 int8，需要使用时咋进行相关学习

// readerwriter可以实现io。readwriter的接口
func RDWT() {
	//建立一个有缓存writer
	b := bytes.NewBuffer(make([]byte, 0)) //
	r := bufio.NewWriter(b)

	//建立一个有缓存reader
	r2 := strings.NewReader("yes")
	r3 := bufio.NewReader(r2)

	//建立一个readerwriter可以实现两种接口功能
	rw := bufio.NewReadWriter(r3, r)

	//读
	s, _ := rw.ReadString('e')
	fmt.Printf("s: %v\n", s)
	//写
	rw.WriteString("shit!")
	rw.Flush()
	fmt.Printf("b: %v\n", b)
}
