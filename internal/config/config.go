package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	LogLevel string

	GitUsername    string
	GitCommitQueue string
	GitDeployKey   string

	TwitterKeyword        string
	TwitterConsumerKey    string
	TwitterConsumerSecret string
	TwitterAccessToken    string
	TwitterAccessSecret   string

	RepoPath string
)

func LoadConfig() {
	LogLevel = parseString("LOG_LEVEL")

	GitUsername = parseString("GIT_USERNAME")
	GitCommitQueue = parseString("GIT_COMMIT_QUEUE")
	GitDeployKey = parseString("GIT_DEPLOY_KEY")

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
