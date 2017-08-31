package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/structs"
	"github.com/infinityworks/go-common/logger"
	"github.com/infinityworks/moby-container-stats/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"

	im "github.com/infinityworks/go-common/metrics"
	cm "github.com/infinityworks/moby-container-stats/metrics"
)

var (
	applicationCfg   config.Config
	eLogger          *logrus.Logger
	containerMetrics map[string]*prometheus.Desc
)

// Exporter Sets up all the runtime and metrics
type Exporter struct {
	containerMetrics map[string]*prometheus.Desc
	config.Config
}

func init() {
	applicationCfg = config.Init()
	containerMetrics = cm.Return()
	eLogger = logger.Start(&applicationCfg)
}

func main() {

	eLogger.WithFields(structs.Map(applicationCfg)).Info("Starting Metric Exporter")

	// Register internal metrics used for tracking the exporter performance
	im.Init()

	exporter := Exporter{
		containerMetrics: containerMetrics,
		Config:           applicationCfg,
	}

	// This invokes the Collect method through the prometheus client libraries.
	prometheus.MustRegister(&exporter)

	// Setup HTTP handler
	port := fmt.Sprintf(":%s", applicationCfg.ListenPort())
	http.Handle(applicationCfg.MetricsPath(), prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		                <head><title>Moby Container Exporter</title></head>
		                <body>
		                   <h1>Moby Container Exporter</h1>
		                   <p><a href='` + applicationCfg.MetricsPath() + `'>Metrics</a></p>
		                   </body>
		                </html>
		              `))
	})
	eLogger.Fatal(http.ListenAndServe(port, nil))
}
