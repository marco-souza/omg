package config

import "os"

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type settings struct {
	Model string
}

func NewSettings() *settings {
	return &settings{
		Model: getEnv("MODEL", "llama3"),
	}
}
