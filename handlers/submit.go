package handlers

import (
	"fmt"
	"gasweb/logic"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type resultValues struct{ V, P int }

func ValueSubmit(w http.ResponseWriter, r *http.Request) {

	currentValue := parseValues(r.FormValue("number"))
	currentPrice := parseValues(r.FormValue("price"))

	t, err := template.ParseFiles("template/calc.html")
	if err != nil {
		fmt.Fprintf(w, "page crushed")
	}
	toFill, err := logic.DoMath(currentValue, currentPrice)
	if err != nil {
		fmt.Fprintf(w, "ошибка обработки данных")

	}
	t.Execute(w, toFill)

}

func parseValues(s string) (i int) {
	convertedFormString, err := strconv.Atoi(s)
	if err != nil {
		log.Default()
	}
	return convertedFormString
}
