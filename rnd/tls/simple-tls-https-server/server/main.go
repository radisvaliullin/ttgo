package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

const (
	port         = ":8443"
	responseBody = "Hello, TLS!"
)

func main() {
	fmt.Println("main")

	cert, err := tls.LoadX509KeyPair("./tmp/server.crt", "./tmp/server.key")
	if err != nil {
		log.Fatalf("failed to load X509 key pair: %v", err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	router := http.NewServeMux()
	router.HandleFunc("/", handleRequest)

	server := &http.Server{
		Addr:      port,
		Handler:   router,
		TLSConfig: config,
	}

	log.Printf("listening on %s...", port)
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseBody))
}
