package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	domain string
}

func main() {
	// set application config
	var app application

	// read from command line

	// connect to DB
	app.domain = "example.com"
	
	// start web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("error listening and serving to http: %s", err))
	}

	log.Println("Starting application on port", port)

}