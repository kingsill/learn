package main

import (
	"fmt"
	"os"
)

func main() {
	//os.args 获取在命令行输入对应程序后的内容，程序名加后续内容
	//例如 ： go run get_config.go a.txt
	list := os.Args
	fmt.Printf("list: %v\n", list)
	fmt.Printf("list: %T\n", list)

	//如果在命令行的输入只有程序名，而没有别的指令，则报错
	if len(list) != 2 {
		fmt.Println("usage:xxx file")
		return
	}

	fileName := list[1]
	//Stat返回一个描述name指定的文件对象的FileInfo。如果指定的文件对象是一个符号链接，返回的FileInfo描述该符号链接指向的文件的信息，本函数会尝试跳转该链接。如果出错，返回的错误值为*PathError类型。
	fi, err := os.Stat(fileName)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("name=%v\n", fi.Name())
	fmt.Printf("size=%v\n", fi.Size())

}
