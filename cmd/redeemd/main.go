package main

import (
	"connectrpc.com/connect"
	"context"
	"flag"
	"github.com/ell/streamd"
	twitchv1 "github.com/ell/streamd/rpc/twitch/v1"
	"github.com/ell/streamd/rpc/twitch/v1/twitchv1connect"
	"github.com/ell/streamd/server"
	"log"
	"net/http"
	"path/filepath"
)

const AssetsPath = "C:\\Home\\projects\\stream\\web\\streamd\\assets"

type EventHandlerClient struct {
	audioPlayer *AudioPlayer
}

func main() {
	apiKey := flag.String("api_key", "apikey123", "service api key")
	apiHost := flag.String("api_host", "http://localhost:8065", "streamd host")

	if *apiKey == "" {
		log.Fatalf("api_key not provided")
	}

	if *apiHost == "" {
		log.Fatalf("api_host not provided")
	}

	interceptors := connect.WithInterceptors(server.NewAuthInterceptor(*apiKey))

	client := twitchv1connect.NewTwitchServiceClient(
		http.DefaultClient,
		*apiHost,
		interceptors,
	)

	log.Println("Watching for events...")

	audioPlayer, err := NewAudioPlayer(AssetsPath)
	if err != nil {
		log.Fatal(err)
	}

	handlerClient := EventHandlerClient{
		audioPlayer: audioPlayer,
	}

	eventsCtx := context.Background()
	eventsCh := make(chan *twitchv1.SubscribeToEventsResponse)

	go func() {
		res, err := client.SubscribeToEvents(
			eventsCtx,
			connect.NewRequest(&twitchv1.SubscribeToEventsRequest{}),
		)

		if err != nil {
			log.Fatal(err)
		}

		for res.Receive() != false {
			eventsCh <- res.Msg()
		}
	}()

	for {
		select {
		case event := <-eventsCh:
			{
				switch event.Message.Payload.Event.Event.(type) {
				case *twitchv1.Event_ChannelPointsCustomRewardRedemption:
					{
						redemptionEvent := event.Message.Payload.Event.Event.(*twitchv1.Event_ChannelPointsCustomRewardRedemption)
						redemption := redemptionEvent.ChannelPointsCustomRewardRedemption

						err := handlerClient.HandleChannelPointsRedemption(redemption)
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			}
		case <-eventsCtx.Done():
			{
				log.Fatal("Events context completed")
			}
		}
	}
}

func (e *EventHandlerClient) HandleChannelPointsRedemption(redemption *twitchv1.ChannelPointsCustomRewardRedemptionEvent) error {
	title := redemption.Reward.Title

	fartRedeems := []string{"Fart", "Fart 2"}

	switch {
	case streamd.SliceIncludes(fartRedeems, title) == true:
		{
			go func() {
				err := e.HandleFartRedemption(title)
				if err != nil {
					log.Fatal(err)
				}
			}()
		}

		break
	}

	return nil
}

func (e *EventHandlerClient) HandleFartRedemption(fartType string) error {
	log.Printf("Got fart %+v\n", fartType)

	fartPath := filepath.Join(AssetsPath, "fart.wav")

	err := e.audioPlayer.Play(fartPath)

	return err
}
