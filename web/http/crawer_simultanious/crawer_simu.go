package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

var baseUrl string = "http://tieba.baidu.com/f?kw=孙笑川&ie=utf-8&pn="

func main() {
	var start, end int
	fmt.Printf("请输入起始页（>= 1）:")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页（>= 起始页）:")
	fmt.Scan(&end)

	//
	SpiderPages(start, end)
}

// 根据起始页对网站进行爬取并保存内容
func SpiderPages(start, end int) {

	fmt.Printf("正在爬取 %d 到 %d 的页面", start, end)
	//通过管道确保所有协程运行完毕
	page := make(chan int)
	//
	for i := start; i <= end; i++ {
		go spider(i, page)
		// // 1）明确目标（要知道你准备在哪个范围或者网站去搜索）
		// url := baseUrl + strconv.Itoa((i-1)*50)
		// fmt.Println("url =", url)

		// // 2）爬（将所有网站的内容抓取下来
		// result, httpGetErr := HttpGet(url)
		// if httpGetErr != nil {
		// 	fmt.Println("HttpGetErr =", httpGetErr)
		// }

		// // 3）把内容写进文件
		// fileName := "孙笑川吧" + strconv.Itoa(i) + ".html"
		// f, fCreateErr := os.Create(fileName)
		// if fCreateErr != nil {
		// 	fmt.Println("fCreateErr =", fCreateErr)
		// 	continue // 这里不能用return，要不后面的也没法存了
		// }
		// f.WriteString(result)
		// f.Close()
	}
	//每当爬完一页之后，将读完的page传回，结束这里阻塞的通道
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}

// 向url发出请求相应，并读取消息体中内容
func HttpGet(url string) (result string, err error) {
	resp, httpGetErr := http.Get(url)
	if httpGetErr != nil {
		return "", httpGetErr
	}

	//
	defer resp.Body.Close()

	//
	r, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Printf("read failed,err2: %v\n", err2)
	}
	result = string(r)

	return result, nil
}

func spider(i int, page chan int) {
	// 1）明确目标（要知道你准备在哪个范围或者网站去搜索）
	url := baseUrl + strconv.Itoa((i-1)*50)
	fmt.Println("url =", url)

	// 2）爬（将所有网站的内容抓取下来
	result, httpGetErr := HttpGet(url)
	if httpGetErr != nil {
		fmt.Println("HttpGetErr =", httpGetErr)
	}

	// 3）把内容写进文件
	fileName := "孙笑川吧" + strconv.Itoa(i) + ".html"
	f, fCreateErr := os.Create(fileName)
	if fCreateErr != nil {
		fmt.Println("fCreateErr =", fCreateErr)
		return // 这里不能用return，要不后面的也没法存了
	}
	f.WriteString(result)
	defer f.Close()
	page <- i
}
