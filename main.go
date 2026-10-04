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
		// added to make sure it only goes to root when root url is matched
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintln(w, "pulse is alive")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
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
