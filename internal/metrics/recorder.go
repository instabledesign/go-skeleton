package metrics

import "github.com/prometheus/client_golang/prometheus"

type Recorder struct {
	GrpcPoolGaugeVec *prometheus.GaugeVec
}

func (r *Recorder) mustRegister() *Recorder {
	prometheus.MustRegister(
		r.GrpcPoolGaugeVec,
	)
	return r
}

func (r *Recorder) CollectGrpcPoolState(pool *Pool) *Recorder {
	name := pool.GetName()
	r.GrpcPoolGaugeVec.WithLabelValues(name, "capacity").Set(float64(pool.Capacity()))
	r.GrpcPoolGaugeVec.WithLabelValues(name, "available").Set(float64(pool.Available()))
	return r
}

func NewRecorder(namespace string) *Recorder {
	return (&Recorder{
		GrpcPoolGaugeVec: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: "grpc_pool",
				Name:      "connection_count",
				Help:      "This represent the number of grpc connection",
			},
			[]string{"service", "type"},
		),
	}).mustRegister()
}
