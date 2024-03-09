package eventsub

import (
	"fmt"
	"log"
	"time"
)

const (
	ClientConnected = iota
	ClientDisconnected
)

type Client struct {
	socketAddress string
	cancelCh      chan bool
	msgCh         chan Message
}

func NewClient() Client {
	cancelCh := make(chan bool)
	msgCh := make(chan Message)

	return Client{
		"wss://eventsub.wss.twitch.tv/ws?keepalive_timeout_seconds=30",
		cancelCh,
		msgCh,
	}
}

func NewTestClient(socketAddress string) Client {
	cancelCh := make(chan bool)
	msgCh := make(chan Message)

	return Client{
		socketAddress,
		cancelCh,
		msgCh,
	}
}

func (c *Client) Listen(events chan Message, clientStatus chan uint) {
	var statusCh = make(chan uint)
	var keepAlive uint = 30

	timeout := time.After(time.Second * time.Duration(keepAlive))

	go ConnectToSocket(c.socketAddress, c.msgCh, c.cancelCh, statusCh)

	for {
		select {
		case <-timeout:
			{
				log.Println("Socket timed out, disconnecting")
				c.cancelCh <- true
			}
		case status := <-statusCh:
			{
				if status == SocketDisconnected {
					clientStatus <- ClientDisconnected

					log.Println("Socket disconnected, reconnecting in 5 seconds...")

					time.Sleep(time.Second * 5)

					log.Println("Reconnecting to twitch eventsub")

					go ConnectToSocket(c.socketAddress, c.msgCh, c.cancelCh, statusCh)
				}

				if status == SocketConnected {
					clientStatus <- ClientConnected
				}
			}
		case message := <-c.msgCh:
			{
				// events <- message

				if message.Metadata.MessageType == "session_welcome" {
					var session = message.Payload.(SessionPayload)
					fmt.Printf("session %+v", session)

					//keepAlive = session.KeepaliveTimeoutSeconds
					//timeout = time.After(time.Duration(keepAlive) * time.Second)

					//out, _ := json.MarshalIndent(session, "", "\t")
					//fmt.Println(string(out))
				}

				if message.Metadata.MessageType == "session_keepalive" {
					timeout = time.After(time.Duration(keepAlive) * time.Second)
				}

				if message.Metadata.MessageType == "session_reconnect" {
					c.cancelCh <- true
				}
			}
		}
	}
}

func (c *Client) Disconnect() {
	c.cancelCh <- true
}
