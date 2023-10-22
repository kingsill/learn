// 实现io.reader、io.readat、io.writeto、io.seeker、io.bytescanner
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	test1()
}

func test1() {
	u := "123456789"
	r := strings.NewReader(u)
	r2 := bytes.NewReader([]byte(u))

	//len	返回未读取部分的长度
	fmt.Printf("r.Len(): %v\n", r.Len())
	fmt.Printf("r2.Len(): %v\n", r2.Len())

	//返回数据总长度
	fmt.Printf("r.Size(): %v\n", r.Size())
	fmt.Printf("r2.Size(): %v\n", r2.Size())
	b := make([]byte, 2)
	for {
		_, err := r.Read(b)
		if err != nil {
			break
		}

		fmt.Println(string(b))
		//	fmt.Println(string(b(:n))) ?
		fmt.Printf("b: %T\n", b)
	}
	fmt.Println("————————————————————————————")

	//设置偏移量,因为上面的操作已经改变了读取位置等信息，这里重新定位到文件头 ？（不写也没事）
	r2.Seek(0, 0)
	for {
		//一个字节一个字节进行读取
		b2, err := r2.ReadByte()
		if err != nil {
			break
		}
		fmt.Printf("string(b2): %v\n", string(b2))

	}
	fmt.Println("————————————————————————————")

	off := int64(0)
	for {
		//指定偏移量读取
		n, err := r2.ReadAt(b, off) //从off个开始读，n返回读取的量
		if err != nil {
			break
		}
		off += int64(n) //计算下次开始读取的地址
		fmt.Println(off, string(b[:n]))
	}

}
