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
		Name:  "wang",
		Age:   21,
		Email: &email, //指针类型可以传null值，别的不传默认为空值
	}
	s2 := Student{ //id自动创建
		Name:  "wang2",
		Age:   22,
		Email: nil, //指针类型可以传null值，别的不传默认为空值
	}

	s := []*Student{ //通过切片形式来生成多条记录
		&s1, &s2,
	}
	//s := []Student{
	//	s1, s2,
	//}
	err := DB.Create(s).Error
	if err != nil {
		println(err)
	}
}
