package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type WeekDay int //定义一个数值的类型别名，方便显示

var WeekStringList = []string{"Sunday", "Monday", "Tuesday", "Thursday", "Wednesday", "Friday", "Saturday"}
var WeekTypeList = []WeekDay{Sunday, Monday, Tuesday, Thursday, Wednesday, Friday, Saturday}

func (WeekDay WeekDay) String() string {

	s := WeekStringList[WeekDay-1]
	fmt.Println(s)
	return s
}

func (WeekDay WeekDay) MarshalJSON() ([]byte, error) {

	s, _ := json.Marshal(WeekDay.String())
	return s, nil
}

// 超级落后
const (
	Sunday WeekDay = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

/*
	func (WeekDay WeekDay) MarshalJSON() ([]byte, error) {
		var str string
		switch WeekDay {
		case Sunday:
			str = "Sunday"
		case Monday:
			str = "Monday"
		case Tuesday:
			str = "Tuesday"
		case Thursday:
			str = "Thursday"
		case Wednesday:
			str = "Wednesday"
		case Friday:
			str = "Friday"
		case Saturday:
			str = "Saturday"
		}
		return json.Marshal(str)
	}
*/

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

	fmt.Printf(string(re))
}
