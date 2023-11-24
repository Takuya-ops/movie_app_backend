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

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = conn

	app.Domain = "example.com"

	log.Printf("http://localhost:%d でアプリケーションを起動", port)

	http.HandleFunc("/", app.Home)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
