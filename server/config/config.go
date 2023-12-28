package config

import (
	"os"
	"strconv"
	"time"
)

var LogLevel = GetEnvOrDefault("LOGLEVEL", "error")
var FlushBytes = int64(GetEnvAsIntOrDefault("FLUSH_BYTES", "1000"))
var PrometheusEnabled = GetEnvAsBoolOrDefault("PROMETHEUS_ENABLED", "false")
var PrometheusNamespace = GetEnvOrDefault("PROMETHEUS_NAMESPACE", "corndogs")
var PrometheusQueueSizeEnabled = GetEnvAsBoolOrDefault("PROMETHEUS_QUEUE_SIZE_ENABLED", "true")
var PrometheusQueueSizeInterval = GetEnvAsDurationOrDefault("PROMETHEUS_QUEUE_SIZE_INTERVAL", "15s")
var PrometheusMetricQueryTimeout = GetEnvAsDurationOrDefault("PROMETHEUS_METRIC_QUERY_TIMEOUT", "5s")
var DefaultQueue = GetEnvOrDefault("DEFAULT_QUEUE", "default")
var DefaultStartingState = GetEnvOrDefault("DEFAULT_STARTING_STATE", "submitted")
var DefaultTimeout = int64(GetEnvAsIntOrDefault("DEFAULT_TIMEOUT", "0"))
var DefaultWorkingSuffix = "-working"
var RequestIdHeader = GetEnvOrDefault("REQUEST_ID_HEADER", "X-Request-ID")

func GetEnvOrDefault(env, defaultValue string) string {
	value := os.Getenv(env)
	if value == "" {
		value = defaultValue
	}
	return value
}

func GetEnvAsIntOrDefault(env, defaultValue string) int {
	value := os.Getenv(env)
	if value == "" {
		value = defaultValue
	}

	intValue, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		panic(err)
	}
	return int(intValue)
}

func GetEnvAsBoolOrDefault(env, defaultValue string) bool {
	value := os.Getenv(env)
	if value == "" {
		value = defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}
	return boolValue
}

func GetEnvAsDurationOrDefault(env, defaultValue string) time.Duration {
	value := os.Getenv(env)
	if value == "" {
		value = defaultValue
	}

	durationValue, err := time.ParseDuration(value)
	if err != nil {
		panic(err)
	}
	return durationValue
}
