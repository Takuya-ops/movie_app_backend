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

	log.Printf("ポート番号：%dでアプリケーションを起動", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
