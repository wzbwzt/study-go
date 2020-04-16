package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//第三方库sqlx能够简化操作，提高开发效率。

var db *sqlx.DB

type user struct {
	ID   int
	Name string
	Age  int
}

func initDB() (err error) {
	//数据库连接信息：dataSourceName
	dsn := "root:123123@tcp(192.168.241.129:3306)/go_demo"
	db, err = sqlx.Connect("mysql", dsn) //判断dsn格式即账号密码是否正确
	if err != nil {
		return
	}
	//db.SetMaxOpenConns(10) //设置连接池中最大的连接数
	//db.SetMaxIdleConns(5)  //设置连接池中的最大闲置连接数
	return
}

func query(n int) {
	sqlStr := "select id, name, age from user where id > ?"
	var u []user
	err := db.Select(&u, sqlStr, n)
	if err != nil {
		fmt.Println("query failed err:", err)
		return
	}
	fmt.Printf("%#v", u)
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Println("connect db failed err:", err)
		return
	}
	query(0)
}
