package main

import (
	"fmt"
	"github-commit-reput/internal/config"
	globalConfig "github-commit-reput/internal/config"
	"github-commit-reput/internal/encrypt"
	"github-commit-reput/internal/file"
	"github-commit-reput/internal/git"
	"github-commit-reput/internal/utils"
	"github.com/rs/zerolog/log"
)

func main() {
	config.LoadConfig()
	log.Info().Msgf("Log Level is %s", utils.InitLogger())

	encrypt.GenerateKey()

	path := fmt.Sprintf("%v/%v", globalConfig.RepoPath, globalConfig.GitRepo)
	if err := file.InitFolder(path); err != nil {
		log.Panic().Msgf("Error initiating folder %v", path)
	}
	if err := git.InitRepo(path, globalConfig.GitRepo, globalConfig.GitUsername, globalConfig.GitDeployKey); err != nil {
		log.Panic().Msgf("Error initiating repo %v", path)
	}

	git.CommitAndPushRepo(globalConfig.GitUsername, globalConfig.GitEmail)

	// twitter.StartStreaming()
	// git.ProcessRepo(globalConfig.GitRepo, globalConfig.GitUsername)
}
