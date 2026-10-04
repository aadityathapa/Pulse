package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var start = time.Now()

var requests = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "pulse_http_requests_total",
		Help: "HTTP requests by route.",
	},
	[]string{"route"},
)

var _ = promauto.NewGaugeFunc(
	prometheus.GaugeOpts{
		Name: "pulse_uptime_seconds",
		Help: "Seconds since the process started.",
	},
	func() float64 { return time.Since(start).Seconds() },
)

func count(route string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requests.WithLabelValues(route).Inc()
		h(w, r)
	}
}

func main() {
	http.HandleFunc("/", count("root", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintln(w, "pulse is alive")
	}))

	http.HandleFunc("/health", count("health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"status":         "ok",
			"uptime_seconds": int(time.Since(start).Seconds()),
		})
	}))

	http.HandleFunc("/crash", func(w http.ResponseWriter, r *http.Request) {
		log.Println("crashing on purpose")
		os.Exit(1)
	})

	http.Handle("/metrics", promhttp.Handler())

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
