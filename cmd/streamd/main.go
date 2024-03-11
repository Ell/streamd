package main

import (
	"github.com/ell/streamd/twitch"
	"github.com/ell/streamd/twitch/eventsub"
	"log"
	"os"
)

func main() {
	accessToken := os.Getenv("TWITCH_ACCESS_TOKEN")
	clientId := os.Getenv("TWITCH_CLIENT_ID")

	client, err := twitch.NewClient(clientId, accessToken)
	if err != nil {
		log.Fatalf("Unable to create new twitch client %s\n", err)
	}

	events := make(chan eventsub.Message)
	client.AddEventListener(events)

	chatCondition := eventsub.ChannelChatMessageCondition{
		BroadcasterUserId: client.User.Id,
		UserId:            client.User.Id,
	}
	err = client.SubscribeToEvent(chatCondition)
	if err != nil {
		log.Fatalf("Unable to subscribe to chat messages %s\n", err)
	}

	log.Println("Twitch eventsub listener started")
	go client.Listen()

	for {
		event := <-events

		eventName := event.Metadata.MessageType
		log.Printf("Got event %s\n", eventName)
	}
}
