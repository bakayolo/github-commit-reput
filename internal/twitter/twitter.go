package twitter

import (
	globalConfig "github-commit-reput/internal/config"
	"github-commit-reput/internal/encrypt"
	"github-commit-reput/internal/file"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func StartStreaming() {
	config := oauth1.NewConfig(globalConfig.TwitterConsumerKey, globalConfig.TwitterConsumerSecret)
	token := oauth1.NewToken(globalConfig.TwitterAccessToken, globalConfig.TwitterAccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	// using demux as documented in https://github.com/dghubble/go-twitter/blob/master/examples/streaming.go
	// even if not using it at the moment
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		processTweet(tweet)
	}

	filterParams := &twitter.StreamFilterParams{
		Track:         []string{globalConfig.TwitterKeyword},
		StallWarnings: twitter.Bool(true),
	}

	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Panic().Err(err)
	}

	go demux.HandleChan(stream.Messages)

	log.Info().Msgf("Starting stream on keyword %v", globalConfig.TwitterKeyword)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Info().Msgf("Stopping stream - received %v", <-ch)

	stream.Stop()
}

func processTweet(tweet *twitter.Tweet) {
	log.Debug().Msgf("Received tweet: %v", tweet)
	message, err := encrypt.Encrypt(tweet.Text)
	if err != nil {
		log.Panic().Err(err).Msg("Error encrypting the message")
	}
	if err := file.WriteRepo(message, tweet.IDStr); err != nil {
		log.Panic().Err(err)
	}
}
