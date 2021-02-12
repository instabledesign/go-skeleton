package service

import (
	"sync"

	httpware_metrics "github.com/gol4ng/httpware/v4/metrics"
	httpware_prometheus "github.com/gol4ng/httpware/v4/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

var metricsRegistryOnce sync.Once

func (container *Container) GetMetricsRegistry() prometheus.Registerer {
	metricsRegistryOnce.Do(func() {
		container.metricsRegistry = prometheus.DefaultRegisterer
	})

	return container.metricsRegistry
}

var httpMetricsRecorderOnce sync.Once

func (container *Container) GetHTTPMetricsRecorder() httpware_metrics.Recorder {
	httpMetricsRecorderOnce.Do(func() {
		httpMetricsRecorder := httpware_prometheus.NewRecorder(httpware_prometheus.Config{})
		httpMetricsRecorder.RegisterOn(container.GetMetricsRegistry())
		container.httpMetricsRecorder = httpMetricsRecorder
	})

	return container.httpMetricsRecorder
}
