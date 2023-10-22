package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("D:/golang/a.log")
	buf := make([]byte, 1024)
	n, _ := f.Read(buf)
	fmt.Printf("buf: %v\n", string(buf[:n]))
}
