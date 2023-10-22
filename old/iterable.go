package main

import "fmt"

type Programmer struct {
	Name     string
	Age      int
	Job      string
	Language []string
}

func main() {
	programmer := Programmer{
		Name:     "jack",
		Age:      19,
		Job:      "coder",
		Language: []string{"Go", "C++"},
	}
	fmt.Println(programmer)
}
