package eventsub

import (
	"time"
)

type Event interface {
	GetEventName() string
}

// ChannelFollowEvent channel.follow type
type ChannelFollowEvent struct {
	UserId               string    `json:"user_id"`
	UserLogin            string    `json:"user_login"`
	UserName             string    `json:"user_name"`
	BroadcasterUserId    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	FollowedAt           time.Time `json:"followed_at"`
}

func (e ChannelFollowEvent) GetEventName() string {
	return "channel.follow"
}

// ChannelAdBreakBeginsEvent channel.ad_break.begin
type ChannelAdBreakBeginsEvent struct {
	DurationSeconds      string    `json:"duration_seconds"`
	StartedAt            time.Time `json:"started_at"`
	IsAutomatic          string    `json:"is_automatic"`
	BroadcasterUserId    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	RequesterUserId      string    `json:"requester_user_id"`
	RequesterUserLogin   string    `json:"requester_user_login"`
	RequesterUserName    string    `json:"requester_user_name"`
}

func (e ChannelAdBreakBeginsEvent) GetEventName() string {
	return "channel.ad_break.begin"
}

// ChannelChatMessageEvent channel.chat.message
type ChannelChatMessageEvent struct {
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	ChatterUserId        string `json:"chatter_user_id"`
	ChatterUserLogin     string `json:"chatter_user_login"`
	ChatterUserName      string `json:"chatter_user_name"`
	MessageId            string `json:"message_id"`
	Message              struct {
		Text      string `json:"text"`
		Fragments []struct {
			Type      string      `json:"type"`
			Text      string      `json:"text"`
			Cheermote interface{} `json:"cheermote"`
			Emote     interface{} `json:"emote"`
			Mention   interface{} `json:"mention"`
		} `json:"fragments"`
	} `json:"message"`
	Color  string `json:"color"`
	Badges []struct {
		SetId string `json:"set_id"`
		Id    string `json:"id"`
		Info  string `json:"info"`
	} `json:"badges"`
	MessageType string `json:"message_type"`
	Cheer       struct {
		bits int
	} `json:"cheer,omitempty"`
	Reply struct {
		ParentMessageId   string `json:"parent_message_id"`
		ParentMessageBody string `json:"parent_message_body"`
		ParentUserId      string `json:"parent_user_id"`
		ParentUserName    string `json:"parent_user_name"`
		ParentUserLogin   string `json:"parent_user_login"`
		ThreadMessageId   string `json:"thread_message_id"`
		ThreadUserId      string `json:"thread_user_id"`
		ThreadUserName    string `json:"thread_user_name"`
		ThreadUserLogin   string `json:"thread_user_login"`
	} `json:"reply,omitempty"`
	ChannelPointsCustomRewardId string `json:"channel_points_custom_reward_id,omitempty"`
}

func (e ChannelChatMessageEvent) GetEventName() string {
	return "channel.chat.message"
}

// ChannelSubscribeEvent channel.subscribe
type ChannelSubscribeEvent struct {
	UserId               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Tier                 string `json:"tier"`
	IsGift               bool   `json:"is_gift"`
}

func (e ChannelSubscribeEvent) GetEventName() string {
	return "channel.subscribe"
}

// ChannelSubscriptionGiftEvent channel.subscription.gift
type ChannelSubscriptionGiftEvent struct {
	UserId               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Total                int    `json:"total"`
	Tier                 string `json:"tier"`
	CumulativeTotal      int    `json:"cumulative_total"`
	IsAnonymous          bool   `json:"is_anonymous"`
}

func (e ChannelSubscriptionGiftEvent) GetEventName() string {
	return "channel.subscription.gift"
}

// ChannelCheerEvent channel.cheer
type ChannelCheerEvent struct {
	IsAnonymous          bool   `json:"is_anonymous"`
	UserId               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Message              string `json:"message"`
	Bits                 int    `json:"bits"`
}

func (e ChannelCheerEvent) GetEventName() string {
	return "channel.cheer"
}

// ChannelRaidEvent channel.raid
type ChannelRaidEvent struct {
	FromBroadcasterUserId    string `json:"from_broadcaster_user_id"`
	FromBroadcasterUserLogin string `json:"from_broadcaster_user_login"`
	FromBroadcasterUserName  string `json:"from_broadcaster_user_name"`
	ToBroadcasterUserId      string `json:"to_broadcaster_user_id"`
	ToBroadcasterUserLogin   string `json:"to_broadcaster_user_login"`
	ToBroadcasterUserName    string `json:"to_broadcaster_user_name"`
	Viewers                  int    `json:"viewers"`
}

func (e ChannelRaidEvent) GetEventName() string {
	return "channel.raid"
}

// ChannelPointsCustomRewardRedemptionAddEvent channel.channel_points_custom_reward_redemption.add
type ChannelPointsCustomRewardRedemptionAddEvent struct {
	Id                   string `json:"id"`
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	UserId               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	UserInput            string `json:"user_input"`
	Status               string `json:"status"`
	Reward               struct {
		Id     string `json:"id"`
		Title  string `json:"title"`
		Cost   int    `json:"cost"`
		Prompt string `json:"prompt"`
	} `json:"reward"`
	RedeemedAt time.Time `json:"redeemed_at"`
}

func (e ChannelPointsCustomRewardRedemptionAddEvent) GetEventName() string {
	return "channel.channel_points_custom_reward_redemption.add"
}

