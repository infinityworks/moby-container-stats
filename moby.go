package main

import (
	"bufio"
	"context"
	"encoding/json"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

// ContainerMetrics is used to track the core JSON response from the stats API
type ContainerMetrics struct {
	ID           string
	Name         string
	Error        error
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

func (e *Exporter) asyncRetrieveMetrics() ([]*ContainerMetrics, error) {

	// Create new docker API client for passed down to the async requests
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, errors.Wrapf(err, "Error creating Docker client")
	}

	// Obtain a list of running containers only
	// Docker stats API won't return stats for containers not in the running state
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: false})
	if err != nil {
		return nil, errors.Wrap(err, "Error obtaining container listing")
	}

	// Channels used to enable concurrent requests
	ch := make(chan *ContainerMetrics, len(containers))
	ContainerMetrics := []*ContainerMetrics{}

	// range through the returned containers to obtain the statistics
	// Done due to there not yet being a '--all' option for the cli.ContainerMetrics function in the engine
	for _, c := range containers {

		go func(cli *client.Client, id, name string) {
			err := retrieveContainerMetrics(*cli, id, name, ch)
			if err != nil {
				errors.Wrapf(err, "Error obtaining stats")
			}

		}(cli, c.ID, c.Names[0][1:])

	}

	for {
		select {
		case r := <-ch:

			if r.Error != nil {
				errors.Wrapf(err, "Error processing stats")
				break
			}

			ContainerMetrics = append(ContainerMetrics, r)

			if len(ContainerMetrics) == len(containers) {
				return ContainerMetrics, nil
			}
		}

	}

}

func retrieveContainerMetrics(cli client.Client, id, name string, ch chan<- *ContainerMetrics) error {

	stats, err := cli.ContainerStats(context.Background(), id, false)
	if err != nil {
		errors.Wrapf(err, "Error obtaining container stats for %s, error: %v", id, err)
	}

	s := bufio.NewScanner(stats.Body)

	for s.Scan() {

		var c *ContainerMetrics

		if err := json.Unmarshal(s.Bytes(), &c); err != nil {
			errors.Wrapf(err, "Could not unmarshal the response from the docker engine for container %s", id)
		}
		c.ID = id
		c.Name = name

		ch <- c
	}

	if s.Err() != nil {
		return errors.Wrapf(err, "Error handling Stats.body from Docker engine")
	}

	return nil
}
