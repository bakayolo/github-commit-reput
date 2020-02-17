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

	GitUsername    string
	GitEmail       string
	GitCommitQueue int
	GitRepo        string
	GitDeployKey   []byte

	TwitterKeyword        string
	TwitterConsumerKey    string
	TwitterConsumerSecret string
	TwitterAccessToken    string
	TwitterAccessSecret   string

	RepoPath string
)

func LoadConfig() {
	var err error

	LogLevel = parseString("LOG_LEVEL")

	Timeout, err = strconv.Atoi(parseString("TIMEOUT"))
	if err != nil {
		log.Panic().Err(err).Msgf("TIMEOUT should be an integer")
	}

	GitUsername = parseString("GIT_USERNAME")
	GitEmail = parseString("GIT_EMAIL")
	GitCommitQueue, err = strconv.Atoi(parseString("GIT_COMMIT_QUEUE"))
	if err != nil {
		log.Panic().Err(err).Msgf("GIT_COMMIT_QUEUE should be an integer")
	}
	GitRepo = parseString("GIT_REPO")
	GitDeployKey, err = base64.StdEncoding.DecodeString(parseString("GIT_DEPLOY_KEY"))
	if err != nil {
		log.Panic().Err(err).Msgf("Error decoding GIT_DEPLOY_KEY")
	}

	TwitterKeyword = parseString("TWITTER_KEYWORD")
	TwitterConsumerKey = parseString("TWITTER_CONSUMER_KEY")
	TwitterConsumerSecret = parseString("TWITTER_CONSUMER_SECRET")
	TwitterAccessToken = parseString("TWITTER_ACCESS_TOKEN")
	TwitterAccessSecret = parseString("TWITTER_ACCESS_SECRET")

	RepoPath = parseString("REPO_PATH")
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
