package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
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

type Article struct {
	ID    uint
	Title string
	Tags  []Tag `gorm:"many2many:article2tag;"`
}

type Tag struct {
	ID   uint
	Text string
	//Articles []Article `gorm:"many2many:article2tag;"` //当使用反向引用时需要在setUpJoinTable时多设置一次这个表的 @@@
}

// Article2tag 自定义连接表
type Article2tag struct {
	ArticleID uint      `gorm:"primaryKey"`
	TagID     uint      `gorm:"primaryKey "` //上两项即为连接表默认项
	CreatedAt time.Time //自定义添加一个创建时间字段
}

func main() {
	//第一个参数为具有连接另一个表的字段的连接表，第二个即为连接字段的字段名，第三个为连接表
	DB.SetupJoinTable(&Article{}, "Tags", &Article2tag{})
	//DB.SetupJoinTable(&Tag{}, "Articles", &Article2tag{}) //与上面反向引用时对应 @@@
	//article2tag必须定义好所需的外键，否则会报错
	DB.AutoMigrate(&Article{}, &Tag{}, &Article2tag{})

}
