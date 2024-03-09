package eventsub

import "github.com/ell/streamd/twitch/helix"

type Message struct {
	Metadata MessageMetadata `json:"metadata"`
	Payload  interface{}     `json:"payload,omitempty"`
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
		KeepaliveTimeoutSeconds uint   `json:"keepalive_timeout_seconds"`
		ReconnectUrl            string `json:"reconnect_url,omitempty"`
	} `json:"session"`
}

type SubscriptionPayload[C interface{}] struct {
	Id        string          `json:"id"`
	Status    string          `json:"status"`
	Type      string          `json:"type"`
	Version   string          `json:"version"`
	Cost      uint            `json:"cost"`
	Condition C               `json:"condition"`
	Transport helix.Transport `json:"transport"`
	CreatedAt string          `json:"created_at"`
}

type NotificationPayload[E interface{}] struct {
	Event E `json:"event"`
}
