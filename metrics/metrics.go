package metrics

import "github.com/prometheus/client_golang/prometheus"

// Return - returns a map of metrics to be used by the exporter
func Return() map[string]*prometheus.Desc {

	containerMetrics := make(map[string]*prometheus.Desc)

	// CPU Stats
	containerMetrics["cpuUsagePercent"] = prometheus.NewDesc(
		prometheus.BuildFQName("container", "cpu", "usage_percent"),
		"CPU usage percent for the specified container",
		[]string{"container_name", "container_id"}, nil,
	)

	// Memory Stats
	containerMetrics["memoryUsagePercent"] = prometheus.NewDesc(
		prometheus.BuildFQName("container", "memory", "usage_percent"),
		"Current memory usage percent for the specified container",
		[]string{"container_name", "container_id"}, nil,
	)
	containerMetrics["memoryUsageBytes"] = prometheus.NewDesc(
		prometheus.BuildFQName("container", "memory", "usage_bytes"),
		"Current memory usage in bytes for the specified container",
		[]string{"container_name", "container_id"}, nil,
	)
	containerMetrics["memoryLimit"] = prometheus.NewDesc(
		prometheus.BuildFQName("container", "memory", "limit"),
		"Memory limit as configured for the specified container",
		[]string{"container_name", "container_id"}, nil,
	)

	// Network Stats
	containerMetrics["rxBytes"] = prometheus.NewDesc(
		prometheus.BuildFQName("container", "net_rx", "bytes"),
		"Network RX Bytes",
		[]string{"container_name", "container_id", "interface"}, nil,
	)
	containerMetrics["rxDropped"] = prometheus.NewDesc(
		prometheus.BuildFQName("container", "net_rx", "dropped"),
		"Network RX Dropped Packets",
		[]string{"container_name", "container_id", "interface"}, nil,
	)
	containerMetrics["rxErrors"] = prometheus.NewDesc(
		prometheus.BuildFQName("container", "net_rx", "errors"),
		"Network RX Packet Errors",
		[]string{"container_name", "container_id", "interface"}, nil,
	)
	containerMetrics["rxPackets"] = prometheus.NewDesc(
		prometheus.BuildFQName("container", "net_rx", "packets"),
		"Network RX Packets",
		[]string{"container_name", "container_id", "interface"}, nil,
	)
	containerMetrics["txBytes"] = prometheus.NewDesc(
		prometheus.BuildFQName("container", "net_tx", "bytes"),
		"Network TX Bytes",
		[]string{"container_name", "container_id", "interface"}, nil,
	)
	containerMetrics["txDropped"] = prometheus.NewDesc(
		prometheus.BuildFQName("container", "net_tx", "dropped"),
		"Network TX Dropped Packets",
		[]string{"container_name", "container_id", "interface"}, nil,
	)
	containerMetrics["txErrors"] = prometheus.NewDesc(
		prometheus.BuildFQName("container", "net_tx", "errors"),
		"Network TX Packet Errors",
		[]string{"container_name", "container_id", "interface"}, nil,
	)
	containerMetrics["txPackets"] = prometheus.NewDesc(
		prometheus.BuildFQName("container", "net_tx", "packets"),
		"Network TX Packets",
		[]string{"container_name", "container_id", "interface"}, nil,
	)

	return containerMetrics
}

//
