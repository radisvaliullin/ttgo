package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "proudly served with Go and HTTPS!")
	})

	srv := &http.Server{
		Addr:    *addr,
		Handler: mux,
	}

	log.Printf("starting server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
