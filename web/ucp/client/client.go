package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
		Zone: "",
	})
	if err != nil {
		fmt.Printf("dial failed,err: %v\n", err)
		return
	}
	defer conn.Close()
	r := bufio.NewReader(os.Stdin)
	for {

		s, _ := r.ReadString('\n')
		_, err2 := conn.Write([]byte(s))
		if err2 != nil {
			fmt.Printf("send failed,err2: %v\n", err2)
		}
		buf := make([]byte, 1024)

		_, _, err3 := conn.ReadFromUDP(buf)
		if err3 != nil {
			fmt.Printf("接受回复失败,err3: %v\n", err3)
			return
		}
		fmt.Printf("收到回复: %v\n", string(buf))
	}
}
