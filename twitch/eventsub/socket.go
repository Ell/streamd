package eventsub

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

func connectToSocket(host string, out chan Message, cancel chan bool, finished chan bool) {
	sockErrCh := make(chan error)

	ws, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		log.Fatalf("Could not create socket %s", err)
	}

	defer func() {
		err := ws.Close()
		if err != nil {
			log.Fatalf("Unable to close socket %s", err)
		}
		finished <- true
	}()

	go func() {
		for {
			var message Message

			_, b, err := ws.ReadMessage()
			if err != nil {
				log.Printf("Unable to receive websocket message %s\n", err)
				sockErrCh <- err
				return
			}

			err = json.Unmarshal(b, &message)
			if err != nil {
				log.Printf("Unable to unmarshal websocket message %s\n", err)
				sockErrCh <- err
				return
			}

			out <- message
		}
	}()

	for {
		select {
		case <-cancel:
			{
				log.Println("Closing socket connection")
				return
			}
		case socketErr := <-sockErrCh:
			{
				log.Printf("Got socket error %s", socketErr)
				return
			}
		}
	}
}
