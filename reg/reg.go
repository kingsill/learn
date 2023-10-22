package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a7c aac 888 a9c tac abbc a99c"
	buf2 := "3.14 567 agsdg 1.23 7. 8.99 1sdljgl 6.66"
	//``代表原生字符串，不丢失信息
	buf3 := `<!DOCTYPE html>
	<html>
	 <body > 
	  <div id="lowframe" style="position: fixed; bottom: 0; left: 0; height: 0; width: 100%; border-top: thin solid grey; background-color: white; overflow: auto;">
		... 
	  </div>
	  <!-- #lowframe --> 
		<h3 id="community">社区</h3> 
		<p> 这些服务可帮你寻找社区提供的开源包。 </p> 
		<p> 这些服务可帮你
		寻找社区提供
		的开源包。 </p> 
	  </div>
	  <!-- #page --> 
	 </body>
	</html>
	`

	//1.解释规则
	r := regexp.MustCompile(`a.c`)      //.代表任意一个字符
	r2 := regexp.MustCompile(`a[0-9]c`) //[0-9]代表0-9中任意一个字符

	r3 := regexp.MustCompile(`\d\.\d`)   //小数点左边一个数字，右边一个数字，以匹配小数
	r4 := regexp.MustCompile(`\d+\.\d+`) //+匹配前一个字符的1次或多次，这样写能够匹配全部小数

	r5 := regexp.MustCompile(`<p> .* </p> `) //*代表前面的字符匹配0次或多次，.*即表示中间匹配多个任意字符
	r6 := regexp.MustCompile(`<p>(?s:.*)</p>`)

	//2.根据规则提取信息
	b2 := r.FindAllString(buf, -1)
	fmt.Printf("规则1: %v\n", b2)
	b := r.Find([]byte(buf))
	fmt.Printf("规则1‘: %v\n", string(b))

	s := r2.FindAllString(buf, -1)
	fmt.Printf("规则2: %v\n", s)

	s2 := r3.FindAllString(buf2, -1)
	fmt.Printf("规则3，小数: %v\n", s2)
	s3 := r4.FindAllString(buf2, -1)
	fmt.Printf("股则4，完整小数: %v\n", s3)

	s4 := r5.FindAllString(buf3, -1)
	fmt.Printf("规则5，匹配xml中内容: %v\n", s4)
	s5 := r6.FindAllString(buf3, -1)
	fmt.Printf("规则6，换行符等也加入匹配s5: %v\n", s5)

}
