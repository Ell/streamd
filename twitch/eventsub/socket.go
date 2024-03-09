package eventsub

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

const (
	SocketConnecting = iota
	SocketConnected
	SocketClosing
	SocketDisconnecting
	SocketDisconnected
)

func ConnectToSocket(host string, out chan Message, cancel chan bool, status chan uint) {
	status <- SocketConnecting

	ws, err := websocket.Dial(host, "", "http://localhost/")
	if err != nil {
		log.Fatalf("Could not create socket %s", err)
	}

	status <- SocketConnected

	defer func(ws *websocket.Conn) {
		status <- SocketDisconnecting

		err := ws.Close()
		if err != nil {
			log.Fatalf("Unable to close socket %s", err)
		}

		status <- SocketDisconnected
	}(ws)

	for {
		select {
		case <-cancel:
			fmt.Println("Closing socket connection")
			status <- SocketClosing

			return
		default:
			message := Message{}

			err = websocket.JSON.Receive(ws, &message)
			if err != nil {
				fmt.Printf("Unable to receive websocket message %s", err)

				status <- SocketClosing
			}

			debug, _ := json.MarshalIndent(message, "", "\t")
			fmt.Println(string(debug))

			out <- message
		}
	}
}
