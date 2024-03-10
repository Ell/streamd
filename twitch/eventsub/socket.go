package eventsub

import (
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

func connectToSocket(host string, out chan Message, cancel chan bool, status chan uint) {
	sockErrCh := make(chan error)

	status <- SocketConnecting

	ws, err := websocket.Dial(host, "", "http://localhost/")
	if err != nil {
		log.Fatalf("Could not create socket %s", err)
	}

	status <- SocketConnected

	defer func() {
		status <- SocketDisconnecting

		err := ws.Close()
		if err != nil {
			log.Fatalf("Unable to close socket %s", err)
		}

		status <- SocketDisconnected
	}()

	go func() {
		for {
			var message Message
			err = websocket.JSON.Receive(ws, &message)

			if err != nil {
				fmt.Printf("Unable to receive websocket message %s\n", err)
				sockErrCh <- err
				return
			}

			out <- message
		}
	}()

	for {
		select {
		case <-cancel:
			fmt.Println("Closing socket connection")
			status <- SocketClosing
			return
		case socketErr := <-sockErrCh:
			{
				fmt.Printf("Got socket error %s", socketErr)
				status <- SocketClosing
				return
			}
		}
	}
}
