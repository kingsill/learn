package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Println("connect")
	//
	defer conn.Close()
	requestHeader := "GET /go HTTP/1.1\r\nHost: 127.0.0.1:8000\r\nConnection: keep-alive\r\nCache-Control: max-age=0\r\nUpgrade-Insecure-Requests: 1\r\nUser-Agent: Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3\r\nAccept-Encoding: gzip, deflate, \r\nAccept-Language: zh-CN,zh;q=0.9,ja;q=0.8,en;q=0.7\r\n\r\n"

	// 先发请求包，服务器才会回响应包
	_, writeErr := conn.Write([]byte(requestHeader))
	fmt.Println("request")
	if err != nil {
		fmt.Println("writeErr =", writeErr)
	}

	// 接收服务器回复的响应包
	buf := make([]byte, 1024*4)
	n, readErr := conn.Read(buf)
	if readErr != nil {
		if readErr != nil {
			fmt.Println("readErr =", readErr)
		} else {
			fmt.Println("null response")
		}
		return
	}

	//
	responseStr := string(buf[:n])
	fmt.Printf("#%v#", responseStr)
}
