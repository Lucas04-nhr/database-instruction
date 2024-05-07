package main

//To run this in linux:
//1. vi go.mod
//2. go mod edit -module=example.com/mod
//3. go get github.com/qiniu/qmgo
//4. click Run menu: Run without debugging
//5. boot mongodb
//cd /opt/mongodbrhel7/bin
//./mongod -f ../etc/mongoDB.conf
//6. boot:  $ /usr/bin/mongodb-compass  => click： CONNECT
//7. create databse: mongotest,  collection: mongouser

import (
	"context"
	"fmt"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	ctx := context.Background()
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017", Database: "mongotest", Coll: "mongouser"})
	defer func() {
		if err = cli.Close(ctx); err != nil {
			panic(err)
		}
	}()
	type UserInfo struct {
		Name   string `bson:"name"`
		Age    uint16 `bson:"age"`
		Weight uint32 `bson:"weight"`
	}

	var userInfo = UserInfo{
		Name:   "xm",
		Age:    7,
		Weight: 40,
	}
	//cli.CreateOneIndex(context.Background(), options.IndexModel{Key: []string{"name"}})
	//cli.CreateIndexes(context.Background(), []options.IndexModel{{Key: []string{"id2", "id3"}}})

	// insert one document
	result, err := cli.InsertOne(ctx, userInfo)
	fmt.Println("插入一条数据")
	fmt.Println(result) // &{ObjectID("633aafb8059f29c39a396235")}

	// find one document
	one := UserInfo{}
	err = cli.Find(ctx, bson.M{"name": userInfo.Name}).One(&one)
	fmt.Println("根据name查询")
	fmt.Println(one) //{xm 7 40}
	// multiple insert
	var userInfos = []UserInfo{
		UserInfo{Name: "a1", Age: 6, Weight: 20},
		UserInfo{Name: "b2", Age: 6, Weight: 25},
		UserInfo{Name: "c3", Age: 6, Weight: 30},
		UserInfo{Name: "d4", Age: 6, Weight: 35},
		UserInfo{Name: "a1", Age: 7, Weight: 40},
		UserInfo{Name: "a1", Age: 8, Weight: 45},
	}
	result1, err := cli.Collection.InsertMany(ctx, userInfos)
	fmt.Println("插入多条数据")
	fmt.Println(result1)
	/*
		&{[ObjectID("633aafb8059f29c39a396236") ObjectID("633aafb8059f29c39a396237")
		ObjectID("633aafb8059f29c39a396238") ObjectID("633aafb8059f29c39a396239")
		ObjectID("633aafb8059f29c39a39623a") ObjectID("633aafb8059f29c39a39623b")]}
	*/

	// 删除数据
	err = cli.Remove(ctx, bson.M{"age": 7})

	// find all 、sort and limit
	batch := []UserInfo{}
	cli.Find(ctx, bson.M{"age": 6}).Sort("weight").Limit(7).All(&batch)
	fmt.Println("sort+limit查询")
	fmt.Println(batch) //[{a1 6 20} {b2 6 25} {c3 6 30} {d4 6 35}]
	count, err := cli.Find(ctx, bson.M{"age": 6}).Count()

	fmt.Printf("age=6 的个数为: %d\n", count) //age=6 的个数为: 4

}
