package handlers

import (
	"encoding/json"
	"fmt"
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

func Handle(w http.ResponseWriter, r *http.Request) {

	q := json.NewDecoder(r.Body)
	var v incomingValues
	err := q.Decode(&v)
	if err != nil {
		fmt.Println("empty json")
	}
	fmt.Println(v)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	//w.Write([]byte("wtf"))

	sameRespons := returnedMath{
		ToFill:    (85 - v.Liters) * 1500 / 100,
		TotalCost: ((85 - v.Liters) * 1500 / 100) * v.Cost,
	}
	json.NewEncoder(w).Encode(sameRespons)
	//fmt.Fprint(w, "qq")

}
