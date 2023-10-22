// 这个包可以实现json的编码和解码，就是将json字符串转换为struct，或者将struct转换为json。
// marshal		 将struct编码成json，可以接受任意类型
// unmarshal		将json转码成struct结构体
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// struct	结构体本身以及其内部的字段都遵守大小写命名的暴露方式
type Person struct {
	Name   string
	Age    int
	Email  string
	Parent []string
}

// marshal 将结构体编码为json
func test1() {
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "wang2001zilong@126.com",
	}

	fmt.Printf("p: %v\n", p)

	b, _ := json.Marshal(p)
	fmt.Printf("json:b: %v\n", string(b))
}

// unmarshal json转换为结构体
func test2() {
	//建立json编码的数据
	b1 := []byte(`{"Name":"tom","Age":20,"Email":"tom@gmail.com"}`) //?

	var m Person

	json.Unmarshal(b1, &m)
	fmt.Printf("m: %v\n", m)

}

// 解析嵌套类型
func test3() {
	b := []byte(`{"Name":"tom","Age":20,"Email":"tom@gmail.com","Parents":["big tom","kite"]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	fmt.Printf("f: %v\n", f)
	fmt.Printf("f: %T\n", f)

	//对嵌套类型进行遍历输出
	/* 	for index, value := range iterable {	index为标签，value为内容
	   	} */
	for k, v := range f.(map[string]interface{}) { //?
		fmt.Printf("%v：%v\n", k, v)
	}
}

// 解析嵌套引用类型
func test4() {
	p := Person{
		Name:   "tom",
		Age:    20,
		Email:  "tom@gmail.com",
		Parent: []string{"big tom", "big kite"},
	}

	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b))
}

// encode 编码，将其他值转化为json值
func test5() {
	f, _ := os.OpenFile("a.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	defer f.Close()

	p := Person{
		Name:   "tom",
		Age:    20,
		Email:  "tom@gmail.com",
		Parent: []string{"big tom", "big kite"},
	}

	e := json.NewEncoder(f)
	err := e.Encode(p)
	fmt.Printf("err: %v\n", err)
}

// decode encode
func test6() {
	// a.json : {"Name":"tom","Age":20,"Email":"tom@gmail.com","Parents":["tom","kite"]}
	f, _ := os.Open("a.json")
	d := json.NewDecoder(f)
	e := json.NewEncoder(os.Stdout)

	var v interface{}
	// var v map[string]interface{}

	//decode解码，将json解码为结构体
	if err := d.Decode(&v); err != nil {
		log.Println(err)
		// return
	}
	fmt.Printf("v: %v\n", v)

	//encode编码，将结构体编码为json，由于stdout，可以直接输出
	if err := e.Encode(&v); err != nil {
		log.Println(err)
	}

}

func main() {
	test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()

}
