package eventsub

import (
	"log"
	"time"
)

const (
	ClientConnected = iota
	ClientDisconnected
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

func (c *Client) Listen(events chan Message, clientStatus chan uint) {
	var keepAlive = 10

	var statusCh = make(chan uint)
	var cancelCh = make(chan bool)
	var msgCh = make(chan Message)

	timer := time.NewTimer(time.Second * time.Duration(keepAlive))

	go connectToSocket(c.socketAddress, msgCh, cancelCh, statusCh)

	for {
		select {
		case <-timer.C:
			{
				log.Println("Socket timed out, disconnecting")
				cancelCh <- true
			}
		case status := <-statusCh:
			{
				if status == SocketDisconnected {
					clientStatus <- ClientDisconnected

					log.Println("Socket disconnected, reconnecting in 5 seconds...")

					time.Sleep(time.Second * 5)

					log.Println("Reconnecting to twitch eventsub")

					go connectToSocket(c.socketAddress, msgCh, cancelCh, statusCh)
				}

				if status == SocketConnected {
					clientStatus <- ClientConnected
				}
			}
		case message := <-msgCh:
			{
				events <- message

				// timer.Reset(time.Duration(keepAlive) * time.Second)

				if message.Metadata.MessageType == "session_welcome" {
					var payload = new(SessionPayload)

					err := UnmarshalMessagePayload(&message, payload)
					if err != nil {
						log.Fatalf("Could not unmarshal payload %s", err)
					}

					keepAlive = payload.Session.KeepaliveTimeoutSeconds
				}

				if message.Metadata.MessageType == "session_reconnect" {
					cancelCh <- true
				}
			}
		}
	}
}
