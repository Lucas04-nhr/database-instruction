package lesson2

import (
	"fmt"
	"net/http"
	"text/template"
)

// must be L case!!!!! for call from another package!!!!
func Lessonora2(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("7ghsdgdfagdfvgaergvdv")
	t, _ := template.ParseFiles("hustdbroot/hustbiodb2oracle19c.html")
	t.Execute(writer, nil)
	//http.Redirect(writer, request, "/hustdbroot/hustbiodb1oracle12.html", http.StatusMovedPermanently)
}

// enc := simplifiedchinese.GBK.NewEncoder()
// fmt.Println("7ghsdgdfagdfvgaergvdv")
// out, _ := enc.String("hustdbroot/hustbiodb1oracle12.html")
// t, _ := template.ParseFiles(out)
