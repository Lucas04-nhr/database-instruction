package controllers

import (
	"fmt"
	"net/http"
)

func Dbcontrol(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is my home page")
}
