package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "0.1.0"

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, version)
}

func getPort() string {
	if p := os.Getenv("PORT"); p != "" {
		return p
	}
	return "8080"
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/version", versionHandler)

	addr := ":" + getPort()
	log.Printf("Pulse %s listening on %s", version, addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
