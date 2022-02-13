package main

import (
	"log"
	"net/http"
	"os"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	message := []byte("ok")
	if _, err := w.Write(message); err != nil {
		log.Fatal(err)
	}
}

func secretHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		secretHandleGet(w, r)
	case "POST":
		secretHandlePost(w, r)
	}
}

func secretHandleGet(w http.ResponseWriter, r *http.Request) {
	message := []byte("GET/secret handler")
	if _, err := w.Write(message); err != nil {
		log.Fatal(err)
	}
}

func secretHandlePost(w http.ResponseWriter, r *http.Request) {
	message := []byte("POST/secret handler")
	if _, err := w.Write(message); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Data file
	dataPath := os.Getenv("DATA_FILE_PATH")
	if dataPath == "" {
		log.Fatal("Environment variable DATA_FILE_PATH not found!")
	}

	// Create data file
	f, err := os.Create(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// HTTP server
	mux := http.NewServeMux()

	// Handlers
	mux.HandleFunc("/", secretHandler)
	mux.HandleFunc("/healthcheck", healthCheckHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
