package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func Sendfile(path string, c net.Conn) {
	//以只读方式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("open file failed ,err: %v\n", err)
		return
	}

	defer f.Close()

	//读文件内容，读多少发多少
	buf := make([]byte, 1024)
	for {
		// 循环读取，每次最多1024个字
		n, err2 := f.Read(buf)
		//发送内容
		fmt.Printf("buf[:n]: %v\n", string(buf[:n]))
		c.Write(buf[:n])
		if err2 == io.EOF {
			fmt.Println("发送完毕")
			return
		}
	}

}

func main() {
	//提示输入文件
	fmt.Println("请输入需要传输的文件：")
	var path string

	//可以回顾scan的各种类型的区别，见 https://blog.csdn.net/weixin_45765795/article/details/112766580
	fmt.Scan(&path)
	fmt.Printf("path: %v\n", path)

	//获取文件名
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	//主动连接服务器
	c, err2 := net.Dial("tcp", "127.0.0.1:20000")
	if err2 != nil {
		fmt.Printf("dial failed,err2: %v\n", err2)
		return
	}
	fmt.Println("连接成功！")
	defer c.Close()

	//给接收方先发送文件名
	_, err3 := c.Write([]byte(fi.Name()))
	if err3 != nil {
		fmt.Printf("write failed ,err3: %v\n", err3)
		return
	}

	//接受对方的回复。如果回复“ok”，说明准备好了
	buf := make([]byte, 1024)
	n, err4 := c.Read(buf)
	fmt.Printf("buf: %v\n", string(buf))

	//上文中的写法与这里有什么区别？
	// r := bufio.NewReader(c)
	// n3, err5 := r.Read(buf)

	if err4 != nil {
		fmt.Printf("read failed,err4: %v\n", err4)
		return
	}
	if string(buf[:n]) == "ok" {
		fmt.Println("接收方已准备好")
		//发送文件内容
		Sendfile(path, c)
	} else {
		fmt.Println("对面没有返回 ok ，请重新准备")
	}
}