// ChannelPollBeginEvent channel.poll.begin
type ChannelPollBeginEvent struct {
	Id                   string `json:"id"`
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Title                string `json:"title"`
	Choices              []struct {
		Id    string `json:"id"`
		Title string `json:"title"`
	} `json:"choices"`
	BitsVoting struct {
		IsEnabled     bool `json:"is_enabled"`
		AmountPerVote int  `json:"amount_per_vote"`
	} `json:"bits_voting"`
	ChannelPointsVoting struct {
		IsEnabled     bool `json:"is_enabled"`
		AmountPerVote int  `json:"amount_per_vote"`
	} `json:"channel_points_voting"`
	StartedAt time.Time `json:"started_at"`
	EndsAt    time.Time `json:"ends_at"`
}

func (e ChannelPollBeginEvent) GetEventName() string {
	return "channel.poll.begin"
}

// ChannelPollProgressEvent channel.poll.progress
type ChannelPollProgressEvent struct {
	Id                   string `json:"id"`
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Title                string `json:"title"`
	Choices              []struct {
		Id                 string `json:"id"`
		Title              string `json:"title"`
		BitsVotes          int    `json:"bits_votes"`
		ChannelPointsVotes int    `json:"channel_points_votes"`
		Votes              int    `json:"votes"`
	} `json:"choices"`
	BitsVoting struct {
		IsEnabled     bool `json:"is_enabled"`
		AmountPerVote int  `json:"amount_per_vote"`
	} `json:"bits_voting"`
	ChannelPointsVoting struct {
		IsEnabled     bool `json:"is_enabled"`
		AmountPerVote int  `json:"amount_per_vote"`
	} `json:"channel_points_voting"`
	StartedAt time.Time `json:"started_at"`
	EndsAt    time.Time `json:"ends_at"`
}

func (e ChannelPollProgressEvent) GetEventName() string {
	return "channel.poll.progress"
}

// ChannelPollEndEvent channel.poll.end
type ChannelPollEndEvent struct {
	Id                   string `json:"id"`
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Title                string `json:"title"`
	Choices              []struct {
		Id                 string `json:"id"`
		Title              string `json:"title"`
		BitsVotes          int    `json:"bits_votes"`
		ChannelPointsVotes int    `json:"channel_points_votes"`
		Votes              int    `json:"votes"`
	} `json:"choices"`
	BitsVoting struct {
		IsEnabled     bool `json:"is_enabled"`
		AmountPerVote int  `json:"amount_per_vote"`
	} `json:"bits_voting"`
	ChannelPointsVoting struct {
		IsEnabled     bool `json:"is_enabled"`
		AmountPerVote int  `json:"amount_per_vote"`
	} `json:"channel_points_voting"`
	Status    string    `json:"status"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
}

func (e ChannelPollEndEvent) GetEventName() string {
	return "channel.poll.end"
}

// ChannelPredictionBeginEvent channel.prediction.begin
type ChannelPredictionBeginEvent struct {
	Id                   string `json:"id"`
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Title                string `json:"title"`
	Outcomes             []struct {
		Id    string `json:"id"`
		Title string `json:"title"`
		Color string `json:"color"`
	} `json:"outcomes"`
	StartedAt time.Time `json:"started_at"`
	LocksAt   time.Time `json:"locks_at"`
}

func (e ChannelPredictionBeginEvent) GetEventName() string {
	return "channel.prediction.begin"
}

// ChannelPredictionProgressEvent channel.prediction.progress
type ChannelPredictionProgressEvent struct {
	Id                   string `json:"id"`
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Title                string `json:"title"`
	Outcomes             []struct {
		Id            string `json:"id"`
		Title         string `json:"title"`
		Color         string `json:"color"`
		Users         int    `json:"users,omitempty"`
		ChannelPoints int    `json:"channel_points,omitempty"`
		TopPredictors []struct {
			UserName          string      `json:"user_name"`
			UserLogin         string      `json:"user_login"`
			UserId            interface{} `json:"user_id"`
			ChannelPointsWon  interface{} `json:"channel_points_won"`
			ChannelPointsUsed int         `json:"channel_points_used"`
		} `json:"top_predictors"`
	} `json:"outcomes"`
	StartedAt time.Time `json:"started_at"`
	LocksAt   time.Time `json:"locks_at"`
}

func (e ChannelPredictionProgressEvent) GetEventName() string {
	return "channel.prediction.progress"
}

// ChannelPredictionEndEvent channel.prediction.end
type ChannelPredictionEndEvent struct {
	Id                   string `json:"id"`
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Title                string `json:"title"`
	WinningOutcomeId     string `json:"winning_outcome_id"`
	Outcomes             []struct {
		Id            string `json:"id"`
		Title         string `json:"title"`
		Color         string `json:"color"`
		Users         int    `json:"users"`
		ChannelPoints int    `json:"channel_points"`
		TopPredictors []struct {
			UserName          string      `json:"user_name"`
			UserLogin         string      `json:"user_login"`
			UserId            interface{} `json:"user_id"`
			ChannelPointsWon  *int        `json:"channel_points_won"`
			ChannelPointsUsed int         `json:"channel_points_used"`
		} `json:"top_predictors"`
	} `json:"outcomes"`
	Status    string    `json:"status"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
}

func (e ChannelPredictionEndEvent) GetEventName() string {
	return "channel.prediction.end"
}
