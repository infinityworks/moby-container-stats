package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

// Describe - loops through the API metrics and passes them to prometheus.Describe
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	for _, m := range e.containerMetrics {
		ch <- m
	}

}

// Collect function, called on by Prometheus Client library
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	eLogger.Debug("Metric collection requested")

	err := e.getStats(ch)
	if err != nil {
		fmt.Println("Error encountered", err)
	}

	eLogger.Debug("Metrics successfully collected")

}
