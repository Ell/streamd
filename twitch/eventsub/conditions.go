package eventsub

type ChannelFollowCondition struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
	ModeratorUserId   string `json:"moderator_user_id"`
}
