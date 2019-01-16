package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	log.Printf("Received request from %s", r.RemoteAddr)
	fmt.Fprintf(w, "Hello from %s\n", hostname)
}

func main() {
	hostname, _ := os.Hostname()
	http.HandleFunc("/", handler)
	log.Printf("Starting server on %s", hostname)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("LISTEN_PORT"), nil))
}
