package config

import (
	"log"
	"os"
	"strconv"
)

const (
	keyEnvironment = "ENV"
	keyDataSrcUrl  = "DATA_SOURCE_URL"
	keyAppPort     = "APPLICATION_PORT"
)

func GetEnv() string {
	return getEnvironmentVariable(keyEnvironment)
}

func GetDataSrcUrl() string {
	return getEnvironmentVariable(keyDataSrcUrl)
}

func GetAppPort() int {
	portStr := getEnvironmentVariable(keyAppPort)
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port %s is invalid", portStr)
	}

	return port
}

func getEnvironmentVariable(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("%s environment variable is missing", key)
	}

	return val
}
