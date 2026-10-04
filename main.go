package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"os"
)

var start = time.Now()

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pulse is alive")
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"status":         "ok",
			"uptime_seconds": int(time.Since(start).Seconds()),
		})
	})

	http.HandleFunc("/crash", func(w http.ResponseWriter, r *http.Request) {
		log.Println("crashing on purpose")
		os.Exit(1)
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
