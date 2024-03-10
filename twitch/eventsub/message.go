package eventsub

import (
	"encoding/json"
	"github.com/ell/streamd/twitch/helix"
)

type Message struct {
	Metadata MessageMetadata        `json:"metadata"`
	Payload  map[string]interface{} `json:"payload,omitempty"`
}

func UnmarshalMessagePayload[P interface{}](message *Message, payload *P) error {
	b, err := json.Marshal(message.Payload)

	if err != nil {
		return err
	}

	return json.Unmarshal(b, payload)
}

type MessageMetadata struct {
	MessageId           string `json:"message_id"`
	MessageType         string `json:"message_type"`
	MessageTimestamp    string `json:"message_timestamp"`
	SubscriptionType    string `json:"subscription_type,omitempty"`
	SubscriptionVersion string `json:"subscription_version,omitempty"`
}

type SessionPayload struct {
	Session struct {
		Id                      string `json:"id"`
		Status                  string `json:"status"`
		ConnectedAt             string `json:"connected_at"`
		KeepaliveTimeoutSeconds int    `json:"keepalive_timeout_seconds"`
		ReconnectUrl            string `json:"reconnect_url,omitempty"`
	} `json:"session"`
}

type SubscriptionPayload[C interface{}] struct {
	Id        string          `json:"id"`
	Status    string          `json:"status"`
	Type      string          `json:"type"`
	Version   string          `json:"version"`
	Cost      int             `json:"cost"`
	Condition C               `json:"condition"`
	Transport helix.Transport `json:"transport"`
	CreatedAt string          `json:"created_at"`
}

type NotificationPayload[E, C interface{}] struct {
	Subscription SubscriptionPayload[C] `json:"subscription"`
	Event        E                      `json:"event"`
}
