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

func main() { //跟新的前提是查询到记录

	//save更新-----------------------------------

	var stu Student

	DB.Take(&stu, 3) //查询id为3的记录，将其转化为结构体实例
	//fmt.Println(stu)
	stu.Age = 18 //修改结构体内容
	stu.Name = "sb"

	DB.Save(&stu) //将结构体内容传回sql

	//save的create模式
	stu = Student{
		12, //原本没有id为12的记录
		"wang",
		26,
		nil,
	}
	DB.Save(&stu)

	//批量更新多条记录,单列
	var students []Student
	newEmail := "1256694541@qq.com"
	//这里查询如果使用其他条件可以参考前面的查询部分
	DB.Find(&students, []int{3, 4, 5}).Update("email", newEmail)
	students = []Student{} //清空students

	//使用model方法进行单列的数据更新
	DB.Model(&students).Where("name=?", "sb").Update("name", "大聪明")
	//DB.Model(&Student{}).Where("name=?", "sb").Update("name", "小聪明")//这种也可以

	students = []Student{} //清空students

	//使用model方法根据主键进行单列数据更新
	DB.Model(&Student{ID: 12}).Update("age", 100)

	//使用model方法进行多列数据更新
	DB.Model(&Student{}).Where("name=?", "wang3").Updates(Student{Name: "帅哥", Age: 18})

}
