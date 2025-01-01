package handlers

import (
	"encoding/json"
	"fmt"
	"gasweb/logic"
	"log"
	"net/http"
)

type incomingValues struct {
	Liters int `json:"liters" `
	Cost   int `json:"cost"`
}
type returnedMath struct {
	ToFill    int `json:"tofill"`
	TotalCost int `json:"totalcost"`
}

type requestMetaData struct {
	metod, url, headerFoo string
}

func ApiCalc(w http.ResponseWriter, r *http.Request) {

	incReq := requestMetaData{metod: r.Method, url: r.URL.Path}
	var v incomingValues
	d := json.NewDecoder(r.Body)
	err := d.Decode(&v)

	if isReqestvalid(incReq) && err != nil {

		zz, err := logic.DoMath(v.Liters, v.Cost)
		if err != nil {
			log.Println(err)
		}
		responsResult := returnedMath{
			ToFill:    zz.ToFullLiters,
			TotalCost: zz.ToFullCost,
		}
		setHeaders(w)
		json.NewEncoder(w).Encode(responsResult)
		return

	}

	http.Error(w, "not allowed", http.StatusForbidden)

}

func isJsonValid(d *json.Decoder) bool { //проверить что не нулевые и инты
	var v incomingValues
	err := d.Decode(&v)
	if err != nil {
		log.Println()
		return false

	}
	return true
}

func isReqestvalid(listOfHeaders requestMetaData) (bb bool) { //проверить урл вызова и метод
	fmt.Println(listOfHeaders.headerFoo)
	if listOfHeaders.url == "/api/v1/calc" && listOfHeaders.metod == "POST" {

		return true
	}
	return false
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("flag", "true")

}
