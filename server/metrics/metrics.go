package metrics

import (
	"net/http"

	"github.com/TnLCommunity/corndogs/server/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var TasksTotal prometheus.Counter
var CompletedTasksTotal prometheus.Counter
var CanceledTasksTotal prometheus.Counter
var TimedOutTasksTotal prometheus.Counter

func StartMetricsEndpoint() {
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":8080", nil)
}

func InitializeMetrics() {
	TasksTotal = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: config.PrometheusNamespace,
		Name:      "tasks_total",
		Help:      "The total tasks that have ever been in the system",
	})
	CompletedTasksTotal = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: config.PrometheusNamespace,
		Name:      "completed_tasks_total",
		Help:      "The total tasks that have been completed",
	})
	CanceledTasksTotal = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: config.PrometheusNamespace,
		Name:      "canceled_tasks_total",
		Help:      "The total tasks that have been canceled",
	})
	TimedOutTasksTotal = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: config.PrometheusNamespace,
		Name:      "timed_out_tasks_total",
		Help:      "The total tasks that have been timed out",
	})
}
