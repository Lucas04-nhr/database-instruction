package dblogin

import (
	"fmt"
	"maindbgo/hustdbroot/dbmodels"

	//"maindbgo/hustdbroot/models"
	"net/http"
)

func BioDbloginHandler(writer http.ResponseWriter, request *http.Request) {
	mycurpath := request.URL.Path //the value is: "/dblogin"
	// fmt.Fprintf(writer, "Hello!")
	fmt.Println("同学你好，你现在即将登录到上机练习！", mycurpath)
	biodbusername := request.FormValue("userName1")
	biodbuserpassword := request.FormValue("password1")

	if biodbusername == "" || biodbuserpassword == "" {
		http.Error(writer, "Your Name? or Password?", http.StatusNotFound)
		return
	}

	//connect to MongoDB to verify login and user/passwd
	mycurUser, err := dbmodels.Mondbconn(writer, request) //wts add to conn mongodb

	//naeme := models.User.user
	if mycurUser.STD_NAME == "" || err != nil {
		http.Error(writer, "You are not the right man!.", http.StatusNotFound)
		return
	}
	if mycurUser.STD_PASSWORD == "" {
		http.Error(writer, "Your password is wrong!.", http.StatusNotFound)
		return
	}

	//这里可以实现网页重定向  wts  2023.11.15
	//http.HandleFunc("/biodbcontent", biodbpraccontent)
	fmt.Print(fmt.Sprintf("%v+", request))
	writer.Header().Set("Cache-Control", "must-revalidate, no-store")
	writer.Header().Set("Content-Type", " text/html;charset=UTF-8")
	writer.Header().Set("Location", "/biodbcontent") //跳转地址设置
	////writer写入状态码
	writer.WriteHeader(http.StatusFound) //关键: StatusFound = 302 // 不能用StatusNotFound
	//writer.WriteHeader(http.StatusOK) //这个也不行!!!!!

	//http.Redirect(writer, request, "http://www.example.com", http.StatusFound)
	//writer.WriteHeader(301) //关键在这里！
	//http.Redirect(writer, request, "www.hust.edu.cn", http.StatusFound)
	//writer.Header().Set("Content-Type", "text/html")
	//biodbpraccontent(writer, request) //here just write original html!!!!!!!!!!!!!!!!!!!!!

	// if mycurpath != "/hello" {
	// 	http.Error(writer, "404 not found.", http.StatusNotFound)
	// 	return
	// }

	// if request.Method != "GET" {
	// 	http.Error(writer, "Method is not supported.", http.StatusNotFound)
	// 	return
	// }
}

// just for test
func Mylog() {
	fmt.Println("You are now in the login Mylog()!!!!!!!!")
}

// 这个是当学生认证成功后来的地方
func Biodbpraccontent(writer http.ResponseWriter, request *http.Request) {
	indexhtm := `<html>
		<head>
	<title>华中科技大学本科教学课程-数据库技术及应用</title>
	</head>
	<body>
	<div><img src="/hustdbroot/images/logo.png"><h2>华中科技大学本科教学课程</h2>
	<div align="center"><h1>《数据库技术及应用》上机实操</h1></div> </div>

	<h2>课程预备知识：</h2>
	数据库上机练习内容：
	<!--   000000   -->
	<form action="Lesson0" method="get" onsubmit="return  checkForm(this);">
	<table>
		<tr>
			<td><input type="submit" value="Lesson0" name="submit2"> 数据库上机预备: Oracle 19.3C数据库及配套软件安装
		</tr>
		</table>
	</form>

	<form action="Lesson1" method="get" onsubmit="return  checkForm(this);">
	<table>
		<tr>
			<td><input type="submit" value="Lesson1" name="submit2"> 上机练习一：Oracle数据库基础: SQL语言及数据库编程
		</tr>
		</table>
	</form>

	<form action="Lesson2" method="get" onsubmit="return  checkForm(this);">
	<table>
		<tr>
			<td><input type="submit" value="Lesson2" name="submit3"> SQL基础语句
		</tr>
		</table>
	</form>
	<form action="Lesson3" method="get" onsubmit="return  checkForm(this);">
	<table>
		<tr>
			<td><input type="submit" value="Lesson3" name="submit3"> 表空间及表的完整性约束
		</tr>
		</table>
	</form>
	<form action="Lesson4" method="get" onsubmit="return  checkForm(this);">
	<table>
		<tr>
			<td><input type="submit" value="Lesson4" name="submit3"> Oracle查寻处理及PL/SQL语言查询优化
		</tr>
		</table>
	</form>
	<form action="Lesson5" method="get" onsubmit="return  checkForm(this);">
	<table>
		<tr>
			<td><input type="submit" value="Lesson5" name="submit3"> 数据库连接：嵌入式SQL语言及数据库编程
		</tr>
		</table>
	</form>
	<form action="Lesson6" method="get" onsubmit="return  checkForm(this);">
	<table>
		<tr>
			<td><input type="submit" value="Lesson6" name="submit3"> 小型数据库设计: 酒店管理数据库
		</tr>
		</table>
	</form>
	<form action="Lesson7" method="get" onsubmit="return  checkForm(this);">
	<table>
		<tr>
			<td><input type="submit" value="Lesson7" name="submit3"> 数据库备份及恢复
		</tr>
		</table>
	</form>
	<form action="Lesson8" method="get" onsubmit="return  checkForm(this);">
	<table>
		<tr>
			<td><input type="submit" value="Lesson8" name="submit3">  MONGODB数据库基础: NoSQL语言及数据库编程
		</tr>
		</table>
	</form>
	<form action="Lessonvector9" method="get" onsubmit="return  checkForm(this);">
	<table>
		<tr>
			<td><input type="submit" value="Lesson9" name="submit3">  向量数据库基础及应用
		</tr>
		</table>
	</form>
	<form action="Lessonmsyql10" method="get" onsubmit="return  checkForm(this);">
	<table>
		<tr>
			<td><input type="submit" value="Lesson10" name="submit3">  phpStudy/wamp/lamp MySQL数据库基础及应用
		</tr>
		</table>
	</form>
	</body>
	</html>`

	writer.WriteHeader(200)
	writer.Write([]byte(indexhtm))
}
