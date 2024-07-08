package main

import (
	"net/http"
	"os"
)

func main() {
	if _, err := os.Stat("cert.pem"); os.IsNotExist(err) {
		generateCert()
	}

	server := http.Server{
		Addr:    "192.0.0.1:8080",
		Handler: nil,
	}

	server.ListenAndServeTLS("cert.pem", "key.pem")
}
