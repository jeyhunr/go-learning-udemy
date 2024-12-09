package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	var app = application{Domain: "example.com"}

	app.Domain = "rahimli.net"

	log.Println("Listening on port", port, "with domain:", app.Domain)

	http.HandleFunc("/", Hello)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
