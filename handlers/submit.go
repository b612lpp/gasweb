package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type resultValues struct{ V, P int }

func ValueSubmit(w http.ResponseWriter, r *http.Request) {

	currentValue := parseValues(r.FormValue("number"))
	currentPrice := parseValues(r.FormValue("price"))
	toFill := doMath(currentPrice, currentValue)

	t, err := template.ParseFiles("template/calc.html")
	if err != nil {
		fmt.Fprintf(w, "page crushed")
	}

	t.Execute(w, toFill)

}

func doMath(a, b int) (r resultValues) {
	toUpperLine := ((85 - a) * 1500) / 100
	totalCost := toUpperLine * b
	result := resultValues{V: toUpperLine, P: totalCost}
	return result

}

func parseValues(s string) (i int) {
	convertedFormString, err := strconv.Atoi(s)
	if err != nil {
		log.Default()
	}
	return convertedFormString
}
