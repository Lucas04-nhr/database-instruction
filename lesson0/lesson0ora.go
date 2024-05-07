package lesmon0

import (
	"fmt"
	"net/http"
	"text/template"
)

// must be
func Less0prep(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("afawefaefasfasfasvsv")
	t, _ := template.ParseFiles("hustdbroot/hustbiodb0oracle19cbasic.html")
	t.Execute(writer, nil)
}
