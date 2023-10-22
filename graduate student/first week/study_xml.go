// xml与json相同
// xml包实现xml解析
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	//反引号
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}

func test1() {
	person := Person{
		Name:  "tom",
		Age:   20,
		Email: "wang@126.com",
	}

	// b, _ := xml.Marshal(person)
	b, _ := xml.MarshalIndent(person, "   ", "  ") //使每个标签从新的一行输出，并增加前缀相关信息
	fmt.Printf("%v\n", string(b))
}

func test2_read() {
	// b, _ := ioutil.ReadFile("a.xml")
	//ioutil包已经被弃用，被os和io包取代

	b, _ := os.ReadFile("a.xml")
	var p Person
	xml.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p)
}

func test3_write() {
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}

	//建立f作为writer
	f, _ := os.OpenFile("a.xml", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	defer f.Close()

	e := xml.NewEncoder(f)
	e.Encode(p)

}

func main() {
	test1()
	// test2_read()
	// test3_write()
}
