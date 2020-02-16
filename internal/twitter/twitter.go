package twitter

import (
	"github-commit-reput/internal/commons"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func StartStreaming(consumerKey, consumerTwitter, accessToken, accessSecret, keyword string) error {
	config := oauth1.NewConfig(consumerKey, consumerTwitter)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	// using demux as documented in https://github.com/dghubble/go-twitter/blob/master/examples/streaming.go
	// even if not using it at the moment
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		commons.ProcessTweet(tweet)
	}

	filterParams := &twitter.StreamFilterParams{
		Track:         []string{keyword},
		StallWarnings: twitter.Bool(true),
	}

	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Error().Err(err).Msgf("Error starting the stream")
		return err
	}

	go demux.HandleChan(stream.Messages)

	log.Info().Msgf("Starting stream on keyword %v", keyword)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Info().Msgf("Stopping stream - received %v", <-ch)

	stream.Stop()

	return nil
}
