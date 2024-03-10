package helix

import (
	"fmt"
)

func (c *Client) subscribeToEvent(eventType, version string, condition interface{}, sessionId string) error {
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
		fmt.Printf("Error making SubscribeToEvent request %s", err)
	}

	return nil
}

func (c *Client) SubscribeToChannelChatMessageEvent(userId, sessionId string) error {
	eventType := "channel.chat.message"
	version := "1"
	condition := ChannelChatMessageCondition{
		BroadcasterUserId: userId,
		UserId:            userId,
	}

	err := c.subscribeToEvent(eventType, version, condition, sessionId)

	return err
}
