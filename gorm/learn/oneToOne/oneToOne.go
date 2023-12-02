package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

type User struct {
	ID     uint    `gorm:"size:16"`
	Name   string  `gorm:"size:8"`
	IDCard *IDCard //通过idcard获得详细信息,进行后续查询，同时通过指针传递防止产生结构体嵌套
	//谁是指针不明确，怎么选择取决于具体的应用
}

type IDCard struct {
	ID     uint   `gorm:"size:16"`
	Age    int    `gorm:"size:4"`
	Addr   string `gorm:"size:16"`
	UserID uint   `gorm:"size:16"` //外键
	User   User   //关联的主表结构体
}

func main() {
	//DB.AutoMigrate(&User{}, &IDCard{})

	//同时添加两个表的信息
	//这里只能使用create，不能使用save （why？）
	/*	DB.Create(&User{
		ID:   1,
		Name: "wang2",
		IDCard: &IDCard{
			ID:   123,
			Age:  18,
			Addr: "shandong",
		},
	})*/

	//将新建的idcard绑定到已有user上
	//DB.Save(&User{Name: "cc"})
	//DB.Save(&IDCard{ID: 456, UserID: 2})

	var user User
	DB.Preload("IDCard").Take(&user, 2)
	fmt.Println(user)
	marshal, _ := json.Marshal(user)
	fmt.Println(string(marshal))

	//DB.Select("IDCard").Delete(&user)
	//
	//DB.Model(&user).Association("IDCard").Delete(&user.IDCard)

}
