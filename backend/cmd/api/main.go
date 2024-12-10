package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN    string
	Domain string
}

func main() {
	var app application

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 dbname=movies user=postgres password=postgres sslmode=disable timezone=UTC connect_timeout=5", "PostgreSQL connection string")
	flag.Parse()

	app.Domain = "rahimli.net"

	log.Println("Listening on port", port, "with domain:", app.Domain)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.route()); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
