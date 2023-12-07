package main

import (
	"encoding/json"
	"fmt"
)

type Status int // 类型别名

// 定义running、except、offline对应的常量，本质为整数数据
// 通过定义常量我们可以在查询时直接返回这里的常量名，而不是数字等容产生歧义
const ( //在线、异常、离线
	Running Status = iota + 1 // iota为golang中常量的行索引，起步为0，每行+1
	Except
	Offline
)

type Host struct {
	ID     uint   `gorm:"json:id"`
	Name   string `gorm:"json:name"`
	Status Status `gorm:"json:status"`
}

// 在status的json序列化过程中进行常量和字符串的转化
func (status Status) MarshalJSON() ([]byte, error) {
	var str string //定义一个字符串方便对status的json转换

	switch status {
	case Running:
		str = "running"
	case Except:
		str = "except"
	case Offline:
		str = "offline"
	}
	return json.Marshal(str)

}

func main() {
	//定义一个host主机实例
	var host = Host{
		Status: Running,
		Name:   "wang2",
	}

	result, _ := json.Marshal(host) //序列化host为json以存储

	fmt.Printf(string(result))
	//查询结果如下，可以看到status的值这里查询为字符
	//{"ID":0,"Name":"wang2","Status":"running"}
}
