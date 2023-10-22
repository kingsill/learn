package main

import "fmt"

func main() {
	for i, j := 1, 2; i < 100 && j < 1000; i++ {
		fmt.Println(i, j)
	}

}
