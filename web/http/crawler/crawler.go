package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//发出HTTP/ HTTPS请求。
	var name string
	var page, page2 int
	fmt.Println("请输入要爬取的贴吧名，开始页数和结束页数，中间用空格隔开")
	fmt.Scan(&name, &page, &page2)

	for i := page; i < page2+1; i++ {

		n := (i - 1) * 50
		// fmt.Printf("n: %v\n", n)
		s := strconv.Itoa(n)

		website := "https://tieba.baidu.com/f?kw=" + name + "&ie=utf-8&pn=" + s
		r, err := http.Get(website)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		//程序在使用完回复后必须关闭回复的主体。
		defer r.Body.Close()
		//使用io。readall进行读取
		b, _ := io.ReadAll(r.Body)
		// fmt.Printf("b: %v\n", string(b))
		fmt.Println("-----------------------------------------------------------------------------------")
		filename := name + "吧" + "第" + strconv.Itoa(i) + "页" + ".html" //页数从int类型转string，可以使用strconv。itoa来进行
		f, err2 := os.Create(filename)

		if err2 != nil {
			fmt.Printf("create file failed ,err2: %v\n", err2)

		}
		f.Write(b)
		fmt.Println(website)
		defer f.Close()
	}
}
