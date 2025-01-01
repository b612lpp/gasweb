package main

import (
	"flag"
	"fmt"
	"gasweb/handlers"

	"log"
	"net/http"
)

type webConfig struct {
	crtPath, keyPath, webPort string
}

func main() {

	webPortFlag := flag.String("port", "8081", "server port")
	flag.Parse()
	startWebParam := webConfig{
		webPort: *webPortFlag,
	}

	q := http.NewServeMux()

	q.HandleFunc("/", handlers.MainPage)
	q.HandleFunc("/submit", handlers.ValueSubmit)
	q.HandleFunc("/api/v1/calc", handlers.ApiCalc)
	fmt.Println("server started on port " + startWebParam.webPort)
	log.Fatal(http.ListenAndServe(":"+startWebParam.webPort, q))

}
