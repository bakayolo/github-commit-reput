package config

import (
	"encoding/base64"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

var (
	LogLevel string

	Timeout int

	GitUsername       string
	GitEmail          string
	GitCommitQueueMin int
	GitCommitQueueMax int
	GitRepo           string
	GitDeployKey      []byte

	TwitterKeyword        string
	TwitterConsumerKey    string
	TwitterConsumerSecret string
	TwitterAccessToken    string
	TwitterAccessSecret   string

	RepoPath string
)

func LoadConfig() {
	var err error

	LogLevel = parseString("LOG_LEVEL", "DEBUG")

	Timeout = parseInt("TIMEOUT", "3600")

	GitUsername = parseString("GIT_USERNAME", "")
	GitEmail = parseString("GIT_EMAIL", "")
	GitCommitQueueMin = parseInt("GIT_COMMIT_QUEUE_MIN", "2")
	GitCommitQueueMax = parseInt("GIT_COMMIT_QUEUE_MAX", "5")
	GitRepo = parseString("GIT_REPO", "commit-reput-bot")
	GitDeployKey, err = base64.StdEncoding.DecodeString(parseString("GIT_DEPLOY_KEY", ""))
	if err != nil {
		log.Panic().Err(err).Msgf("Error decoding GIT_DEPLOY_KEY")
	}

	TwitterKeyword = parseString("TWITTER_KEYWORD", "COVID2019")
	TwitterConsumerKey = parseString("TWITTER_CONSUMER_KEY", "")
	TwitterConsumerSecret = parseString("TWITTER_CONSUMER_SECRET", "")
	TwitterAccessToken = parseString("TWITTER_ACCESS_TOKEN", "")
	TwitterAccessSecret = parseString("TWITTER_ACCESS_SECRET", "")

	RepoPath = parseString("REPO_PATH", "/tmp")
}

func parseString(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" && defaultValue == "" {
		log.Panic().Msgf("Could not find env var: %v", key)
	} else if value == "" && defaultValue != "" {
		value = defaultValue
	}
	log.Info().Msgf("Successfully loaded env var: %v=%v", key, value)
	return value
}

func parseInt(key, defaultValue string) int {
	intValue, err := strconv.Atoi(parseString(key, defaultValue))
	if err != nil {
		log.Panic().Err(err).Msgf("%v should be an integer", key)
	}

	return intValue
}
