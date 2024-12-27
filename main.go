package main

import (
	"gasweb/handlers"

	"log"
	"net/http"
)

type webConfig struct {
	crtPath, keyPath, webPort string
}

func main() {
	startWebParam := webConfig{
		webPort: ":8081"}

	q := http.NewServeMux()

	q.HandleFunc("/", handlers.MainPage)
	q.HandleFunc("/submit", handlers.ValueSubmit)
	q.HandleFunc("/api/v1/calc", handlers.Handle)

	log.Fatal(http.ListenAndServe(startWebParam.webPort, q))

}
