/*
该GO例程可实现对DMDB插入数据，修改数据，删除数据，数据查询等基本操作。需要以下几步：
1. 安装Go版的DM数据库驱动
Go 语言标准库 database/sql（https://golang.google.cn/pkg/database/sql/）提供了一系列数据库操作的标准接口，
DM 数据库基于 GO1.13 版本通过实现 database/sql 包的接口，向开发人员提供 DM 数据库操作的 Go 语言接口。
达梦安装后在安装目录的drivers下有go目录，该目录下面的dm-go-driver.zip
/opt/dm8/drivers/go/dm-go-driver.zip
将解压后的dm-go-driver.zip后的达梦go驱动放到GOROOT目录下的src/
//即拷到GOROOT: opt/myoptgo/go/src/dm !!!!!!

2. mkdir -p /opt/myoptgo/go/src/golang.org/x
手工下载** golang.org/x/text **
cd /opt/myoptgo/go/src/golang.org/x
$ git clone https://github.com/golang/text.git
select * from production.product;
3. 需要DMDB的示例数据库： select * from "PRODUCTION"."PRODUCT";
*/
package main

// 引入相关包
import (
	"database/sql"
	"dm"
	"fmt"
	"io/ioutil"
	"time"
)

var db *sql.DB
var err error

// drop schema "TEST" cascade;
func main() {
	driverName := "dm"
	//dataSourceName := "dm://sysdba:WTSgyh1972!@localhost:5236?socketTimeout=10&logLevel=info&schema=TEST" //dameng SQL> create user "test" identified by "WTSgyh1972!";
	dataSourceName := "dm://sysdba:WTSgyh1972!@localhost:5236?socketTimeout=10&logLevel=info" // no nned: create user is ok!!!!!!wts
	if db, err = connect(driverName, dataSourceName); err != nil {
		fmt.Println(err)
		return
	}
	//if err = insertTable(); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	if err = updateTable(); err != nil {
		fmt.Println(err)
		return
	}
	if err = queryTable(); err != nil {
		fmt.Println(err)
		return
	}
	//if err = deleteTable(); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	if err = disconnect(); err != nil {
		fmt.Println(err)
		return
	}
}

/* 创建数据库连接 */
func connect(driverName string, dataSourceName string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	if db, err = sql.Open(driverName, dataSourceName); err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Printf("connect to \"%s\" succeed.\n", dataSourceName)
	return db, nil
}

// insert
func insertTable() error {
	var inFileName = "./wts.jpg"
	//BOOKSHOP中表空间, production表
	//SELECT * FROM BOOKSHOP.PRODUCTION.PRODUCT;   //Python ODBC module操作DM数据库
	//java INSERT INTO PRODUCTION.BIG_DATA (photo, describe, txt)VALUES(?,?,?);
	//select * from "PRODUCTION"."PRODUCT";    //CENTOS 7.9 disql
	var sql = `INSERT INTO production.product(name,author,publisher,publishtime, product_subcategoryid,productno,satetystocklevel,originalprice,nowprice,discount,
description,photo,type,papertotal,wordtotal,sellstarttime,sellendtime)
VALUES(:1,:2,:3,:4,:5,:6,:7,:8,:9,:10,:11,:12,:13,:14,:15,:16,:17);`
	data, err := ioutil.ReadFile(inFileName)
	if err != nil {
		return err
	}
	t1, _ := time.Parse("2006-Jan-02", "2005-Apr-01")
	t2, _ := time.Parse("2006-Jan-02", "2006-Mar-20")
	t3, _ := time.Parse("2006-Jan-02", "1900-Jan-01")
	_, err = db.Exec(sql, "三国演义", "罗贯中", "中华书局", t1, 4, "9787101046121", 10, 19.0000, 15.2000,
		8.0, "《三国演义》是中国第一部长篇章回体小说，中国小说由短篇发展至长篇的原因与说书有关",
		data, "25", 943, 93000, t2, t3)
	if err != nil {
		return err
	}
	fmt.Println("insertTable succeed")
	return nil
}

// update
func updateTable() error {
	var sql = "UPDATE production.product SET name = :name WHERE productid = 11;"
	if _, err := db.Exec(sql, "三国演义（上）"); err != nil {
		return err
	}
	fmt.Println("updateTable succeed")
	return nil
}

// query
func queryTable() error {
	var productid int
	var name string
	var author string
	var description dm.DmClob
	var photo dm.DmBlob
	outFileName := "out.jpg"

	//var jsql = "SELECT photo FROM production.product WHERE productid=15"
	//jpgrows, err := db.Query(jsql)
	//var d = []byte(jpgrows)
	//ioutil.WriteFile(outFileName, d, 0666)
	var sql = "SELECT productid,name,author,description,photo FROM production.product WHERE	productid=11"
	rows, err := db.Query(sql)
	rc, _ := rows.Columns()

	if err != nil {
		return err
	}
	defer rows.Close()
	fmt.Println("queryTable results:")
	for rows.Next() {
		if err = rows.Scan(&productid, &name, &author, &description, &photo); err != nil {
			return err
		}
		blobLen, _ := photo.GetLength() //dm.DmBlob
		//var wjpg = [blobLen]byte
		wjpg := make([]byte, blobLen)
		// 参考DAMENG官方手册!!!
		////func (blob *DmBlob) ReadAt(pos int, dest []byte)  (int, error)
		//从指定偏移读取 DmBlob 对象存储的字节数组
		//pos：偏移，范围为 1~Blob 对象长度； dest：读取的字节存放到 dest 中。
		_, _ = photo.ReadAt(1, wjpg) //返回 int：实际读取字节数
		ioutil.WriteFile(outFileName, wjpg, 0666)
		fmt.Printf("%v %v %v %v %v\n", productid, name, author, description, blobLen)
		fmt.Printf("*******************************")
		fmt.Printf("%v\n", photo, rc)
	}
	return nil
}

//// delete
//func deleteTable() error {
//	var sql = "DELETE FROM production.product WHERE productid = 11;"
//	if _, err := db.Exec(sql); err != nil {
//		return err
//	}
//	fmt.Println("deleteTable succeed")
//	return nil
//}

/* 关闭数据库连接 */
func disconnect() error {
	if err := db.Close(); err != nil {
		fmt.Printf("db close failed: %s.\n", err)
		return err
	}
	fmt.Println("disconnect succeed")
	return nil
}
