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

	metrics, err := e.asyncRetrieveMetrics()

	for _, e := range err {
		if e != nil {
			fmt.Println("Error detected in metric retrieval : ", e)
		}
	}

	if len(metrics) == 0 {
		fmt.Println("No valid container metrics to process")
		return
	}

	for _, b := range metrics {
		e.setPrometheusMetrics(b, ch)
	}

	eLogger.Debug("Metric collection completed")

}

// setPrometheusMetrics takes the pointer to ContainerMetrics and uses the data to set the guages and counters
func (e *Exporter) setPrometheusMetrics(stats *ContainerMetrics, ch chan<- prometheus.Metric) {

	// Set CPU metrics
	ch <- prometheus.MustNewConstMetric(e.containerMetrics["cpuUsagePercent"], prometheus.GaugeValue, calcCPUPercent(stats), stats.ID, stats.Name)

	// Set Memory metrics
	ch <- prometheus.MustNewConstMetric(e.containerMetrics["memoryUsagePercent"], prometheus.GaugeValue, calcMemoryPercent(stats), stats.ID, stats.Name)
	ch <- prometheus.MustNewConstMetric(e.containerMetrics["memoryUsageBytes"], prometheus.GaugeValue, float64(stats.MemoryStats.Usage), stats.ID, stats.Name)
	ch <- prometheus.MustNewConstMetric(e.containerMetrics["memoryLimit"], prometheus.GaugeValue, float64(stats.MemoryStats.Limit), stats.ID, stats.Name)

	// Network interface stats (loop through the map of returned interfaces)
	for key, net := range stats.NetIntefaces {

		ch <- prometheus.MustNewConstMetric(e.containerMetrics["rxBytes"], prometheus.GaugeValue, float64(net.RxBytes), stats.ID, stats.Name, key)
		ch <- prometheus.MustNewConstMetric(e.containerMetrics["rxDropped"], prometheus.GaugeValue, float64(net.RxDropped), stats.ID, stats.Name, key)
		ch <- prometheus.MustNewConstMetric(e.containerMetrics["rxErrors"], prometheus.GaugeValue, float64(net.RxErrors), stats.ID, stats.Name, key)
		ch <- prometheus.MustNewConstMetric(e.containerMetrics["rxPackets"], prometheus.GaugeValue, float64(net.RxPackets), stats.ID, stats.Name, key)
		ch <- prometheus.MustNewConstMetric(e.containerMetrics["txBytes"], prometheus.GaugeValue, float64(net.TxBytes), stats.ID, stats.Name, key)
		ch <- prometheus.MustNewConstMetric(e.containerMetrics["txDropped"], prometheus.GaugeValue, float64(net.TxDropped), stats.ID, stats.Name, key)
		ch <- prometheus.MustNewConstMetric(e.containerMetrics["txErrors"], prometheus.GaugeValue, float64(net.TxErrors), stats.ID, stats.Name, key)
		ch <- prometheus.MustNewConstMetric(e.containerMetrics["txPackets"], prometheus.GaugeValue, float64(net.TxPackets), stats.ID, stats.Name, key)
	}

}
