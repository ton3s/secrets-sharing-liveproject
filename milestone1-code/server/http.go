package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Request & Response structs
type SecretRequest struct {
	Secret string `json:"plain_text"`
}

type SecretResponse struct {
	Id string `json:"id"`
}

func NewHTTPServer(addr string) *http.Server {
	httpsrv := newHTTPServer()

	// Data file
	dataPath := os.Getenv("DATA_FILE_PATH")
	if dataPath == "" {
		log.Fatal("Environment variable DATA_FILE_PATH not found!")
	}
	httpsrv.Secrets.init(dataPath)

	// HTTP server
	mux := http.NewServeMux()

	// Handlers
	mux.HandleFunc("/", httpsrv.secretHandler)
	mux.HandleFunc("/healthcheck", httpsrv.healthCheckHandler)

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

type httpServer struct {
	Secrets *Secrets
}

func newHTTPServer() *httpServer {
	return &httpServer{
		Secrets: newSecrets(),
	}
}

func (s *httpServer) secretHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.secretHandleGet(w, r)
	case "POST":
		s.secretHandlePost(w, r)
	}
}

func (s *httpServer) secretHandleGet(w http.ResponseWriter, r *http.Request) {
	message := []byte("GET/secret handler")
	if _, err := w.Write(message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *httpServer) secretHandlePost(w http.ResponseWriter, r *http.Request) {

	// Decode the body of the request and populate the SecretRequest struct
	var req SecretRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// MD5 has the secret and store in a map
	res := s.Secrets.addSecret(req.Secret)

	// Respond with the a id of the MD5 hash of the secret
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Server handlers
func (s *httpServer) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	message := []byte("ok")
	if _, err := w.Write(message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
