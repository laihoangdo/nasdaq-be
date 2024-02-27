package metric

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	skipPaths = map[string]bool{
		"/api/v1/liveness":  true,
		"/api/v1/readiness": true,
		"/metrics":          true,
		"/favicon.ico":      true,
	}
)

// App Metrics interface
type Metrics interface {
	IncHits(status int, method, path string)
	ObserveResponseTime(status int, method, path string, observeTime float64)
	IsSkipPath(path string) bool
	SetSkipPath(paths []string)
	UnSetSkipPath(paths []string)
	RunServer(port string)
}

// Prometheus Metrics struct
type PrometheusMetrics struct {
	ServiceName         string
	RequestTotal        prometheus.Counter
	RequestsTotalByPath *prometheus.CounterVec
	RequestDuration     *prometheus.HistogramVec
}

// Create metrics with address and name
func NewMetrics(serviceName string) (Metrics, error) {
	metr := PrometheusMetrics{
		ServiceName: serviceName,
		RequestTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "lof_http_request_total",
		}),
		RequestsTotalByPath: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "lof_http_requests_total",
			},
			[]string{"service_name", "status", "method", "path"},
		),
		RequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "lof_http_request_duration_seconds",
			},
			[]string{"service_name", "status", "method", "path"},
		),
	}

	if err := prometheus.Register(metr.RequestTotal); err != nil {
		return nil, err
	}

	if err := prometheus.Register(metr.RequestsTotalByPath); err != nil {
		return nil, err
	}

	if err := prometheus.Register(metr.RequestDuration); err != nil {
		return nil, err
	}

	if err := prometheus.Register(collectors.NewBuildInfoCollector()); err != nil {
		return nil, err
	}

	return &metr, nil
}

func (metr *PrometheusMetrics) RunServer(port string) {

	router := gin.New()
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	log.Printf("Metrics server is running on port: %s", port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting Server: %v", err)
	}
}

func (metr *PrometheusMetrics) IncHits(status int, method, path string) {
	metr.RequestTotal.Inc()
	metr.RequestsTotalByPath.WithLabelValues(metr.ServiceName, strconv.Itoa(status), method, path).Inc()
}

func (metr *PrometheusMetrics) ObserveResponseTime(status int, method, path string, observeTime float64) {
	metr.RequestDuration.WithLabelValues(metr.ServiceName, strconv.Itoa(status), method, path).Observe(observeTime)
}

func (metr *PrometheusMetrics) IsSkipPath(path string) bool {
	return skipPaths[path]
}

func (metr *PrometheusMetrics) SetSkipPath(paths []string) {
	for _, val := range paths {
		skipPaths[val] = true
	}
}

func (metr *PrometheusMetrics) UnSetSkipPath(paths []string) {
	for _, val := range paths {
		delete(skipPaths, val)
	}
}
