package main

import (
	"fmt"
	"os"
)

func creat() {
	adr, err := os.Create("D:/goLang/os/尝试.txt") //创建txt文件
	//正常如果函数运行成功，err一般为nil，否则为错误信息，代表执行失败
	if err != nil {
		fmt.Println("err!")
	} else {
		fmt.Printf("adr: %T\n", adr)
	}
}

func makeDir() {
	err := os.Mkdir("D:/goLang/os/test/deep", os.ModePerm) //创建目录
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println("ok!")
	}
	//删除目录（也可以用来删除文件）
	os.Remove("D:/goLang/os/test/deep")
	os.Remove("D:/goLang/os/尝试.txt")

}

func wd() {
	dir, _ := os.Getwd() //获取当前工作目录
	fmt.Printf("dir: %v\n", dir)

	os.Chdir("d:/") //修改工作目录
	dir, _ = os.Getwd()
	fmt.Printf("dir: %v\n", dir)

	//获取临时工作目录
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

func rename() {
	err := os.Rename("D:/goLang/os/尝试.txt", "D:/goLang/os/成功.txt") //rename（old，new）
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}

// 读取文件
func Read() {
	c, err := os.ReadFile("D:/goLang/os/尝试.txt")
	fmt.Printf("read_err: %v\n", err)
	fmt.Printf("c: %v\n", c)
}

// 读取文件
func write() {
	err := os.WriteFile("D:/goLang/os/成功.txt", []byte("hello"), os.ModePerm) //这里[]byte("hello")实际为将hello从string转换为slice
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Printf("write_err: %v\n", err)
	}
}

func main() {
	// creat()
	// makeDir()
	// wd()
	// rename()
	Read()
	write()
}
