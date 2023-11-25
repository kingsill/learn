package main

type Student struct {
	ID    uint    `gorm:"size:10"`
	Name  string  `gorm:"size:16"`
	Age   int     `gorm:"size:3"`
	Email *string `gorm:"size:128"`
}

func main() {
	//DB.AutoMigrate(&Student{})

	email := "wang2001zilong@126.com"
	//添加记录 结构实例化
	s1 := Student{ //id自动创建
		Name:  "wang3",
		Age:   24,
		Email: &email, //指针类型可以传null值，别的不传默认为空值
	}
	err := DB.Create(&s1).Error
	if err != nil {
		println(err)
	}
}
