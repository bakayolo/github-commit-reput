package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	LogLevel string
)

func LoadConfig() {
	LogLevel = parseString("LOG_LEVEL")
}

func parseString(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Panic().Msgf("Could not find env var: %v", key)
	} else {
		log.Info().Msgf("Successfully loaded env var: %v=%v", key, value)
	}
	return value
}
