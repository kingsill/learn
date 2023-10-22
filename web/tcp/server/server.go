package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func process(conn net.Conn) {

	//针对当前的连接做数据的发送和接受操作
	fmt.Printf("conn.LocalAddr(): %v\n", conn.LocalAddr())

	a := conn.RemoteAddr()
	fmt.Printf("conn.RemoteAddr(): %v\n", a)
	defer fmt.Println("连接关闭", a)
	defer conn.Close() //处理完之后关闭这个连接
	for {

		reader := bufio.NewReader(conn)
		// var buf [128]byte
		buf := make([]byte, 1024*4)
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Printf("read failed,err: %v\n", err)
			break
		}
		if string(buf[:n]) == "exit" {
			defer fmt.Println("服务器断开")
			return
		}
		fmt.Printf("%v: %v\n", a, string(buf[:n]))
		fmt.Printf("#%v#", string(buf[:n]))
		// conn.Write(buf[:n]) //把收到的数据返回给客户端
	}
}

// 为服务端增加发送消息功能
func send(conn net.Conn) {
	defer conn.Close() //处理完之后关闭这个连接
	//2.利用连接进行数据的发送和接受
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)

		//如果读到Q，则退出
		if strings.ToUpper(s) == "Q" {
			return
		}
		_, err3 := conn.Write([]byte(s))
		if err3 != nil {
			fmt.Printf("write failed,err3: %v\n", err3)
			return
		}
	}
}

func main() {
	//1.启动服务，启动监听端口
	l, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed,err: %v\n", err)
	}

	for {
		//2.等待客户端来连接
		conn, err2 := l.Accept()
		if err2 != nil {
			fmt.Printf("accept failed,err2: %v\n", err2)
			continue
		}
		fmt.Println("connect!")
		//3.启动一个单独的goroutine来处理连接
		go process(conn)
		send(conn)

	}

}
