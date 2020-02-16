package commons

import (
	globalConfig "github-commit-reput/internal/config"
	"github-commit-reput/internal/encrypt"
	"github-commit-reput/internal/file"
	"github-commit-reput/internal/git"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/rs/zerolog/log"
)

// TODO See if it makes sense to have a commons. I am sure there is a better to way to do that.
// Purpose is to have a place to "everything" - the rest should be really independent

func ProcessTweet(tweet *twitter.Tweet) {
	log.Debug().Msgf("Received tweet: %v", tweet)
	message, err := encrypt.Encrypt(tweet.Text)
	if err != nil {
		log.Panic().Err(err).Msg("Error encrypting the tweet")
	}
	if err := file.WriteInFolder(message, tweet.IDStr); err != nil {
		log.Panic().Err(err).Msg("Error writing the tweet into the folder")
	}
	if err := git.CommitAndPushRepo(globalConfig.GitUsername, globalConfig.GitEmail); err != nil {
		log.Panic().Err(err).Msg("Error committing/pushing in repo")
	}
}
