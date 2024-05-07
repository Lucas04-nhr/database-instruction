package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-oci8"
)

var db *sql.DB

func initDB() (err error) {
	// dsn 用户名/密码@oracle机器IP:端口/服务名
	dsn := "c##drugdbuser/wtsgyh1972@localhost:1521/drugdb"
	db, err = sql.Open("oci8", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接(校验dsn是否正确)
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%s\n", err)
		return
	}
	defer db.Close()
	var user string
	err = db.QueryRow("select user from dual").Scan(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Successful  connection. Current user is: %v\n", user)
}
