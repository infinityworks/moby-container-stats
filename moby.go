package main

import (
	"bufio"
	"context"
	"encoding/json"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

// ContainerStats is used to track the core JSON response from the stats API
type ContainerStats struct {
	NetIntefaces map[string]struct {
		RxBytes   int `json:"rx_bytes"`
		RxDropped int `json:"rx_dropped"`
		RxErrors  int `json:"rx_errors"`
		RxPackets int `json:"rx_packets"`
		TxBytes   int `json:"tx_bytes"`
		TxDropped int `json:"tx_dropped"`
		TxErrors  int `json:"tx_errors"`
		TxPackets int `json:"tx_packets"`
	} `json:"networks"`
	MemoryStats struct {
		Usage int `json:"usage"`
		Limit int `json:"limit"`
	} `json:"memory_stats"`
	CPUStats struct {
		CPUUsage struct {
			PercpuUsage       []int `json:"percpu_usage"`
			UsageInUsermode   int   `json:"usage_in_usermode"`
			TotalUsage        int   `json:"total_usage"`
			UsageInKernelmode int   `json:"usage_in_kernelmode"`
		} `json:"cpu_usage"`
		SystemCPUUsage int64 `json:"system_cpu_usage"`
	} `json:"cpu_stats"`
	PrecpuStats struct {
		CPUUsage struct {
			PercpuUsage       []int `json:"percpu_usage"`
			UsageInUsermode   int   `json:"usage_in_usermode"`
			TotalUsage        int   `json:"total_usage"`
			UsageInKernelmode int   `json:"usage_in_kernelmode"`
		} `json:"cpu_usage"`
		SystemCPUUsage int64 `json:"system_cpu_usage"`
	} `json:"precpu_stats"`
}

// ListContainers returns a list of containers on the local system
func (e *Exporter) getStats(ch chan<- prometheus.Metric) error {

	cli, err := client.NewEnvClient()
	if err != nil {
		return errors.Wrapf(err, "Error creating Docker client")
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: false})
	if err != nil {
		return errors.Wrap(err, "Error obtaining container listing")
	}

	for _, c := range containers {

		stats, err := cli.ContainerStats(context.Background(), c.ID, false)
		if err != nil {
			return errors.Wrap(err, "Error obtaining container stats")
		}

		s := bufio.NewScanner(stats.Body)
		for s.Scan() {
			var v *ContainerStats
			if err := json.Unmarshal(s.Bytes(), &v); err != nil {
				return errors.Wrapf(err, "Could not unmarshal the response from the docker engine for container %s", c.ID)
			}

			// Set CPU metrics
			ch <- prometheus.MustNewConstMetric(e.containerMetrics["cpuUsagePercent"], prometheus.GaugeValue, calcCPUPercent(v), c.Names[0][1:], c.ID)

			// Set Memory metrics
			ch <- prometheus.MustNewConstMetric(e.containerMetrics["memoryUsagePercent"], prometheus.GaugeValue, calcMemoryPercent(v), c.Names[0][1:], c.ID)
			ch <- prometheus.MustNewConstMetric(e.containerMetrics["memoryUsageBytes"], prometheus.GaugeValue, float64(v.MemoryStats.Usage), c.Names[0][1:], c.ID)
			ch <- prometheus.MustNewConstMetric(e.containerMetrics["memoryLimit"], prometheus.GaugeValue, float64(v.MemoryStats.Limit), c.Names[0][1:], c.ID)

			for key, net := range v.NetIntefaces {

				ch <- prometheus.MustNewConstMetric(e.containerMetrics["rxBytes"], prometheus.GaugeValue, float64(net.RxBytes), c.Names[0][1:], c.ID, key)
				ch <- prometheus.MustNewConstMetric(e.containerMetrics["rxDropped"], prometheus.GaugeValue, float64(net.RxDropped), c.Names[0][1:], c.ID, key)
				ch <- prometheus.MustNewConstMetric(e.containerMetrics["rxErrors"], prometheus.GaugeValue, float64(net.RxErrors), c.Names[0][1:], c.ID, key)
				ch <- prometheus.MustNewConstMetric(e.containerMetrics["rxPackets"], prometheus.GaugeValue, float64(net.RxPackets), c.Names[0][1:], c.ID, key)
				ch <- prometheus.MustNewConstMetric(e.containerMetrics["txBytes"], prometheus.GaugeValue, float64(net.TxBytes), c.Names[0][1:], c.ID, key)
				ch <- prometheus.MustNewConstMetric(e.containerMetrics["txDropped"], prometheus.GaugeValue, float64(net.TxDropped), c.Names[0][1:], c.ID, key)
				ch <- prometheus.MustNewConstMetric(e.containerMetrics["txErrors"], prometheus.GaugeValue, float64(net.TxErrors), c.Names[0][1:], c.ID, key)
				ch <- prometheus.MustNewConstMetric(e.containerMetrics["txPackets"], prometheus.GaugeValue, float64(net.TxPackets), c.Names[0][1:], c.ID, key)
			}
			if s.Err() != nil {
				return errors.Wrapf(err, "Error handling Stats.body from Docker engine")
			}
		}
	}

	return nil
}
