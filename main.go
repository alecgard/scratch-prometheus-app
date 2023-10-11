package main

// Mock Prometheus client that increments Counter and Gauge metrics every second

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	v := 0
	go func() {
		fmt.Println("Starting inital 5s sleep")
		time.Sleep(5 * time.Second)
		fmt.Println("Finished inital sleep")

		for {
			opsProcessedCounter.Inc()
			opsProcessedGauge.Add(1)
			v++
			fmt.Printf("Metrics incremented: %d \n", v)
			time.Sleep(1 * time.Millisecond)
		}
	}()
}

var (
	opsProcessedCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "mock_counter",
		Help: "Mock counter",
	})
)

var (
	opsProcessedGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "mock_gauge",
		Help: "Mock gauge",
	})
)

func main() {
	// Only expose specified metrics
	r := prometheus.NewRegistry()
	r.MustRegister(opsProcessedCounter, opsProcessedGauge)
	handler := promhttp.HandlerFor(r, promhttp.HandlerOpts{})

	recordMetrics()

	fmt.Println("Metrics available on: http://localhost:2112/metrics")
	http.Handle("/metrics", handler)
	http.ListenAndServe(":2112", nil)
}
