package twitter

import (
	"github-commit-reput/internal/commons"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/rs/zerolog/log"
)

var stream *twitter.Stream

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

	var err error
	stream, err = client.Streams.Filter(filterParams)
	if err != nil {
		log.Error().Err(err).Msgf("Error starting the stream")
		return err
	}

	go demux.HandleChan(stream.Messages)

	log.Info().Msgf("Starting stream on keyword %v", keyword)

	return nil
}

func StopStreaming() {
	log.Info().Msgf("Stopping stream")
	stream.Stop()
}
