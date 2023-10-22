//缓冲区是具有读取和写入方法的可变大小的字节缓冲区。Buffer的零值是准备使用的空缓冲区。

package main

import (
	"bytes"
	"fmt"
)

func test1() {
	u := "nihao"
	b := bytes.NewBufferString(u)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b: %T\n", b)

	var bb bytes.Buffer
	fmt.Printf("bb: %v\n", bb)
	fmt.Printf("bb: %T\n", &bb)

	//将切片d写入
	d := []byte("nihao,wang")
	bb.Write(d)
	fmt.Printf("bb: %v\n", &bb)

	//readbyte
	for {
		b2, err := bb.ReadByte() //读取第一个byte，bb的第一个byte被拿掉，赋值给 a => a, _ := b.ReadByte()
		if err != nil {
			break
		}
		fmt.Printf("b2: %v\n", string(b2))
	}
	fmt.Printf("bb: %v\n", &bb) //读取完毕之后可以看到内容为空

	//writestring、writerune写入
	u2 := "yibanban"
	bb.WriteString(u2)
	fmt.Printf("bb: %s\n", &bb)

	r := []rune("你好")
	bb.WriteRune(r[1]) //0 你	1 好
	fmt.Printf("bb: %v\n", &bb)

	//read
	c := make([]byte, 4) //这里建立3个容量的切边，每次只读3个
	for {
		_, err := bb.Read(c) //每次读3个byte
		if err != nil {
			break
		}
		fmt.Printf("c: %v\n", string(c))
	}
	fmt.Printf("bb: %v\n", &bb)
	// -------------------------------------------------------------------
	//writestring、writerune写入
	u3 := "yibanban"
	bb.WriteString(u3)
	fmt.Printf("bb: %s\n", &bb)

	r1 := []rune("你好")
	bb.WriteRune(r1[1]) //0 你	1 好
	fmt.Printf("bb: %v\n", &bb)
	// -----------------------------------------------------------------

	//readrune	读取第一个rune，bb的第一个rune被拿掉，赋值给 r => r, _ := b.ReadRune()
	r2, _, _ := bb.ReadRune()
	fmt.Printf("r2: %v\n", string(r2)) //? 返回y

}

func main() {
	test1()
}
