package main

import (
	"encoding/json"
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

func main() {
	//主键查询
	var student Student
	student.ID = 1
	err := DB.Take(&student).Error //获取查询是否失败
	fmt.Println(student, err)

	//其他条件查询
	var student2 Student
	affected := DB.Take(&student2, "age=?", 22).RowsAffected //获取查询的记录数
	fmt.Println(student2, affected)

	//查询多条记录
	var students []Student //首先创建切片

	rowsAffected := DB.Find(&students).RowsAffected //查询全部记录
	fmt.Println(students, rowsAffected)

	for _, s := range students {
		fmt.Println(s)

		marshal, _ := json.Marshal(s) //通过序列化可以读出指针所指的内容
		fmt.Println(string(marshal))
	}

	students = []Student{} //重新赋值

	//根据主键查询多条记录
	DB.Find(&students, []int{1, 2})
	//DB.Take(&students, []int{1, 3}) //take只能获取单条记录

	fmt.Println(students)

	students = []Student{} //重新赋值

	//根据其他条件查询多条记录，  可以与其他条件查询 单条记录进行对比
	DB.Find(&students, "age in ?", []int{21, 23}) //查询年龄为21和23岁的

	fmt.Println(students)

}
