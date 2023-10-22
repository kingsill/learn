package main

import "fmt"

func main() {
	//结构体——类型定义
	//type name struct{member definition;member2 definition}
	type person struct {
		id       int
		name     string
		email    string
		age      int
		language []string
	}
	type man struct {
		id, age     int
		name, email string
	}
	var zilong person
	zilong.id = 1
	zilong.name = "zilong"
	zilong.age = 18
	zilong.email = "wangzilong@.com"
	zilong.language = []string{"go", "c"}
	fmt.Printf("zilong: %T\n", zilong)
	fmt.Printf("zilong: %v\n", zilong)
	// uzi := man{
	// 	id:    01,
	// 	age:   18,
	// 	name:  "jianzihao",
	// 	email: "nul",
	// }
	uzi := man{
		01,
		18,
		"jianzihao",
		"nul",
	}
	fmt.Printf("uzi: %v\n", uzi)
}
