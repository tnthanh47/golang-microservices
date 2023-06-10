package main

import (
	"log"
	"net/http"
)

const port = ":80"

type Config struct{}

func main() {

	app := Config{}

	log.Printf("Starting broker service in port %s\n", port)

	srv := &http.Server{
		Addr:    port,
		Handler: app.routes(),
	}

	//start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
