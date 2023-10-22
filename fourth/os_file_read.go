package main

import (
	"fmt"
	"io"
	"os"
)

// 打开文件，打开文件之后就可以进行c.read,c.close等方法
func open() *os.File {
	c, _ := os.Open("D:/goLang/os/尝试.txt") //底层为os.openfile
	return c

	//os.creat等价于os.openfile(,o_rdwr|o_creat|o_trunic,0666)
}

func read(c *os.File) {

	for { //这里先读再进行判断
		buf := make([]byte, 10) //设置一个5字节的缓冲区
		n, err := c.Read(buf)
		fmt.Printf("buf: %v\n", string(buf)) //string（slice）将切片转换为字符串
		fmt.Printf("n: %v\n", n)
		if err == io.EOF { //表示文件读取完毕
			c.Close() //关闭文件，否则可能造成系统资源的浪费
			break
		}
	}

	buf := make([]byte, 10)
	n, _ := os.Open("D:/goLang/os/成功.txt")
	x, _ := n.ReadAt(buf, 3) //从第3个字符开始读，0 1 2 3 4
	fmt.Printf("x: %v\n", x)
	fmt.Printf("ATbuf: %v\n", string(buf))
	n.Close()

	//读取目录,使用遍历的方法
	//这里也可以使用先open在读取的操作
	dir, _ := os.ReadDir("D:/goLang")
	for key, v := range dir {
		fmt.Printf("%v: %v\n", key, v.Name())
	}
}

func write() {

}

func main() {
	read(open())
}
