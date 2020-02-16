package main

import (
	"fmt"
	"github-commit-reput/internal/config"
	globalConfig "github-commit-reput/internal/config"
	"github-commit-reput/internal/encrypt"
	"github-commit-reput/internal/file"
	"github-commit-reput/internal/git"
	"github-commit-reput/internal/twitter"
	"github-commit-reput/internal/utils"
	"github.com/rs/zerolog/log"
)

// TODO I had to duplicate several lines of code in order to be faster. Deal with that later.

func main() {
	config.LoadConfig()
	log.Info().Msgf("Log Level is %s", utils.InitLogger())

	// Generate encryption key for encrypting the messages
	if err := encrypt.GenerateKey(); err != nil {
		log.Panic().Msgf("Error generating encryption key")
	}

	path := fmt.Sprintf("%v/%v", globalConfig.RepoPath, globalConfig.GitRepo)
	// create folder on the local
	if err := file.InitFolder(path); err != nil {
		log.Panic().Msgf("Error initiating folder %v", path)
	}
	// init repository in the folder
	if err := git.InitRepo(path, globalConfig.GitRepo, globalConfig.GitUsername, globalConfig.GitDeployKey); err != nil {
		log.Panic().Msgf("Error initiating repo %v", path)
	}

	// start twitter streaming
	if err :=
		twitter.StartStreaming(
			globalConfig.TwitterConsumerKey,
			globalConfig.TwitterConsumerSecret,
			globalConfig.TwitterAccessToken,
			globalConfig.TwitterAccessSecret,
			globalConfig.TwitterKeyword); err != nil {
		log.Panic().Msgf("Error streaming")
	}
}
