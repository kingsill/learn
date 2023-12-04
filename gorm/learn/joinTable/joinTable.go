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

/*type Article struct {
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
	ArticleID uint `gorm:"primaryKey"`
	TagID     uint `gorm:"primaryKey "` //上两项即为连接表默认项

	CreatedAt time.Time //自定义添加一个创建时间字段
}*/

/*// 自定义主键部分-----------------------------------------
type Article struct {
	ID    uint
	Title string
	Tags  []Tag `gorm:"many2many:article2tag;joinForeignKey:A_ID;joinReferences:T_ID"`
}

type Tag struct {
	ID       uint
	Text     string
	Articles []Article `gorm:"many2many:article2tag;joinForeignKey:T_ID;joinReferences:A_ID"` //当使用反向引用时需要在setUpJoinTable时多设置一次这个表的 @@@
}

// Article2tag 自定义连接表
type Article2tag struct {
	A_ID      uint      `gorm:"primaryKey"`
	T_ID      uint      `gorm:"primaryKey "` //上两项即为连接表默认项
	CreatedAt time.Time //自定义添加一个创建时间字段
}
//------------------------------------------*/

// 查询连接表-------------------------------------------
type Article struct {
	ID    uint
	Title string
	Tags  []Tag `gorm:"many2many:article2tag;"`
}

type Tag struct {
	ID       uint
	Text     string
	Articles []Article `gorm:"many2many:article2tag;"` //当使用反向引用时需要在setUpJoinTable时多设置一次这个表的 @@@
}

// Article2tag 自定义连接表
type Article2tag struct {
	ArticleID uint    `gorm:"primaryKey"`
	TagID     uint    `gorm:"primaryKey "` //上两项即为连接表默认项
	Article   Article `gorm:"foreignKey:ArticleID"`
	Tag       Tag     `gorm:"foreignKey:TagID"`

	CreatedAt time.Time //自定义添加一个创建时间字段
}

//-------------------------------------------

func main() {
	//第一个参数为具有连接另一个表的字段的连接表，第二个即为连接字段的字段名，第三个为连接表
	DB.SetupJoinTable(&Article{}, "Tags", &Article2tag{})
	DB.SetupJoinTable(&Tag{}, "Articles", &Article2tag{}) //与上面反向引用时对应 @@@
	//article2tag必须定义好所需的外键，否则会报错
	DB.AutoMigrate(&Article{}, &Tag{}, &Article2tag{})

	/*	//添加文章并添加标签，并自动关联
		DB.Create(&Article{
			Title: "golang_study",
			Tags: []Tag{{Text: "go"},
				{Text: "study"}},
		})
	*/

	/*	//添加文章，关联已有标签
		var tag Tag
		DB.Take(&tag, "text=?", "study")
		DB.Create(&Article{
			Title: "how To Study",
			Tags:  []Tag{tag},
		})
	*/

	/*	//给已有文章关联标签
		DB.Save(&Article{Title: "gorm"}) //创建gorm文章
		//直接操作连接表，比较极端，知道文章和tag的ID，直接进行添加
		DB.Create(&Article2tag{
			ArticleID: 3,
			TagID:     2,
			CreatedAt: time.Time{},
		})
	*/
	/*	//先查询相关标签，在关联
		var tags []Tag //这里主要根据tag切片作为识别，切片名字可以随便起
		DB.Find(&tags, "text in?", []string{"go", "study"})

		var article Article
		DB.Take(&article, "title=?", "gin")

		DB.Model(&article).Association("Tags").Append(&tags)
	*/

	/*	//替换已有文章的标签
		var article Article
		DB.Preload("Tags").Take(&article, 3) //preload中参数需严格对应
		fmt.Println(article)
		var tag Tag
		DB.Take(&tag, "text=?", "go")
		DB.Model(&article).Association("Tags").Replace(&tag)
	*/

	/*	//查询文章列表并显示标签
		var articles []Article
		DB.Preload("Tags").Find(&articles)
		fmt.Println(articles)

		var article Article
		DB.Preload("Tags").Take(&article, 1)
		fmt.Println(article)

		DB.Preload("Tags").Take(&article, 1)
		var tags []Tag
		DB.Model(article).Limit(1).Association("Tags").Find(&tags) //这样可以分页，但是也无法查询连接表中其他字段，如关联时间
		fmt.Println(tags)*/

	var articleTag []Article2tag
	DB.Preload("Article").Preload("Tag").Where(map[string]any{"Article_id": 1}).Find(&articleTag) //大小写可忽略，符号不可忽略，具体需查看表内字段
	fmt.Println(articleTag)
	DB.Preload("Article").Preload("Tag").Where("article_id=?", 1).Find(&articleTag) //大小写可忽略，符号不可忽略，具体需查看表内字段
	fmt.Println(articleTag)
	DB.Preload("Article").Preload("Tag").Find(&articleTag, "article_id=?", 1) //大小写可忽略，符号不可忽略，具体需查看表内字段
	fmt.Println(articleTag)
}
