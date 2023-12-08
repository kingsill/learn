package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type WeekDay int //定义一个数值的类型别名，方便显示

// 定义weekday与整数的对应
const (
	Sunday WeekDay = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

var WeekStringList = []string{"Sunday", "Monday", "Tuesday", "Thursday", "Wednesday", "Friday", "Saturday"}
var WeekTypeList = []WeekDay{Sunday, Monday, Tuesday, Thursday, Wednesday, Friday, Saturday}

// 通过定义weekday的string方法将其本质int映射为对应的string 自定义》字符
// 实现了自定义weekday类型的string接口，会在输出、打印weekday时自动调用，其被包含在stringer包中，具体可查阅该包
func (W WeekDay) String() string {
	s := WeekStringList[W]
	//fmt.Println("from string 方法", s)
	return s
}

func (W WeekDay) MarshalJSON() ([]byte, error) {
	s, _ := json.Marshal(W.String())
	//fmt.Println("marshal启用")
	return s, nil

}

// ParseStrWeekday 将字符串类型转化为自定义类型 字符串》自定义
func (W *WeekDay) ParseStrWeekday(week string) {
	for i, v := range WeekStringList {
		if v == week {
			*W = WeekTypeList[i]
			fmt.Println(*W)
		}
	}
}

// ParseIntWeekday Int》weekday自定义类型
func (W *WeekDay) ParseIntWeekday(week int) {
	*W = WeekTypeList[week-1]
}

type DayInfo struct {
	WeekDay WeekDay   `json:"WeekDay"`
	Data    time.Time `json:"Data"`
}

func main() {
	December7th := DayInfo{
		WeekDay: Wednesday,
		Data:    time.Now(),
	}
	re, _ := json.Marshal(December7th)

	fmt.Println(string(re))

	//通过ParseStrWeekday和ParseIntWeekday这两个方法，实现了输入字符串和int都可以转换为weekday类型，同时通过该类型的string接口的实现可以在输出时自动以对应的字符串了类型进行输出
	var w WeekDay
	w.ParseStrWeekday("Sunday")
	fmt.Println("我输入的字符串“Sunday”，输出结果为：", w)

	w.ParseIntWeekday(2)
	fmt.Println("我输入的int类型 2，输出结果为： ", w)

}
