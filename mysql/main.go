package main

import (
	DBsql "DBsql/DB"
	"database/sql"
	"fmt"
)

// 定义 db 为相关方法
var db *sql.DB

// 定义 test表内内容 的结构体，方便后续查询
type test struct {
	ID   int
	name string
	age  int
}

func main() {

	//初始化连接
	db = DBsql.InitDB()
	//最后 断开连接
	defer db.Close()

	// insert_test("王二", 23)

	// QueryOneRow()
	// QueryOneRow_ID(2)

	QueryRows()

	UpdateData()

	DeleteData()
}

func insert_test(name string, age int) {
	//第一句为DML语句，其中‘？’为相关参数，在后面进行补充
	ret, err := db.Exec("insert into test (name,age) values (?,?)", name, age)

	if err != nil {
		fmt.Println("插入失败")
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("ok!")
	}
	//这里查询新插入数据的ID
	i, err2 := ret.LastInsertId()
	if err2 != nil {
		fmt.Println("get lastInsertID failed")
	}
	fmt.Printf("insert success, ID is: %v\n", i)
}

func QueryOneRow() {
	var t test
	//queryrow后调用scan方法进行读取，并且释放数据库连接
	err := db.QueryRow("select * from test where id = ?", 1).Scan(&t.ID, &t.name, &t.age)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("t: %v\n", t)
}

// 根据ID查数据
func QueryOneRow_ID(id int) {
	var t test
	//queryrow后调用scan方法进行读取，并且释放数据库连接
	err := db.QueryRow("select * from test where id = ?", id).Scan(&t.ID, &t.name, &t.age)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("t: %v\n", t)
}

// 查询多行数据
func QueryRows() {
	var t test
	r, err := db.Query("select * from test where age > ?", 20)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	//使用next循环读取数据
	for r.Next() {
		err2 := r.Scan(&t.ID, &t.name, &t.age)
		if err2 != nil {
			fmt.Printf("err2: %v\n", err2)
		}
		fmt.Printf("t: %v\n", t)
	}
}

// 更新数据
func UpdateData() {
	r, err := db.Exec("update test set name = ? where id =?", "王小波", 2)
	if err != nil {
		fmt.Println("数据更新失败，", err)
	}
	//获得影响的行数
	i, _ := r.RowsAffected()
	fmt.Println("更新数据成功，更新行数:", i)
}

func DeleteData() {
	r, err := db.Exec("delete from test where id = ?", 1)
	if err != nil {
		fmt.Println("删除成功")
	}
	i, _ := r.RowsAffected()
	fmt.Println("删除行数：", i)
}
