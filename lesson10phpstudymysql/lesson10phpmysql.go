package lesson9

import (
	"fmt"
	"net/http"
	"text/template"
)

// must be L case!!!!! for call from another package!!!!
func LessonPHPMYSql10(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("phpStudy and WAMP/LAMP practise 10 for MySQL db")
	t, _ := template.ParseFiles("hustdbroot/lesson10phpstudymysql/mysqlphpstudy.html")
	t.Execute(writer, nil)
	//http.Redirect(writer, request, "/hustdbroot/hustbiodb1oracle12.html", http.StatusMovedPermanently)
}

// enc := simplifiedchinese.GBK.NewEncoder()
// fmt.Println("7ghsdgdfagdfvgaergvdv")
// out, _ := enc.String("hustdbroot/hustbiodb1oracle12.html")
// t, _ := template.ParseFiles(out)
