package main

import (
	"log"
	"net/http"
	"os"
)

func healthCheckHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello from health check handler!")
	if _, err := writer.Write(message); err != nil {
		log.Fatal(err)
	}
}

func secretCheckHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello from secret handler!")
	if _, err := writer.Write(message); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Data file
	if dataPath := os.Getenv("DATA_FILE_PATH"); dataPath == "" {
		log.Fatal("Environment variable DATA_FILE_PATH not found!")
	}

	// HTTP server
	mux := http.NewServeMux()

	// Handlers
	mux.HandleFunc("/", secretCheckHandler)
	mux.HandleFunc("/healthcheck", healthCheckHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
