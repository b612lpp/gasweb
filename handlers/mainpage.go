package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Fprintf(w, "smth went wrong 500")
	}
	t.Execute(w, nil)

}
