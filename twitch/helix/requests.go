package helix

import (
	"log"
	"time"
)

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

type DeleteEventSubSubscriptionRequest struct {
	Id string `url:"id"`
}

type DeleteEventSubSubscriptionResponse struct{}

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

type GetChannelInformationRequest struct {
	BroadcasterId string `url:"broadcaster_id"`
}

type GetChannelInformationResponse struct {
	BroadcasterId               string   `json:"broadcaster_id"`
	BroadcasterLogin            string   `json:"broadcaster_login"`
	BroadcasterName             string   `json:"broadcaster_name"`
	BroadcasterLanguage         string   `json:"broadcaster_language"`
	GameId                      string   `json:"game_id"`
	GameName                    string   `json:"game_name"`
	Title                       string   `json:"title"`
	Delay                       int      `json:"delay"`
	Tags                        []string `json:"tags"`
	ContentClassificationLabels []string `json:"content_classification_labels"`
	IsBrandedContent            bool     `json:"is_branded_content"`
}

type CustomRewardsResponse struct {
	BroadcasterName     string      `json:"broadcaster_name"`
	BroadcasterLogin    string      `json:"broadcaster_login"`
	BroadcasterId       string      `json:"broadcaster_id"`
	Id                  string      `json:"id"`
	Image               interface{} `json:"image"`
	BackgroundColor     string      `json:"background_color"`
	IsEnabled           bool        `json:"is_enabled"`
	Cost                int         `json:"cost"`
	Title               string      `json:"title"`
	Prompt              string      `json:"prompt"`
	IsUserInputRequired bool        `json:"is_user_input_required"`
	MaxPerStreamSetting struct {
		IsEnabled    bool `json:"is_enabled"`
		MaxPerStream int  `json:"max_per_stream"`
	} `json:"max_per_stream_setting"`
	MaxPerUserPerStreamSetting struct {
		IsEnabled           bool `json:"is_enabled"`
		MaxPerUserPerStream int  `json:"max_per_user_per_stream"`
	} `json:"max_per_user_per_stream_setting"`
	GlobalCooldownSetting struct {
		IsEnabled             bool `json:"is_enabled"`
		GlobalCooldownSeconds int  `json:"global_cooldown_seconds"`
	} `json:"global_cooldown_setting"`
	IsPaused     bool `json:"is_paused"`
	IsInStock    bool `json:"is_in_stock"`
	DefaultImage struct {
		Url1X string `json:"url_1x"`
		Url2X string `json:"url_2x"`
		Url4X string `json:"url_4x"`
	} `json:"default_image"`
	ShouldRedemptionsSkipRequestQueue bool   `json:"should_redemptions_skip_request_queue"`
	RedemptionsRedeemedCurrentStream  int    `json:"redemptions_redeemed_current_stream,omitempty"`
	CooldownExpiresAt                 string `json:"cooldown_expires_at,omitempty"`
}

type CreateCustomRewardsRequest struct {
	Title string `json:"title"`
	Cost  string `json:"cost"`
}

type GetCustomRewardsRequest struct {
	BroadcasterId string `url:"broadcaster_id"`
}

type UpdateCustomRewardsParams struct {
	BroadcasterId string `url:"broadcaster_id"`
	Id            string `url:"id"`
}

type UpdateCustomRewardsRequest struct {
	Title                             string `json:"title,omitempty"`
	Prompt                            string `json:"prompt,omitempty"`
	Cost                              uint64 `json:"cost,omitempty"`
	BackgroundColor                   string `json:"background_color,omitempty"`
	IsEnabled                         bool   `json:"is_enabled,omitempty"`
	IsUserInputRequired               bool   `json:"is_user_input_required,omitempty"`
	IsMaxPerStreamEnabled             bool   `json:"is_max_per_stream_enabled,omitempty"`
	MaxPerStream                      uint64 `json:"max_per_stream,omitempty"`
	IsMaxPerUserPerStreamEnabled      bool   `json:"is_max_per_user_per_stream_enabled,omitempty"`
	MaxPerUserPerStream               uint64 `json:"max_per_user_per_stream,omitempty"`
	IsGlobalCooldownEnabled           bool   `json:"is_global_cooldown_enabled,omitempty"`
	GlobalCooldownSeconds             uint64 `json:"global_cooldown_seconds,omitempty"`
	IsPaused                          bool   `json:"is_paused,omitempty"`
	ShouldRedemptionsSkipRequestQueue bool   `json:"should_redemptions_skip_request_queue,omitempty"`
}

