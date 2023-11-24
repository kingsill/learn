package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Student struct {
	ID    uint    `gorm:"size:10"`
	Name  string  `gorm:"size:16"`
	Age   int     `gorm:"size:3"`
	Email *string `gorm:"size:128"`
}

var DB *gorm.DB

func init() {
	username := "root"   //账号
	password := "123456" //密码
	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	port := 3306         //数据库端口
	Dbname := "gorm"     //数据库名
	timeout := "10s"     //连接超时，10秒

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	DB = db
	// 连接成功
	fmt.Println(db)
}

//var ss Student

func main() {
	//根据主键删除
	DB.Delete(&Student{}, 1)           //删除主键为1的记录
	DB.Delete(&Student{}, []int{1, 2}) //删除主键为1、2的记录

	//删除id为1的记录
	DB.Delete(&Student{ID: 1})

	//删除id为5的记录
	DB.Where("id=?", 5).Delete(&Student{})

	//使用其它条件进行删除，这里也是可以用于多条记录的删除
	email := "1256694651@qq.com"
	DB.Where("email=?", &email).Delete(&Student{})

	//使用其他条件进行删除
	DB.Where("email=?", "wang2001zilong@126.com").Delete(&Student{})
}
