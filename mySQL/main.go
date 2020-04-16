package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//数据库连接示例：Mysql

var db *sql.DB

func initDB() (err error) {
	//数据库连接信息：dataSourceName
	dsn := "root:123123@tcp(192.168.241.129:3306)/go_demo"
	db, err = sql.Open("mysql", dsn) //判断dsn格式是否正确
	if err != nil {
		return
	}
	err = db.Ping() //尝试连接数据库；验证账号密码
	if err != nil {
		return
	}
	//db.SetMaxOpenConns(10) //设置连接池中最大的连接数
	//db.SetMaxIdleConns(5)  //设置连接池中的最大闲置连接数
	return
}

type user struct {
	id   int
	name string
	age  int
}

//单行查询
func queryOne(id int) (u user, err error) {
	sqlStr := `select id,name,age from user where id=?;`
	err = db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age) // 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	if err != nil {
		return
	}
	return
}

//多行查询
func queryRows(n int) {
	sqlStr := `select id,name,age from user where id>?;`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("query  failed, err:%v\n", err)
		return
	}
	defer rows.Close() //释放连接
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("query scan failed, err:%v\n", err)
			return
		}
		fmt.Println("query:", u)
	}
	return
}

//插入
func insert() {
	sqlStr := `insert into user(name,age)values(?,?)`
	res, err := db.Exec(sqlStr, "铁蛋", "888")
	if err != nil {
		fmt.Println("exec failed;err:", err)
		return
	}
	id, err := res.LastInsertId() //新插入数据的id
	if err != nil {
		fmt.Println("ressult return failed;err:", err)
		return
	}
	fmt.Println(id)
}

//更新
func updata() {
	sqlStr := `update user set age=? where id=?`
	res, err := db.Exec(sqlStr, "9991", 2)
	if err != nil {
		fmt.Println("exec  failed;err:", err)
		return
	}
	n, err := res.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Println("result  return failed;err:", err)
		return
	}
	fmt.Println(n)
}

//删除
func delete(id int) {
	sqlStr := `delete from user where id=?`
	res, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("exec  failed;err:", err)
		return
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println("result return  failed;err:", err)
		return
	}
	fmt.Println(n)
}

//mysql的预处理;当需要重复执行SQL的方法时；
//可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
//避免sql注入
func prepareQuery(n int) {
	sqlStr := `select id,name,age from user where id>?`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(" mysql prepare failed;err:", err)
		return
	}
	rows, err := stmt.Query(n)
	if err != nil {
		fmt.Println(" query  failed;err:", err)
		return
	}
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println(" query scan  failed;err:", err)
			return
		}
		fmt.Println(u)
	}
}

//事务  demo
func transaction() {
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		fmt.Println(" transaction begin  failed;err:", err)
		return
	}
	sqlStr1 := `update user set age=age-2 where id=1`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback()
		fmt.Println(" sql1 exec  failed;err:", err)
		return
	}
	sqlStr2 := `update user set age=age-2 where id=2`
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		fmt.Println(" sql2 exec  failed;err:", err)
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("transaction commit  failed;err:", err)
		return
	}

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initDB failed err:%v\n", err)
		return
	}
	fmt.Println("connect success!")
	u, err := queryOne(2)
	if err != nil {
		fmt.Printf("queryRow failed, err:%v\n", err)
		return
	}
	fmt.Println("queryRow:", u)
	// queryRows(0)
	// insert()
	// updata()
	// delete(3)
	// prepareQuery(1)
	transaction()
}
