// readstring只能读有缓冲的reader
// 有缓存的reader	1.先建立reader	2.建立有缓存reader 	3.读取	or	先建立另一个缓存再读取
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 读字符串
func read() {
	r := strings.NewReader("hello world") //打开一个reader接口
	r2 := bufio.NewReader(r)              //创建一个具有默认缓存大小(4096)的缓冲，从r读取的reader
	s, _ := r2.ReadString('\n')           //读取r2，并且在读到\n即换行时停止
	fmt.Printf("s: %v\n", s)
}

// 读文件
func fileread() {
	f, _ := os.Open("D:/goLang/os/尝试.txt")
	r3 := bufio.NewReader(f)
	n, _ := r3.ReadString('\n') //读到第一次换行
	fmt.Printf("n: %v\n", n)
	f.Close()
}

// reset 置零并导向
func rst() {
	//将 数字 放入 r	；将 字母 放进r2	并读取
	r := strings.NewReader("123456")
	r2 := strings.NewReader("abcdefg")

	//创建从r读取的有缓存reader r3
	r3 := bufio.NewReader(r)
	//r4 := bufio.NewReader(r2)

	//用s读reader r3中的数据
	s, _ := r3.ReadString('\n')
	fmt.Printf("s: %v\n", s) //打印r3

	//清零r3的缓存，并将r3读取的reader从r改为r2
	r3.Reset(r2)

	//再用s读reader r3 接口的数据
	s, _ = r3.ReadString('\n')
	fmt.Printf("s: %v\n", s) //打印r3

}

// 双缓冲的读取
func read_buf() {
	r := strings.NewReader("ABCDEFGHI123456")
	r2 := bufio.NewReader(r) //建立有缓冲的reader接口
	buf := make([]byte, 10)
	for {
		n, err := r2.Read(buf) //通过缓存进行第一步读取
		if err == io.EOF {
			break
		} else {
			fmt.Printf("buf: %v\n", string(buf[0:n]))
		}
	}
}

// unreadbyte反读上一个字节，不要重复使用
func rtr() {
	r := strings.NewReader("WANG")
	r2 := bufio.NewReader(r)

	b, _ := r2.ReadByte() //每次读取一个字符
	fmt.Printf("b: %v\n", string(b))

	b, _ = r2.ReadByte()
	fmt.Printf("b: %v\n", string(b))

	r2.UnreadByte()
	b, _ = r2.ReadByte()
	fmt.Printf("b: %v\n", string(b))

} //readrune与之类似，可以用来读取日文等字符，后续有需要可以对比学习

//readslice\readbytes(delim byte )在读到delim字节时停止，类似于readstring，使用时再进行学习
//readbytes 返回 切片类型，readstring 返回 字符串

// write to，实现了io.writer to 接口
func write() {
	r := strings.NewReader("ABCDEFG!")
	r2 := bufio.NewReader(r)
	// b := bytes.NewBuffer(make([]byte, 0))

	//写入
	// r2.WriteTo(b)
	// fmt.Printf("b: %v\n", b)

	//写入文件
	w, _ := os.OpenFile("D:/goLang/os/尝试.txt", os.O_RDWR|os.O_APPEND, 777)
	r2.WriteTo(w)
	w.Close()
}

func main() {
	read()
	fileread()
	rst()
	read_buf()
	rtr()
	write()
}
