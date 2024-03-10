package helix

import "fmt"

type APIResponse[T interface{}] struct {
	Data       []T `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination,omitempty"`
	Total        uint `json:"total,omitempty"`
	TotalCost    uint `json:"total_cost,omitempty"`
	MaxTotalCost uint `json:"max_total_cost,omitempty"`
}

type Transport struct {
	Method    string `json:"method"`
	SessionId string `json:"session_id,omitempty"`
}

type StartCommercialRequest struct {
	BroadcasterId string `json:"broadcaster_id"`
	Length        uint   `json:"length"`
}

type StartCommercialResponse struct {
	CommercialLength uint   `json:"length"`
	Message          string `json:"message"`
	RetryAfter       uint   `json:"retry_after"`
}

type CreateEventSubSubscriptionRequest struct {
	SubscriptionType string      `json:"type"`
	Version          string      `json:"version"`
	Condition        interface{} `json:"condition"`
	Transport        Transport   `json:"transport"`
}

type CreateEventSubSubscriptionResponse struct {
	Id               string      `json:"id"`
	Status           string      `json:"status"`
	SubscriptionType string      `json:"type"`
	Version          string      `json:"version"`
	Cost             uint        `json:"cost"`
	Condition        interface{} `json:"condition"`
	CreatedAt        string      `json:"created_at"`
	Transport        Transport   `json:"transport"`
}

type SendChatMessageRequest struct {
	BroadcasterId        string `json:"broadcaster_id"`
	SenderId             string `json:"sender_id"`
	Message              string `json:"message"`
	ReplyParentMessageId string `json:"reply_parent_message_id"`
}

type SendChatMessageResponse struct {
	MessageId  string `json:"message_id"`
	IsSent     bool   `json:"is_sent"`
	DropReason []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
}

// SendChatMessage Sends a message to the broadcaster’s chat room.
func (c *Client) SendChatMessage(req SendChatMessageRequest) (SendChatMessageResponse, error) {
	apiUrl := c.baseUrl + "/chat/messages"
	resp, err := makeHelixPostRequest[SendChatMessageRequest, SendChatMessageResponse](c, apiUrl, req)

	if err != nil {
		fmt.Printf("Error making SendChatMessage request %s", err)
		return SendChatMessageResponse{}, err
	}

	return resp.Data[0], nil
}

type GetUsersRequest struct {
	Id    string `url:"id,omitempty"`
	Login string `url:"login,omitempty"`
}

type GetUsersResponse struct {
	Id              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageUrl string `json:"profile_image_url"`
	OfflineImageUrl string `json:"offline_image_url"`
	ViewCount       uint   `json:"view_count"`
	Email           string `json:"email"`
	CreatedAt       string `json:"created_at"`
}

// GetCurrentUser Gets information about one or more users. Only fetches user associated with current client id
func (c *Client) GetCurrentUser() (GetUsersResponse, error) {
	var req struct{}

	apiUrl := c.baseUrl + "/users"

	resp, err := makeHelixGetRequest[interface{}, GetUsersResponse](c, apiUrl, req)
	if err != nil {
		fmt.Printf("Error making GetCurrentUser request %s", err)
	}

	return resp.Data[0], nil
}
