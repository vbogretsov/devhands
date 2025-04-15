package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "path", "status"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
	/* RequestUserCPUMms = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_rusage_user_cpu_ms",
			Help: "User CPU consumed by request",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	) */
	/* RequestSystemCPUMms = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_rusage_system_cpu_ms",
			Help: "System CPU consumed by request",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	) */
	/* RequestsInOps = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_rusage_in_ops",
			Help: "Value of ru_inblock consumed during request",
		},
		[]string{"method", "path"},
	) */
	/* RequestsOutOps = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_rusage_out_ops",
			Help: "Value of ru_oublock consumed during request",
		},
		[]string{"method", "path"},
	) */
)

func init() {
	prometheus.MustRegister(RequestsTotal)
	prometheus.MustRegister(RequestDuration)
	// prometheus.MustRegister(RequestUserCPUMms)
	// prometheus.MustRegister(RequestSystemCPUMms)
	// prometheus.MustRegister(RequestsInOps)
	// prometheus.MustRegister(RequestsOutOps)
}
