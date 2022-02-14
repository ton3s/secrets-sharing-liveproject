package main

import (
	"log"

	"github.com/ton3s/secrets-sharing-liveproject/milestone1-code/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())

}
