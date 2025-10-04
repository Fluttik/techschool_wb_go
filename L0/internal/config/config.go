package config

import (
	"os"
)

type Config struct {
	PostgresURL string
	ServerPort  string
	KafkaTopic  string
}

func New() *Config {
	return &Config{
		PostgresURL: getEnv("POSTGRESQL_URL", ""),
		ServerPort:  getEnv("PORT", "8000"),
		KafkaTopic:  getEnv("TOPIC_NAME", "some_topic"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
