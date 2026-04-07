package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Root Endpoint with HTML Links
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<h1>Hello! This is a DevOps-Ready Go API</h1>
			<p>Deployed via GitHub Actions CI/CD pipeline.</p>
			<p><strong>Current Timestamp:</strong> %s</p>
			<hr>
			<h3>Explore DevOps Endpoints:</h3>
			<ul>
				<li><a href="/health">/health</a> - Check application status</li>
				<li><a href="/metrics">/metrics</a> - View Prometheus raw metrics</li>
			</ul>
		`, time.Now().Format(time.RFC1123))
	})

	// Health Check Endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Healthy"))
	})

	// Metrics Endpoint
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server is starting on port 8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