type DeleteCustomRewardsRequest struct {
	BroadcasterUserId string `url:"broadcaster_id"`
	Id                string `url:"id"`
}

type UpdateRedemptionStatusParams struct {
	Id            string `url:"id"`
	BroadcasterId string `url:"broadcaster_id"`
	RewardId      string `url:"reward_id"`
}

type UpdateRedemptionStatusRequest struct {
	Status string `json:"status"`
}

type UpdateRedemptionStatusResponse struct {
	BroadcasterName  string    `json:"broadcaster_name"`
	BroadcasterLogin string    `json:"broadcaster_login"`
	BroadcasterId    string    `json:"broadcaster_id"`
	Id               string    `json:"id"`
	UserId           string    `json:"user_id"`
	UserName         string    `json:"user_name"`
	UserLogin        string    `json:"user_login"`
	UserInput        string    `json:"user_input"`
	Status           string    `json:"status"`
	RedeemedAt       time.Time `json:"redeemed_at"`
	Reward           struct {
		Id     string `json:"id"`
		Title  string `json:"title"`
		Prompt string `json:"prompt"`
		Cost   int    `json:"cost"`
	} `json:"reward"`
}

type SendShoutoutRequest struct {
	FromBroadcasterId string `url:"from_broadcaster_id"`
	ToBroadcasterId   string `url:"to_broadcaster_id"`
	ModeratorId       string `url:"moderator_id"`
}

type SendAnnouncementParams struct {
	BroadcasterId string `url:"broadcaster_id"`
	ModeratorId   string `url:"moderator_id"`
}

type SendAnnouncementRequest struct {
	Message string `json:"message"`
	Color   string `json:"color,omitempty"`
}

type GetChattersParams struct {
	BroadcasterId string `url:"broadcaster_id,omitempty"`
	ModeratorId   string `url:"moderator_id,omitempty"`
	First         int    `url:"first,omitempty"`
	After         string `url:"after,omitempty"`
}

type GetChattersResponse struct {
	UserId    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
}

func (c *Client) GetChatters(req *GetChattersParams) ([]GetChattersResponse, error) {
	apiUrl := c.baseUrl + "/chat/chatters"

	if req.First == 0 {
		req.First = 1000
	}

	if req.ModeratorId != "" && req.BroadcasterId == "" {
		req.BroadcasterId = req.ModeratorId
	}

	if req.BroadcasterId != "" && req.ModeratorId == "" {
		req.ModeratorId = req.BroadcasterId
	}

	var results []GetChattersResponse

	resp, err := makeHelixGetRequest[GetChattersParams, GetChattersResponse](c, apiUrl, *req)
	if err != nil {
		log.Printf("Error making GetChatters request %s\n", err)
		return nil, err
	}

	results = append(results, resp.Data...)

	cursor := resp.Pagination.Cursor
	for cursor != "" {
		resp, err = makeHelixGetRequest[GetChattersParams, GetChattersResponse](c, apiUrl, *req)
		if err != nil {
			log.Printf("Error making GetChatters request %s\n", err)
			return nil, err
		}

		results = append(results, resp.Data...)
		cursor = resp.Pagination.Cursor
	}

	return results, nil
}

func (c *Client) SendChatAnnouncement(params *SendAnnouncementParams, req *SendAnnouncementRequest) error {
	apiUrl := c.baseUrl + "/chat/announcements?" + "broadcaster_id=" + params.BroadcasterId + "&moderator_id=" + params.ModeratorId

	_, err := makeHelixPostRequest[SendAnnouncementRequest, interface{}](c, apiUrl, *req)
	if err != nil {
		log.Printf("Error making SendChatAnnouncement request %s\n", err)
	}

	return err
}

func (c *Client) SendChatShoutout(req *SendShoutoutRequest) error {
	apiUrl := c.baseUrl + "/chat/shoutouts?" + "from_broadcaster_id=" + req.FromBroadcasterId + "&moderator_id=" + req.ModeratorId + "&to_broadcaster_id=" + req.ToBroadcasterId

	_, err := makeHelixPostRequest[SendShoutoutRequest, interface{}](c, apiUrl, *req)
	if err != nil {
		log.Printf("Error making SendChatShoutout request %s\n", err)
	}

	return err
}

