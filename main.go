package main

import (
	"gasweb/handlers"

	"log"
	"net/http"
)

type webConfig struct {
	crtPath, keyPath, webPort string
}

type transferValues struct{ V, P int }

func main() {
	startWebParam := webConfig{
		webPort: ":8081"}

	q := http.NewServeMux()

	q.HandleFunc("/", handlers.MainPage)
	q.HandleFunc("/submit", handlers.ValueSubmit)

	log.Fatal(http.ListenAndServe(startWebParam.webPort, q))

}
