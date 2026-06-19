package main

import (
	"fmt"
	"net/http"
	"time"
)


func runServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Port: " + addr + " URL: " + r.URL.String()))
	})

	server := &http.Server{
		Addr: addr,
		Handler: mux,
		WriteTimeout: 10 * time.Second,
		ReadTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on " + addr + " port...")
	server.ListenAndServe()
}

func main() {
	go runServer(":8080")
	runServer(":8081")
}