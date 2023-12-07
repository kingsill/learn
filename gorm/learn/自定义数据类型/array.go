package main

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
)

var DB *gorm.DB

func init() {
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	Dbname := "gorm"
	timeout := "10s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("连接数据库失败, error=", err)
		return
	}
	DB = db
	fmt.Println("数据库连接成功")
}

type Ports []string

type HostModel struct {
	ID    int
	IP    string
	Ports Ports `gorm:"type:string"`
}

/*// Info json的序列化与反序列化的实例，定义Info的信息，方便后续进行转化及查询
type Info struct {
	Status     string `json:"status"`
	Addr       string `json:"addr"`
	Age        int    `json:"age"`
	LiveOrDead bool   `json:"liveOrDead"`
}

// User 定义User表，表中的Info字段想要传入的即为json类型的数据
type User struct {
	Name string
	Info Info `gorm:"type:string"` //这里由于我们已经实现了Scanner和Valuer接口，当不属于基本数据类型的数据传入时，会自动调用这两个接口，自动赋予类型。当然我们这里也可以提前指定好，我们这里选择string类型
}*/

func main() {
	G15 := HostModel{

		IP:    "192.168.1.1",
		Ports: []string{"1", "2", "3"},
	}

	DB.AutoMigrate(&HostModel{})
	DB.Create(&G15)

	var QueryUser HostModel
	DB.Take(&QueryUser)
	fmt.Printf("类型：%T\n内容：%v", QueryUser.Ports, QueryUser)
	//类型：main.Info
	//内容：{wang2 {ok zibo 18 true}}
}

//json方法----------------------------------------
//注意Scan方法传入为指针，而value直接传入结构体

/*// Scan 从数据库读取，将数据库中读取出来的数据类型还原为json,实现了sql.Scanner 接口
func (i *Ports) Scan(value interface{}) error {

	v, _ := value.([]byte) //类型断言，断定为[]byte类型，我们在value方法中也是转换为[]byte类型输入到数据库中的

	var receiver Ports
	err := json.Unmarshal(v, &receiver) //反序列化，将[]byte类型转化为我们需要的结构体
	if err != nil {
		return err
	}
	//fmt.Println(receiver)
	*i = receiver //将其内容传输给info

	return nil

}

// Value 存入数据库，将json转换为数据库可接受类型数据，实现dirver.Valuer接口
func (i Ports) Value() (driver.Value, error) {

	return json.Marshal(i) //由结构体转换为json类型数据，返回[]byte

}*/

//分割字符串方法------------------------------------

// Scan 从数据库读取，将数据库中读取出来的数据类型还原为json,实现了sql.Scanner 接口
func (i *Ports) Scan(value any) error {

	//v := value.(string) //!!!!错误版本，？？？？类型断言，断定为[]byte类型，我们在value方法中也是转换为[]byte类型输入到数据库中的
	//panic: interface conversion: interface {} is []uint8, not string

	v := value.([]byte) //我们在Scan函数中value的断言都选择[]byte方法

	cache := strings.Split(string(v), "|") //将v转变为字符串之后用|分隔符分割,还原为[]string类型数据

	*i = cache
	//fmt.Println(cache)
	return nil

}

// Value 存入数据库，将json转换为数据库可接受类型数据，实现dirver.Valuer接口
func (i Ports) Value() (driver.Value, error) {
	return strings.Join(i, "|"), nil //由结构体转换为json类型数据，返回string
}
