package twitch

import (
	"github.com/ell/streamd/twitch/eventsub"
	"github.com/ell/streamd/twitch/helix"
	"log"
	"sync"
	"time"
)

type Client struct {
	clientId             string
	accessToken          string
	sessionId            string
	helixClient          helix.Client
	eventSubClient       eventsub.Client
	subscriptions        map[string]eventsub.Condition
	events               chan eventsub.Message
	eventListeners       []chan eventsub.Message
	eventListenersMutex  *sync.Mutex
	statusListeners      []chan uint
	statusListenersMutex *sync.Mutex
	User                 User
}

func NewClient(clientId, accessToken string) (*Client, error) {
	err := ValidateAccessToken(accessToken)
	if err != nil {
		log.Fatalf("Unable to validate accessToken %s\n", err)
	}

	go func() {
		ticker := time.NewTicker(10 * time.Minute)

		for {
			<-ticker.C

			log.Println("Validating access token")

			err := ValidateAccessToken(accessToken)
			if err != nil {
				log.Fatalf("Unable to validate accessToken %s\n", err)
			}
		}
	}()

	var client Client

	eventListeners := make([]chan eventsub.Message, 0)
	statusListeners := make([]chan uint, 0)

	sessionId := ""
	subscriptions := make(map[string]eventsub.Condition)

	helixClient, err := helix.NewClient(clientId, accessToken)
	if err != nil {
		log.Println("Unable to create helix client", err)
		return &client, err
	}

	userData, err := helixClient.GetCurrentUser()
	if err != nil {
		log.Fatalf("Unable to get current user from helix client %s\n", err)
	}

	user := User{
		Id:          userData.Id,
		Username:    userData.Login,
		Displayname: userData.DisplayName,
	}

	eventSubClient := eventsub.NewClient()

	events := make(chan eventsub.Message)

	client = Client{
		clientId:        clientId,
		accessToken:     accessToken,
		sessionId:       sessionId,
		helixClient:     *helixClient,
		eventSubClient:  eventSubClient,
		subscriptions:   subscriptions,
		events:          events,
		eventListeners:  eventListeners,
		statusListeners: statusListeners,
		User:            user,
	}

	go client.handleEvents(events)

	return &client, nil
}

func (c *Client) Listen() {
	c.eventSubClient.Listen(c.events)
}

func (c *Client) handleEvents(events chan eventsub.Message) {
	for {
		message := <-events
		if message.Metadata.MessageType == "session_welcome" {
			var payload = new(eventsub.SessionPayload)

			err := eventsub.UnmarshalMessagePayload(&message, &payload)
			if err != nil {
				log.Fatalf("Unable to unmarshal event payload %s\n", err)
			}

			c.sessionId = payload.Session.Id

			err = c.subscribeToEvents()
			if err != nil {
				log.Fatalf("Unable to subscribe to events %s\n", err)
			}
		}

		if message.Metadata.MessageType == "notification" {
			for _, listener := range c.eventListeners {
				listener := listener

				go func() {
					listener <- message
				}()
			}
		}
	}
}

func (c *Client) subscribeToEvents() error {
	for _, condition := range c.subscriptions {
		log.Println("Subscribing to event", condition.GetEventName())

		err := c.subscribeToEvent(condition)
		if err != nil {
			log.Println("Unable to subscribe to event", err)
			return err
		}
	}

	return nil
}

func (c *Client) subscribeToEvent(condition eventsub.Condition) error {
	eventName := condition.GetEventName()
	eventVersion := condition.GetEventVersion()

	return c.helixClient.SubscribeToEvent(eventName, eventVersion, condition, c.sessionId)
}

func (c *Client) SubscribeToEvent(condition eventsub.Condition) error {
	c.subscriptions[condition.GetEventName()] = condition

	if c.sessionId != "" {
		err := c.subscribeToEvent(condition)
		if err != nil {
			log.Println("Unable to subscribe to event", err)
			return err
		}
	}

	return nil
}

func (c *Client) AddEventListener(messages chan eventsub.Message) {
	c.eventListeners = append(c.eventListeners, messages)
}

func (c *Client) AddStatusListener(statuses chan uint) {
	c.statusListeners = append(c.statusListeners, statuses)
}
