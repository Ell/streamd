package eventsub

import (
	"log"
	"time"
)

type Client struct {
	socketAddress string
}

func NewClient() Client {
	return Client{"wss://eventsub.wss.twitch.tv/ws?keepalive_timeout_seconds=30"}
}

func NewTestClient(socketAddress string) Client {
	return Client{socketAddress}
}

func (c *Client) Listen(events *chan Message) {
	keepAlive := time.Second * 30
	timer := time.NewTimer(keepAlive)

	socketFinishedCh := make(chan bool)
	cancelCh := make(chan bool)
	msgCh := make(chan Message)

	go connectToSocket(c.socketAddress, msgCh, cancelCh, socketFinishedCh)

	for {
		select {
		case <-socketFinishedCh:
			{
				log.Println("Socket disconnected, reconnecting in 5 seconds...")
				time.Sleep(time.Second * 5)
				log.Println("Reconnecting to twitch eventsub")
				go connectToSocket(c.socketAddress, msgCh, cancelCh, socketFinishedCh)
				log.Println("Reconnected to twitch eventsub")
			}
		case message := <-msgCh:
			{
				*events <- message

				timer.Reset(keepAlive)

				if message.Metadata.MessageType == "session_welcome" {
					var payload = new(SessionPayload)

					err := UnmarshalMessagePayload(&message, payload)
					if err != nil {
						log.Fatalf("Could not unmarshal payload %s", err)
					}

					keepAlive = time.Second * time.Duration(payload.Session.KeepaliveTimeoutSeconds+10)
				}

				if message.Metadata.MessageType == "session_reconnect" {
					cancelCh <- true
				}
			}
		case <-timer.C:
			{
				log.Println("Socket timed out, disconnecting")
				cancelCh <- true
			}
		}
	}
}
