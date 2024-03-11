package helix

import (
	"log"
)

func (c *Client) SubscribeToEvent(eventType, version string, condition interface{}, sessionId string) error {
	req := CreateEventSubSubscriptionRequest{
		SubscriptionType: eventType,
		Version:          version,
		Condition:        condition,
		Transport: Transport{
			Method:    "websocket",
			SessionId: sessionId,
		},
	}

	apiUrl := c.baseUrl + "/eventsub/subscriptions"

	_, err := makeHelixPostRequest[
		CreateEventSubSubscriptionRequest,
		CreateEventSubSubscriptionResponse,
	](c, apiUrl, req)

	if err != nil {
		log.Printf("Error making SubscribeToEvent request %s", err)
	}

	return nil
}
