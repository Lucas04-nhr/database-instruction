package dbmodels

import (
	"context"
	"fmt"
	"net/http"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type UserInfo struct {
	STD_ID       uint32 `bson:"stdid"`
	STD_NAME     string `bson:"stdname"`
	STD_NUM      string `bson:"stdnum"`
	STD_MAJOR    string `bson:"stdmajor"`
	STD_PASSWORD string `bson:"stdpasswd"`
	STD_RECORDS  uint16 `bson:"stdrecords"`
	ISVIP        string `bson:"isvip"`
}

// 创建一个学生实例
var userInfo = UserInfo{
	STD_ID:       2020001,
	STD_NAME:     "孙思宇",
	STD_NUM:      "U202011069",
	STD_MAJOR:    "生实202001",
	STD_PASSWORD: "U202011069",
	STD_RECORDS:  1000,
	ISVIP:        "NO", //here must has a  "," !!!!!!
}

func (u UserInfo) GetUserById(id int) {
	// 根据id获取用户信息
}

func (u UserInfo) CreatUser(user UserInfo) bool {
	//创建用户
	return true
}
func (u UserInfo) UpdateUser(user UserInfo) bool {
	return true
}

func (u UserInfo) DeleteUserById(id int) bool {
	return true
}

func Mondbconn(w http.ResponseWriter, r *http.Request) (curUser UserInfo, err error) {
	curUser.STD_NAME = r.FormValue("userName1")
	curUser.STD_PASSWORD = r.FormValue("password1")

	ctx := context.Background()
	//如果你的连接是指向固定的 database 和 collection，我们推荐使用下面的更方便的方法初始化连接，
	//后续操作都基于cli而不用再关心 database 和 collection
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017", Database: "biomondb", Coll: "biodbcol"})

	//初始化成功后，请defer来关闭连接
	defer func() {
		if err := cli.Close(ctx); err != nil {
			panic(err)
		}
	}()

	// 一定要初始化,否则报错,网上很多文章都没有做这一步
	//var findOptions = &options.FindOptions{}
	// find one document
	one := &UserInfo{}
	//err = cli.Find(ctx, bson.M{"stdnum": userInfo.STD_NUM}).One(&one)
	err = cli.Find(ctx, bson.M{"stdname": curUser.STD_NAME}).One(&one)
	fmt.Println("根据std_name查询")
	fmt.Println(one.STD_NAME) //{xm 7 40}

	//var many []UserInfo
	//err := cli.Find(ctx, bson.M{"stdid": bson.M{"$gte": 2020001, "$lte": 2020021}}).All(&many)
	//must use GET Method from login form, then here can get username1!!!!
	//mynextnm := r.FormValue("userName1")
	if one.STD_NAME != curUser.STD_NAME {
		fmt.Println("同学，你的上机信息未录入，请联系管理员或任课老师！", err)
		// insert one document ! wtssure inserted one line into collection
		//result, err := cli.InsertOne(ctx, userInfo)
		//fmt.Println("插入一条数据", err)
		//fmt.Println(result) // &{ObjectID("633aafb8059f29c39a396235")}
		curUser := &UserInfo{}
		http.Error(w, "Your Name??", http.StatusNotFound)
		return *curUser, err
	}
	if one.STD_PASSWORD != curUser.STD_PASSWORD {
		curUser := &UserInfo{}
		http.Error(w, "Your Password?", http.StatusNotFound)
		return *curUser, err
	}

	// multiple insert
	//var userInfos = []UserInfo{
	//{STD_ID: 20232001, STD_NAME: "Gyhhuststd", STD_NUM: "", STD_MAJOR: "", STD_PASSWORD: "", STD_RECORDS: 1003, ISVIP: ""},
	//{STD_ID: 20232002, STD_NAME: "Wzthuststd", STD_NUM: "", STD_MAJOR: "", STD_PASSWORD: "", STD_RECORDS: 1003, ISVIP: ""},
	//{STD_ID: 20232003, STD_NAME: "Weihuststd", STD_NUM: "", STD_MAJOR: "", STD_PASSWORD: "", STD_RECORDS: 1003, ISVIP: ""},
	//{STD_ID: 20232004, STD_NAME: "Xuphuststd", STD_NUM: "", STD_MAJOR: "", STD_PASSWORD: "", STD_RECORDS: 1003, ISVIP: ""},
	//{STD_ID: 20232005, STD_NAME: "Ulahuststd", STD_NUM: "", STD_MAJOR: "", STD_PASSWORD: "", STD_RECORDS: 1003, ISVIP: ""},
	//{STD_ID: 20232006, STD_NAME: "Kaohuststd", STD_NUM: "", STD_MAJOR: "", STD_PASSWORD: "", STD_RECORDS: 1003, ISVIP: ""},
	//{STD_ID: 20232007, STD_NAME: "Luphuststd", STD_NUM: "", STD_MAJOR: "", STD_PASSWORD: "", STD_RECORDS: 1003, ISVIP: ""},
	//}
	//result1, err := cli.Collection.InsertMany(ctx, userInfos)
	//fmt.Println("插入多条数据")
	//fmt.Println(result1)

	if err != nil {
		fmt.Println(err)
	}

	return *one, err
}
