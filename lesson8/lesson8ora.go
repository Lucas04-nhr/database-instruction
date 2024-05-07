package lesson8

import (
	"fmt"
	"net/http"
	"text/template"
)

// must be L case!!!!! for call from another package!!!!
func Lessonora8(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Oracle practise 8 for mongo db")
	t, _ := template.ParseFiles("hustdbroot/hustbiodb8mongodb.html")
	t.Execute(writer, nil)
	//http.Redirect(writer, request, "/hustdbroot/hustbiodb1oracle12.html", http.StatusMovedPermanently)
}

// enc := simplifiedchinese.GBK.NewEncoder()
// fmt.Println("7ghsdgdfagdfvgaergvdv")
// out, _ := enc.String("hustdbroot/hustbiodb1oracle12.html")
// t, _ := template.ParseFiles(out)
