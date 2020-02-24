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
	"os"
	"os/signal"
	"syscall"
	"time"
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
	if err := git.InitRepo(path, globalConfig.GitRepo, globalConfig.GitDeployKey, globalConfig.GitCommitQueueMin, globalConfig.GitCommitQueueMax); err != nil {
		log.Panic().Msgf("Error initiating repo %v", path)
	}

	err := twitter.StartStreaming(
		globalConfig.TwitterConsumerKey,
		globalConfig.TwitterConsumerSecret,
		globalConfig.TwitterAccessToken,
		globalConfig.TwitterAccessSecret,
		globalConfig.TwitterKeyword)

	if err != nil {
		log.Panic().Msgf("Error streaming")
	}

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sign := <-ch:
		log.Info().Msgf("Killing signal received %v", sign)
	case <-time.After(time.Duration(globalConfig.Timeout) * time.Second):
		_ = git.CommitAndPushRepo(globalConfig.GitUsername, globalConfig.GitEmail) // push before timeout
		log.Info().Msgf("Timeout after %v seconds", time.Duration(globalConfig.Timeout)*time.Second)
	}

	twitter.StopStreaming()
}
