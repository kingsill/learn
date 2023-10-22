package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//发出HTTP/ HTTPS请求。
	r, err := http.Get("http://www.baidu.com")
	// r, err := http.Get("127.0.0.1:8000")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	//程序在使用完回复后必须关闭回复的主体。
	defer r.Body.Close()
	//使用io。readall进行读取
	b, _ := io.ReadAll(r.Body)
	fmt.Printf("b: %v\n", string(b))
	//循环读取，直到读完所有内容
	// buf := make([]byte, 1024)
	// var temp string
	// for {
	// 	n, err := r.Body.Read(buf)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	temp += string(buf[:n])
	// }
	// fmt.Printf("r.Body: %v\n", r.Body)
	// fmt.Printf("response: %v\n", temp)

	fmt.Printf("r.Header: %v\n", r.Header)
}
