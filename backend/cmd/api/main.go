package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN    string
	Domain string
	DB     *sql.DB
}

func main() {
	var app application

	// read from cmd
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 dbname=movies user=postgres password=postgres sslmode=disable timezone=UTC connect_timeout=5", "PostgreSQL connection string")
	flag.Parse()

	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	app.DB = conn
	app.Domain = "rahimli.net"

	log.Println("Listening on port", port, "with domain:", app.Domain)

	if err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.route()); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
