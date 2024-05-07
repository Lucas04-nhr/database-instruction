package main

//two:  方式: 使用go get github.com/godror/godror
//方式一
//使用go get github.com/mattn/go-oci8
//need:
//$ cd /opt/app/oracle/product/12.2.0/dbhome_1/bin/
//$ ./lsnrctl start
//]$ ./sqlplus /nolog
//SQL> conn as sysdba
//请输入用户名:  sys
//输入口令:
//已连接到空闲例程。
//SQL> startup
//ORACLE 例程已经启动。
//SQL> conn as sysdba
//数据库装载完毕。
//数据库已经打开。
//SQL>

import (
	"database/sql"
	"fmt"
	"log"

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

	//search students
	//need: CREATE TABLE drugdbuser( STD_ID NUMBER NOT NULL, STD_NAME VARCHAR2(100) NOT NULL, STD_NUM VARCHAR2(50) NOT NULL,STD_MAJOR VARCHAR2(50) NOT NULL, STD_PASSWORD VARCHAR2(200), STD_RECORDS NUMBER NOT NULL);
	//ORA-00942: 表或视图不存在
	//ORA-00990: 权限缺失或无效
	//SQL> grant unlimited tablespace to C##DRUGDBUSER;    need SYS in ORACLE 12 !!!!
	stdrows, err := db.Query("select STD_ID from drugdbuser")
	if err != nil {
		log.Fatalln(err)
	}

	defer stdrows.Close()
	for stdrows.Next() {
		var stdu string
		stdrows.Scan(&stdu)
		fmt.Println(stdu)
	}

	if err = stdrows.Err(); err != nil {
		log.Fatalln(err)
	}

	//search hustdb table     course
	row2, err := db.Query("SELECT STD_ID FROM BIODBSTD")
	if err != nil {
		log.Fatalln(err)
	}

	defer row2.Close()
	for row2.Next() {
		var data string
		row2.Scan(&data)
		fmt.Println(data)
	}

	if err = row2.Err(); err != nil {
		log.Fatalln(err)
	}
}
