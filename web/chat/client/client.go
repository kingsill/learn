package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//1.与服务端连接
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Printf("dial fialed,err: %v\n", err)
		return
	}
	fmt.Println("与", "127.0.0.1:20000", "连接成功")

	//3.从服务器端接受回复的消息
	go func() {
		var buf [1024]byte //类比server.go中办法，也可以直接建立切片
		for {
			reader := bufio.NewReader(conn)
			n, err2 := reader.Read(buf[:]) //数组后面加[:]即复制一个内容相同的切片，[]的用法可查
			if err2 != nil {
				fmt.Printf("read failed,err2: %v\n", err2)
				return
			}
			fmt.Printf("%v\n", string(buf[:n]))
		}
	}()

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
