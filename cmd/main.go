package main

import (
	"github-commit-reput/internal/config"
	"github-commit-reput/internal/twitter"
	"github-commit-reput/internal/utils"
	"github.com/rs/zerolog/log"
)

func main() {
	config.LoadConfig()
	log.Info().Msgf("Log Level is %s", utils.InitLogger())
	twitter.StartStreaming()
}
