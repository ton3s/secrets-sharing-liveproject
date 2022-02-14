package server

import "net/http"

func (s *httpServer) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	message := []byte("ok")
	if _, err := w.Write(message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
