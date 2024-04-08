package twitch

import (
	"github.com/ell/streamd/twitch/eventsub"
	"github.com/ell/streamd/twitch/helix"
	"log"
	"sync"
)

type subscription struct {
	Condition eventsub.Condition
	Id        string
}

type Client struct {
	ClientId      string
	AccessToken   string
	sessionId     string
	subscriptions []*subscription
	listeners     []chan eventsub.Message
	Helix         helix.Client
	Eventsub      eventsub.Client
	lock          sync.Mutex
}

func NewClient(clientId, accessToken string) *Client {
	eventsubClient := eventsub.NewClient()
	helixClient := helix.NewClient(clientId, accessToken)

	client := &Client{
		ClientId:    clientId,
		AccessToken: accessToken,
		Eventsub:    eventsubClient,
		Helix:       *helixClient,
	}

	return client
}

func NewTestClient(clientId, accessToken, wsAddress, apiAddress string) *Client {
	eventsubClient := eventsub.NewTestClient(wsAddress)
	helixClient := helix.NewTestClient(clientId, accessToken, apiAddress)

	client := &Client{
		ClientId:    clientId,
		AccessToken: accessToken,
		Eventsub:    eventsubClient,
		Helix:       *helixClient,
	}

	return client
}

func (c *Client) AddListener() *chan eventsub.Message {
	c.lock.Lock()
	defer c.lock.Unlock()

	ch := make(chan eventsub.Message)
	c.listeners = append(c.listeners, ch)

	return &ch
}

func (c *Client) Listen(conditions ...eventsub.Condition) {
	c.lock.Lock()

	for _, condition := range conditions {
		sub := subscription{
			Condition: condition,
			Id:        "",
		}

		c.subscriptions = append(c.subscriptions, &sub)
	}

	c.lock.Unlock()

	events := make(chan eventsub.Message)

	go c.Eventsub.Listen(&events)

	for {
		message := <-events

		switch message.Metadata.MessageType {
		case "session_welcome":
			{
				var payload = &eventsub.SessionPayload{}

				err := eventsub.UnmarshalMessagePayload(&message, &payload)
				if err != nil {
					log.Fatalf("Unable to unmarshal event payload %s\n", err)
				}

				log.Println("Connected to Eventsub")

				for _, sub := range c.subscriptions {
					id, err := c.Helix.SubscribeToEvent(sub.Condition.GetEventName(), sub.Condition.GetEventVersion(), sub.Condition, payload.Session.Id)
					if err != nil {
						log.Fatalf("Unable to subscribe to Eventsub event %s\n", err)
					}

					log.Println("Subscribed to event with id", id)

					c.lock.Lock()
					sub.Id = id
					c.lock.Unlock()
				}

				c.sessionId = payload.Session.Id
			}
		case "notification":
			{
				conditionName := message.Metadata.SubscriptionType
				condition, err := eventsub.GetConditionForConditionName(message.Metadata.SubscriptionType)
				if err != nil {
					log.Println("Unknown subscription type received", conditionName)
					continue
				}

				log.Printf("Got notification for %s event\n", condition.GetEventName())

				c.lock.Lock()

				listeners := make([]chan eventsub.Message, 0)
				for i, listener := range c.listeners {
					select {
					case <-listener:
						listeners = append(c.listeners[:i], c.listeners[i+1:]...)
						c.listeners = listeners
						continue
					default:
						listener <- message
						continue
					}
				}

				c.lock.Unlock()
			}
		}
	}
}

func (c *Client) subscribeToEvent(condition eventsub.Condition) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	subId := ""

	if c.sessionId != "" {
		subId, _ = c.Helix.SubscribeToEvent(condition.GetEventName(), condition.GetEventVersion(), condition, c.sessionId)
	}

	c.subscriptions = append(c.subscriptions, &subscription{
		Condition: condition,
		Id:        subId,
	})

	return nil
}
