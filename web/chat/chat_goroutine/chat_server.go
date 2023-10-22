package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//-----所标记的地方 意指与原文不同

type client struct {
	Name string      //用户名
	Addr string      //网络地址
	C    chan string //用于发送数据的管道
}

// type cliAddr string //为了方便理解，讲cliAddr定义为string

// var onlinemap = make(map[string]client) //用户信息对应 map为引用类型，不初始化分配空间不能填充数据-------
// map+structure时由于map中的value本身是不可寻址的，所以将结构体改为指针类型
var onlinemap = make(map[string]*client) //用户信息对应 map为引用类型，不初始化分配空间不能填充数据-------
// var onlinemap map[string]client

var message = make(chan string) //make的原因与 map同理，引用类型--------

// 专门发送信息
func WriteMessageToClient(cli client, conn net.Conn) {
	for msg := range cli.C { //给当前客户端发送信息
		conn.Write([]byte(msg))

		fmt.Printf("msg: %v\n", msg)
	}
}

// 打包封装msg
func MakeMsg(cli client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ":" + msg
	return
}

// 处理用户连接
func handleConn(conn net.Conn) {
	// defer conn.Close()

	//获取客户端的网络地址
	addr := conn.RemoteAddr().String()

	fmt.Printf("addr: %v\n", addr)

	//创建用户对应的结构体
	cli := client{
		Name: addr,
		Addr: addr,
		C:    make(chan string),
	}

	fmt.Printf("cli: %v\n", cli)

	isQuit := make(chan bool)  //用于确认用户是否退出
	hasData := make(chan bool) //用于确认对方是否有数据发送

	//将结构体加入到map中
	// onlinemap[addr] = cli
	onlinemap[addr] = &cli

	fmt.Printf("onlinemap: %v\n", onlinemap)

	//新开一个协程，专门给当前客户端发送信息
	go WriteMessageToClient(cli, conn)

	//广播用户上线
	// message <- "[" + cli.Addr + "]" + cli.Name + ":login"
	message <- MakeMsg(cli, "login")

	//提示我是谁
	cli.C <- "I am here,addr:" + addr

	//新建一个协程，接受用户发送的数据
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)

			if err != nil {
				isQuit <- true
				fmt.Printf("read failed,err: %v\n", err)
				return
			}

			// if n == 0 {
			// 	isQuit <- true
			// 	fmt.Printf("read failed,err: %v\n", err)
			// }

			msg := string(buf[:n])

			if msg == "who" { //查询当前在线用户
				//遍历map
				conn.Write([]byte("user list:"))
				for _, cli := range onlinemap {
					info := cli.Name + cli.Addr
					message <- info
					// conn.Write([]byte(info))
				}
			}

			if len(msg) > 8 && msg[:6] == "rename" && strings.Contains(msg, "|") { //重命名
				cli.Name = strings.Split(msg, "|")[1]
				onlinemap[addr].Name = strings.Split(msg, "|")[1]
			}

			fmt.Printf("Msg: %v\n", msg)
			//转发此内容
			message <- MakeMsg(cli, msg)

			hasData <- true //代表有数据进入
		}
	}()

	for {
		select {
		case <-isQuit:
			delete(onlinemap, addr)        //删除用户
			message <- MakeMsg(cli, "已退出") //广播下线
			return
		case <-hasData:
		case <-time.After(60 * time.Second): //常用select+time.after来实现超时控制
			delete(onlinemap, addr)               //删除用户
			message <- MakeMsg(cli, "超时未发送信息，踢出") //广播下线
			return
		}
	}
}

// 转发消息，遍历map，给map每个成员发送通道内消息
func Manage() {
	//缺少给map分配空间这一步，留待后续添加-------------

	for {
		msg := <-message //没有消息时，这里会堵塞
		for _, client := range onlinemap {
			client.C <- msg
		}
	}

}

func main() {
	//监听
	l, _ := net.Listen("tcp", ":8000") //当ip地址为空时，默认使用回环地址
	defer l.Close()

	//创建用于转发消息的协程，只要有消息，即进行遍历map转发
	go Manage()

	//主协程，循环阻塞、等待用户连接
	for {
		c, err2 := l.Accept()
		if err2 != nil {
			fmt.Printf("accept failed,err2: %v\n", err2)
			continue //不使用return是因为return会结束整个程序
		}
		fmt.Println("连接成功")
		go handleConn(c)
	}

}
