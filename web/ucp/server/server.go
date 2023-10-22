package main

import (
	"fmt"
	"net"
)

func main() {
	//启动服务，进行监听
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
		Zone: "",
	})
	if err != nil {
		fmt.Printf("listen failed,err: %v\n", err)
		return
	}
	defer listen.Close()

	//无连接，不需要goroutine来进行多个连接的处理
	for {
		buf := make([]byte, 1024)
		_, u, err2 := listen.ReadFromUDP(buf)
		if err2 != nil {
			fmt.Printf("read failed,err2: %v\n", err2)
		}
		fmt.Printf("接收到的数据: %v\n", string(buf))
		listen.WriteToUDP(buf, u)

	}
}
