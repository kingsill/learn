package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func receiveFile(filename string, conn net.Conn) {
	//创建文件
	fmt.Printf("filename: %v\n", filename)
	f, err := os.Create(filename)

	if err != nil {
		fmt.Printf("create failed,err: %v\n", err)
		return
	}
	defer f.Close()
	//接受多少，写入多少内容
	buf := make([]byte, 1024)
	for {
		n, err2 := conn.Read(buf)
		fmt.Printf("buf: %v\n", string(buf[:n]))

		if err2 != nil {
			if err2 == io.EOF {
				fmt.Println("读取完毕")
				return
			}

		}
		f.Write(buf[:n])
	}

}

func main() {
	//监听串口
	l, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed ,err: %v\n", err)
		return
	}

	//阻塞等待用户连接
	c, err2 := l.Accept()
	if err2 != nil {
		fmt.Printf("accept failed,err2: %v\n", err2)
		return
	}
	fmt.Println("连接成功")
	defer c.Close()

	//读取对方发送的文件名
	buf := make([]byte, 1024)

	r := bufio.NewReader(c)
	n, err3 := r.Read(buf)
	if err3 != nil {
		fmt.Printf("read failed,err3: %v\n", err3)
		return
	}
	fileName := string(buf[:n])
	fmt.Printf("fileName: %v\n", fileName)

	//返回ok.
	if n > 0 {
		c.Write([]byte("ok"))
		//介绍文件内容
		receiveFile(fileName, c)
	}

}
