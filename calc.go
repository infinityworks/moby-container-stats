package main

func calcCPUPercent(stats *ContainerMetrics) float64 {

	var CPUPercent float64

	cpuDelta := float64(stats.CPUStats.CPUUsage.TotalUsage - stats.PrecpuStats.CPUUsage.TotalUsage)
	sysDelta := float64(stats.CPUStats.SystemCPUUsage - stats.PrecpuStats.SystemCPUUsage)

	if sysDelta > 0.0 && cpuDelta > 0.0 {
		CPUPercent = (cpuDelta / sysDelta) * float64(len(stats.CPUStats.CPUUsage.PercpuUsage)) * 100.0
	}

	return CPUPercent
}

func calcMemoryPercent(stats *ContainerMetrics) float64 {
	return float64(stats.MemoryStats.Usage) * 100.0 / float64(stats.MemoryStats.Limit)
}
