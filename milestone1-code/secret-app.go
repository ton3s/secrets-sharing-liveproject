package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type Secrets struct {
	mu   sync.Mutex
	file *os.File
	data map[string]string
}

func newSecrets() *Secrets {
	return &Secrets{
		data: make(map[string]string),
	}
}

var secrets = newSecrets()

func (s *Secrets) addSecret(secret string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Generate MD5 hash of the secret
	hash := fmt.Sprintf("%x", md5.Sum([]byte(secret)))

	// Add to map
	s.data[hash] = secret

	// TODO: Save map to file

	return hash
}

// Request & Response structs
type SecretRequest struct {
	Secret string `json:"plain_text"`
}

type SecretResponse struct {
	Id string `json:"id"`
}

// Server handlers
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	message := []byte("ok")
	if _, err := w.Write(message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func secretHandlePost(w http.ResponseWriter, r *http.Request) {

	// Decode the body of the request and populate the SecretRequest struct
	var req SecretRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// MD5 has the secret and store in a map
	res := secrets.addSecret(req.Secret)

	// Respond with the a id of the MD5 hash of the secret
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {

	// Data file
	dataPath := os.Getenv("DATA_FILE_PATH")
	if dataPath == "" {
		log.Fatal("Environment variable DATA_FILE_PATH not found!")
	}

	// TODO: Check if data file already exists
	// TODO: Read data from file into the map
	// Create data file if it doesn't exist
	f, err := os.Create(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	secrets.file = f

	// HTTP server
	mux := http.NewServeMux()

	// Handlers
	mux.HandleFunc("/", secretHandler)
	mux.HandleFunc("/healthcheck", healthCheckHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
