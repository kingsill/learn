package main

import (
	"fmt"
	"net/http"
)

// w：给客户端回复数据    req：读取客户端发送的数据
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
	fmt.Printf("r.Method: %v\n", r.Method)
	fmt.Printf("r.URL: %v\n", r.URL)
	fmt.Printf("r.Header: %v\n", r.Header)
}

func main() {
	//注册处理函数，用户连接，自动调用指定处理函数
	http.HandleFunc("/go", myHandler)

	//在指定的地址进行监听，开启一个http
	http.ListenAndServe("127.0.0.1:8000", nil)
}
