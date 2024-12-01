package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

type webConfig struct {
	crtPath, keyPath, webPort string
}

type transferValues struct{ V, P int }

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Fprintf(w, "smth went wrong 500")
	}
	t.Execute(w, nil)

}

func valueSubmit(w http.ResponseWriter, r *http.Request) {
	va, err := strconv.Atoi(r.FormValue("number"))
	pr, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		fmt.Fprintf(w, "ошибка в данных")
	}
	toFill := transferValues{V: ((85 - va) * 1500) / 100}
	toFill.P = pr * toFill.V
	//	fmt.Fprint(w, "литров дозаправить", toFill, "стоимость ", toFill*p)

	t, err := template.ParseFiles("template/calc.html")
	if err != nil {
		fmt.Fprintf(w, "smth went wrong 500")
	}

	t.Execute(w, toFill)

}

func main() {
	startWebParam := webConfig{
		crtPath: filepath.Join("crt", "localhost.crt"),
		keyPath: filepath.Join("crt", "localhost.key"),
		webPort: ":15443"}

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/submit", valueSubmit)

	log.Fatal(http.ListenAndServeTLS(startWebParam.webPort, startWebParam.crtPath, startWebParam.keyPath, nil))

}
