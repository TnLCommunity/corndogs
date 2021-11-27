package config

import (
	"os"
	"strconv"
)

var LogLevel = GetEnvOrDefault("LOGLEVEL", "error")
var FlushBytes = int64(GetEnvAsIntOrDefault("FLUSH_BYTES", "1000"))
var PrometheusEnabled = GetEnvAsBoolOrDefault("PROMETHEUS_ENABLED", "false")
var DefaultQueue = GetEnvOrDefault("DEFAULT_QUEUE", "default")
var DefaultStartingState = GetEnvOrDefault("DEFAULT_STARTING_STATE", "submitted")

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
