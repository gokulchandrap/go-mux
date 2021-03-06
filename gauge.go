package main

import (
	"time"
	"github.com/prometheus/client_golang/prometheus"
)

// ExampleGauge which will increment the incValue gauge by 1 every 5 seconds and promote it to /metrics
func ExampleGauge(prometheusRegistry *prometheus.Registry) {
	incValue := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "acme_company",
		Subsystem: "blob_storage",
		Name:      "inc_value",
		Help:      "just an increased number.",
	})
	// register incValue
	prometheusRegistry.MustRegister(incValue)
	// loop over the ticker and call Inc function
	for range time.Tick(time.Second * 5) {
		// increment incValue by 1 every 5 seconds
		go incValue.Inc()
	}
}
