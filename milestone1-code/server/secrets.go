package server

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type Secrets struct {
	mu   sync.Mutex
	path string
	data map[string]string
}

func newSecrets() *Secrets {
	return &Secrets{
		data: make(map[string]string),
	}
}

func (s *Secrets) init(dataPath string) {
	var f *os.File

	// Check if data file already exists
	if _, err := os.Stat(dataPath); err != nil {
		// Create data file if it doesn't exist
		if f, err = os.Create(dataPath); err != nil {
			log.Fatal("Error creating new file", err)
		}
	} else {
		// Open existing file
		if f, err = os.Open(dataPath); err != nil {
			log.Fatal("Error opening existing file", err)
		}
		// Read the contents of file into map
		jsonData, err := io.ReadAll(f)
		if err != nil {
			log.Fatal("Error reading JSON data!", err)
		}
		if len(jsonData) != 0 {
			if err := json.Unmarshal(jsonData, &s.data); err != nil {
				log.Fatal("Error converting JSON to map!", err)
			}
		}
		fmt.Println(s.data)
	}
	defer f.Close()
	s.path = dataPath
}

func (s *Secrets) addSecret(secret string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Generate MD5 hash of the secret
	hash := fmt.Sprintf("%x", md5.Sum([]byte(secret)))

	// Add to map
	s.data[hash] = secret

	// Marshal the map into a JSON string
	data, err := json.Marshal(s.data)
	if err != nil {
		log.Fatal(err)
	}

	// Save map to file
	f, err := os.Create(s.path)
	if err != nil {
		log.Fatal("Error creating new file", err)
	}
	if _, err := f.Write(data); err != nil {
		log.Fatal(err)
	}
	return hash
}
