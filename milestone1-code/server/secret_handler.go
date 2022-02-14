package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ton3s/secrets-sharing-liveproject/milestone1-code/types"
)

func (s *httpServer) secretHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.secretHandleGet(w, r)
	case "POST":
		s.secretHandlePost(w, r)
	}
}

func (s *httpServer) secretHandleGet(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/")
	if len(id) == 0 {
		http.Error(w, "No Secret ID specified", http.StatusNotFound)
		return
	}

	// Check if the id exists
	secret, exists := s.Secrets.data[id]
	if !exists {
		http.Error(w, "Secret ID does not exist", http.StatusNotFound)
		return
	}

	// Send the plain text of the secret
	res := types.GetSecretResponse{Data: secret}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *httpServer) secretHandlePost(w http.ResponseWriter, r *http.Request) {
	// Decode the body of the request and populate the SecretRequest struct
	var req types.CreateSecretRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// MD5 has the secret and store in a map
	hash := s.Secrets.addSecret(req.Secret)

	// Respond with the a id of the MD5 hash of the secret
	res := types.CreateSecretResponse{Id: hash}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
