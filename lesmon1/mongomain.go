package lesmon1

import (
	"fmt"
	"net/http"
	"text/template"
)

// must be
func Lessmongo(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("afawefaefasfasfasvsv")
	t, _ := template.ParseFiles("hustdbroot/hustbiodb2mongodb.html")
	t.Execute(writer, nil)
}
