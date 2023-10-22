//errors包实现了操作错误的函数。语言使用error类型来返回函数执行过程中遇到的错误，
//如果返回的error值为nil，则表示未遇到错误，
//否则error会返回一个字符串，用于说明遇到了什么错误。

// 自定义错误
package main

import (
	"fmt"
	"time"
)

// MyError is an error implementation that include a time and message.
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
