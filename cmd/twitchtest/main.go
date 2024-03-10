package main

import (
	"fmt"
	"github.com/ell/streamd/twitch/eventsub"
	"github.com/ell/streamd/twitch/helix"
	"log"
	"os"
)

var clientId = "4yomih22hlk0hda17lofkd9e7e3kpe"

func main() {
	accessToken := os.Getenv("TWITCH_ACCESS_TOKEN")

	client, _ := helix.NewClient(clientId, accessToken)

	user, err := client.GetCurrentUser()
	if err != nil {
		log.Fatalf("Unable to get current user %s", err)
	}

	fmt.Printf("Got user: %v\n", user)

	fmt.Println("Connecting to eventsub")

	eventSubClient := eventsub.NewClient()

	var eventsCh = make(chan eventsub.Message)
	var statusCh = make(chan uint)

	go eventSubClient.Listen(eventsCh, statusCh)

	for {
		select {
		case event := <-eventsCh:
			{
				fmt.Printf("Got Event: %v\n", event)
				if event.Metadata.MessageType == "session_welcome" {
					var payload = new(eventsub.SessionPayload)

					err = eventsub.UnmarshalMessagePayload(&event, &payload)
					if err != nil {
						log.Fatalf("Unable to unmarshal event payload %s", err)
					}

					sessionId := payload.Session.Id

					err = client.SubscribeToChannelChatMessageEvent(user.Id, sessionId)
					if err != nil {
						log.Fatalf("Unable to subscribe to event %s", err)
					}
				}
			}
		case status := <-statusCh:
			{
				fmt.Printf("Got Status: %v\n", status)
			}
		}
	}
}
