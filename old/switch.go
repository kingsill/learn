package main

import "fmt"

func main() {
	str := "f"
	switch str {
	case "a":
		str += "a"
	case "b":
		str += "b"
	default:
		str = "cc"
	}
	fmt.Println(str)
}
