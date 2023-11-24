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

	var app application

	app.Domain = "example.com"

	log.Printf("http://localhost:%d でアプリケーションを起動", port)

	http.HandleFunc("/", app.Home)
	// http.HandleFunc("/sample", Sample)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
