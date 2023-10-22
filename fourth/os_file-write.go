package main

import (
	"fmt"
	"os"
)

/*
- os.O_RDONLY：只读模式，打开文件后只能读取文件内容，不能修改文件内容。
- os.O_WRONLY：只写模式，打开文件后只能写入文件内容，不能读取文件内容。
- os.O_RDWR：读写模式，打开文件后既可以读取文件内容，也可以写入文件内容。
- os.O_CREATE：如果文件不存在，则创建文件。
- os.O_APPEND：追加模式，打开文件后写入的内容会追加到文件末尾。
- os.O_trunk:覆盖
*/
func write() {
	//这里的如果使用open进行打开的话权限不够
	w, _ := os.OpenFile("D:/goLang/os/尝试.txt", os.O_RDWR|os.O_APPEND, 777) //这里代表这里的openfile命令可以进行读写，同时新输入的数据从后面进行插入
	n, err := w.Write([]byte("heelo\n"))                                   //将hello转变为切边输入
	fmt.Printf("n: %v\n", n)
	fmt.Printf("err: %v\n", err)

	w.WriteString("新的来喽")
	w.Close()
}

func writeat() {
	//使用writeat时不能再openfile命令中加入append
	w, _ := os.OpenFile("D:/goLang/os/尝试.txt", os.O_RDWR, 0777)
	_, er := w.WriteAt([]byte("QAQ"), 30)
	fmt.Printf("er: %v\n", er)
	w.Close()
}

func main() {
	write()
	writeat()
}