func (c *Client) UpdateRedemptionStatus(params *UpdateRedemptionStatusParams, req *UpdateRedemptionStatusRequest) (*UpdateRedemptionStatusResponse, error) {
	apiUrl := c.baseUrl + "/channel_points/custom_rewards/redemptions?"
	apiUrl = apiUrl + "id=" + params.Id + "&broadcaster_id=" + params.BroadcasterId + "&reward_id=" + params.RewardId

	resp, err := makeHelixPatchRequest[UpdateRedemptionStatusRequest, UpdateRedemptionStatusResponse](c, apiUrl, *req)
	if err != nil {
		log.Printf("Error making UpdateRedemptionStatus request %s\n", err)
		return nil, err
	}

	return &resp.Data[0], nil
}

func (c *Client) DeleteCustomReward(req *DeleteCustomRewardsRequest) error {
	apiUrl := c.baseUrl + "/channel_points/custom_rewards"

	_, err := makeHelixDeleteRequest[DeleteCustomRewardsRequest, interface{}](c, apiUrl, *req)
	if err != nil {
		log.Printf("Error making DeleteCustomReward request %s\n", err)
	}
	return err
}

func (c *Client) UpdateCustomReward(params *UpdateCustomRewardsParams, req *UpdateCustomRewardsRequest) (*CustomRewardsResponse, error) {
	apiUrl := c.baseUrl + "/channel_points/custom_rewards?broadcaster_id=" + params.BroadcasterId + "&id=" + params.Id

	resp, err := makeHelixPatchRequest[UpdateCustomRewardsRequest, CustomRewardsResponse](c, apiUrl, *req)
	if err != nil {
		log.Printf("Error making UpdateCustomReward request %s\n", err)
	}

	return &resp.Data[0], nil
}

func (c *Client) GetCustomRewards(req *GetCustomRewardsRequest) ([]CustomRewardsResponse, error) {
	apiUrl := c.baseUrl + "/channel_points/custom_rewards"

	resp, err := makeHelixGetRequest[GetCustomRewardsRequest, CustomRewardsResponse](c, apiUrl, *req)
	if err != nil {
		log.Printf("Error making GetChannelInformation request %s\n", err)
		return nil, err
	}

	return resp.Data, nil
}

func (c *Client) CreateCustomReward(req *CreateCustomRewardsRequest) (*CustomRewardsResponse, error) {
	apiUrl := c.baseUrl + "/channel_points/custom_rewards"

	resp, err := makeHelixPostRequest[CreateCustomRewardsRequest, CustomRewardsResponse](c, apiUrl, *req)
	if err != nil {
		log.Printf("Error making GetChannelInformation request %s\n", err)
		return nil, err
	}

	return &resp.Data[0], nil
}

func (c *Client) GetChannelInformation(req *GetChannelInformationRequest) (*GetChannelInformationResponse, error) {
	apiUrl := c.baseUrl + "/channels"

	resp, err := makeHelixGetRequest[GetChannelInformationRequest, GetChannelInformationResponse](c, apiUrl, *req)
	if err != nil {
		log.Printf("Error making GetChannelInformation request %s\n", err)
		return nil, err
	}

	return &resp.Data[0], nil
}

func (c *Client) GetUsers(req *GetUsersRequest) (GetUsersResponse, error) {
	apiUrl := c.baseUrl + "/users"

	resp, err := makeHelixGetRequest[GetUsersRequest, GetUsersResponse](c, apiUrl, *req)
	if err != nil {
		log.Printf("Error making GetCurrentUser request %s\n", err)
	}

	return resp.Data[0], nil
}

func (c *Client) SubscribeToEvent(eventType, version string, condition interface{}, sessionId string) (string, error) {
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

	resp, err := makeHelixPostRequest[
		CreateEventSubSubscriptionRequest,
		CreateEventSubSubscriptionResponse,
	](c, apiUrl, req)

	if err != nil {
		log.Printf("Error making SubscribeToEvent request %s", err)
		return "", err
	}

	data := resp.Data[0]

	return data.Id, nil
}

func (c *Client) UnsubscribeFromEvent(eventId string) error {
	req := DeleteEventSubSubscriptionRequest{
		Id: eventId,
	}

	apiUrl := c.baseUrl + "/eventsub/subscriptions"

	_, err := makeHelixDeleteRequest[DeleteEventSubSubscriptionRequest, DeleteEventSubSubscriptionResponse](c, apiUrl, req)

	return err
}

func (c *Client) SendChatMessage(req *SendChatMessageRequest) (SendChatMessageResponse, error) {
	apiUrl := c.baseUrl + "/chat/messages"
	resp, err := makeHelixPostRequest[SendChatMessageRequest, SendChatMessageResponse](c, apiUrl, *req)

	if err != nil {
		log.Printf("Error making SendChatMessage request %s", err)
		return SendChatMessageResponse{}, err
	}

	return resp.Data[0], nil
}
