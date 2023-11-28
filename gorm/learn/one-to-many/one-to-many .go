package main

import (
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

//表结构关联-----------------------
// User 用户表 一个用户可以有多篇文章

type User struct {
	ID       uint   `gorm:"size:4"`
	Name     string `gorm:"size:8"`
	Articles []Article
}

// Article 文章表 一篇文章属于一个用户
type Article struct {
	ID     uint   `gorm:"size:4"`
	Title  string `gorm:"size:16"`
	UserID uint   `gorm:"size:4"`
	User   User
}

//重写外键关联----------------------------
//gorm的foreignKey备注写在对应的两个表的关联上
//
//type User1 struct {
//	ID       uint       `gorm:"size:4"`
//	Name     string     `gorm:"size:8"`
//	Articles []Article1 `gorm:"foreignKey:UID"`
//}
//
//type Article1 struct {
//	ID    uint   `gorm:"size:4"`
//	Title string `gorm:"size:16"`
//	UID   uint   `gorm:"size:4"`
//	User  User1  `gorm:"foreignKey:UID"`
//}

//重写引用----------------------------
//备注写在对应的两个表的关联上

//type User2 struct {
//	ID       uint       `gorm:"size:4"`
//	Name     string     `gorm:"size:8"`
//	Articles []Article2 `gorm:"foreignKey:UserName;references:Name"`
//}
//
//type Article2 struct {
//	ID       uint   `gorm:"size:4"`
//	Title    string `gorm:"size:16"`
//	UserName string `gorm:"size:8"`
//	User     User2  `gorm:"references:Name"`
//}

func main() {
	//DB.AutoMigrate(&User{}, &Article{})
	//DB.AutoMigrate(&User1{}, &Article1{})
	//DB.AutoMigrate(&User2{}, &Article2{})

	////创建用户的同时创建文章，并将两者关联
	//DB.Save(&User{
	//	Name: "wang2",
	//	Articles: []Article{
	//		{Title: "golang"},
	//		{Title: "python"},
	//	},
	//})

	////创建文章，关联已有用户
	////1.直接传入关联外键部分
	//DB.Save(&Article{Title: "easyGo", UserID: 2})
	//
	////2.查询过后传入对应结构体中
	//var user User
	//DB.Take(&user, 1)                           //查询已有用户
	//DB.Save(&Article{Title: "c++", User: user}) //将关联部分的User结构体传入

	//外键添加
	//常规方法-----------------------将id为8的文章和id为2的用户绑定
	//现有用户，这里id=2
	//var user User
	//DB.Take(&user, 2)

	//现有文章，id=8，未设置用户
	//var article Article
	//DB.Take(&article, 8)

	//方法1.给现有用户绑定文章
	//user.Articles = []Article{article}//[{8 c语言 2 {0  []}}]
	//DB.Save(&user)

	//方法2.给现有文章关联用户
	//article.User = user //{2 wang2 []}
	//DB.Save(&article)
	//-------------------------------------
	//append方法--------将id为8的文章和id为2的用户绑定
	var user User
	DB.Take(&user, 2)
	var article Article
	DB.Take(&article, 8)
	//1.用户绑定文章
	//model在选表的同时也在选择对象
	//DB.Model(&User{ID: 2}).Association("Articles").Append(&article)
	DB.Model(&user).Association("Articles").Append(&article)
	//2.文章关联用户
	DB.Model(&article).Association("User").Append(&user)
}
