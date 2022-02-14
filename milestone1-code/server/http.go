package server

import (
	"log"
	"net/http"
	"os"
)

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
