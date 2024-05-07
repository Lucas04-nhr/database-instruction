package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/godror/godror"
)

func main() {
	//c##drugdbuser/wtsgyh1972@localhost:1521/drugdb
	db, err := sql.Open("godror", `user="c##drugdbuser" password="wtsgyh1972" connectString="localhost:1521/drugdb"`)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//查询ZHUJI表中STD_NUM='U202112156'的sql:
	//SELECT * FROM biodbstd where STD_NUM = 'U202112156'
	sqlStatement := "SELECT * FROM biodbstd where STD_ID = :2021002"
	stmt, err := db.Prepare(sqlStatement)
	defer stmt.Close()
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query("STD_NUM") //输入sql中对应参数的值
	if err != nil {
		panic(err)
	}

	defer rows.Close() //defer关闭查询连接
	//获取列相关信息
	strings, _ := rows.Columns()

	for i := 0; i < len(strings); i++ {
		fmt.Print(" ", strings[i])
	}
	fmt.Print("\n")
	//构造切片存储json
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})

	var STD_ID, STD_NUM, STD_NAME string
	for rows.Next() {
		rows.Scan(&STD_ID, &STD_NUM, &STD_NAME) //写入查询数据集的所有列名称
		fmt.Printf("code is %s, STD_NUM is %s\n", STD_ID, STD_NUM)
		m1["CODE"] = STD_ID
		m1["OVALUE1"] = STD_NUM
		m1["OVALUE2"] = STD_NAME
		slice = append(slice, m1) //分片中追加信息
	}
	if err = rows.Err(); err != nil {
		// handle the error here
	}
	defer stmt.Close()

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err = %v\n", err)
	}
	//输出序列化后的结果 json字符串
	fmt.Printf("序列化后 = %v\n", string(data))
}
