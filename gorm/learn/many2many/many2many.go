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

type Article struct {
	ID    uint   `gorm:"size:8"`
	Title string `gorm:"size:16"`
	Tags  []Tag  `gorm:"many2many:article_tags;"` //用于确定多对多的关系并指定第三张连接表的名字
}
type Tag struct {
	ID       uint      `gorm:"size:8"`
	Text     string    `gorm:"size:16"`
	Articles []Article `gorm:"many2many:article_tags;"` //反向引用，可以用来查询具有相同标签的文章
}

/*type User struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}*/

func main() {
	DB.AutoMigrate(&Article{}, &Tag{})
	//DB.AutoMigrate(&User{}, &Language{})

	//DB.Save(&Article{
	//	ID:    1,
	//	Title: "go",
	//	Tags: []Tag{
	//		{Text: "language"},
	//		{Text: "study"},
	//	},
	//})

	//var tag Tag
	//DB.Take(&tag, "Text=?", "study") //从已有标签中查询
	//DB.Save(&Article{
	//	Title: "Study notes",
	//	Tags:  []Tag{tag},
	//})

	//DB.Save(&User{
	//	Languages: []Language{
	//		{Name: "chinese"},
	//		{Name: "english"},
	//	},
	//})

	var article Article
	DB.Preload("Tags").Take(&article, 1)
	DB.Model(&article).Association("Tags").Delete(article.Tags)
	fmt.Println(article)

	var tag Tag
	DB.Preload("Articles").Take(&tag, 2)
	fmt.Println(tag)

	article = Article{}
	var tags []Tag

	DB.Find(&tags, []int{1, 2, 3}) //找到想要添加的标签

	DB.Preload("Tags").Take(&article, 1) //预加载要修改的文章

	DB.Model(&article).Association("Tags").Replace(tags) //替换文章标签

}
