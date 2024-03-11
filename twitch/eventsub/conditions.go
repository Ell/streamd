package eventsub

type Condition interface {
	GetEventName() string
	GetEventVersion() string
}

type ChannelFollowCondition struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
	ModeratorUserId   string `json:"moderator_user_id"`
}

func (c ChannelFollowCondition) GetEventName() string {
	return "channel.follow"
}

func (c ChannelFollowCondition) GetEventVersion() string {
	return "2"
}

type ChannelAdBreakBeginsCondition struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

func (c ChannelAdBreakBeginsCondition) GetEventName() string {
	return "channel.ad_break.begin"
}

func (c ChannelAdBreakBeginsCondition) GetEventVersion() string {
	return "2"
}

type ChannelChatMessageCondition struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
	UserId            string `json:"user_id"`
}

func (c ChannelChatMessageCondition) GetEventName() string {
	return "channel.chat.message"
}

func (c ChannelChatMessageCondition) GetEventVersion() string {
	return "1"
}

type ChannelSubscribeCondition struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

func (c ChannelSubscribeCondition) GetEventName() string {
	return "channel.subscribe"
}

func (c ChannelSubscribeCondition) GetEventVersion() string {
	return "1"
}

type ChannelSubscriptionGiftCondition struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

func (c ChannelSubscriptionGiftCondition) GetEventName() string {
	return "channel.subscription.gift"
}

func (c ChannelSubscriptionGiftCondition) GetEventVersion() string {
	return "1"
}

type ChannelCheerCondition struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

func (c ChannelCheerCondition) GetEventName() string {
	return "channel.cheer"
}

func (c ChannelCheerCondition) GetEventVersion() string {
	return "1"
}

type ChannelRaidCondition struct {
	ToBroadcasterUserId string `json:"to_broadcaster_user_id"`
}

func (c ChannelRaidCondition) GetEventName() string {
	return "channel.raid"
}

func (c ChannelRaidCondition) GetEventVersion() string {
	return "1"
}

type ChannelPointsCustomRewardRedemptionAddCondition struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
	RewardId          string `json:"reward_id,omitempty"`
}

func (c ChannelPointsCustomRewardRedemptionAddCondition) GetEventName() string {
	return "channel.channel_points_custom_reward_redemption.add"
}

func (c ChannelPointsCustomRewardRedemptionAddCondition) GetEventVersion() string {
	return "1"
}

type ChannelPollBeginCondition struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

func (c ChannelPollBeginCondition) GetEventName() string {
	return "channel.poll.begin"
}

func (c ChannelPollBeginCondition) GetEventVersion() string {
	return "1"
}

type ChannelPollProgressCondition struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

func (c ChannelPollProgressCondition) GetEventName() string {
	return "channel.poll.progress"
}

func (c ChannelPollProgressCondition) GetEventVersion() string {
	return "1"
}

type ChannelPollEndCondition struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

func (c ChannelPollEndCondition) GetEventName() string {
	return "channel.poll.end"
}

func (c ChannelPollEndCondition) GetEventVersion() string {
	return "1"
}

type ChannelPredictionBeginCondition struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

func (c ChannelPredictionBeginCondition) GetEventName() string {
	return "channel.prediction.begin"
}

func (c ChannelPredictionBeginCondition) GetEventVersion() string {
	return "1"
}

type ChannelPredictionProgressCondition struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

func (c ChannelPredictionProgressCondition) GetEventName() string {
	return "channel.prediction.progress"
}

func (c ChannelPredictionProgressCondition) GetEventVersion() string {
	return "1"
}

type ChannelPredictionEndCondition struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

func (c ChannelPredictionEndCondition) GetEventName() string {
	return "channel.prediction.end"
}

func (c ChannelPredictionEndCondition) GetEventVersion() string {
	return "1"
}
