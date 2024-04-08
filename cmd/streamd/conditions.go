package main

import (
	"github.com/ell/streamd/twitch/eventsub"
)

func createConditionList(userId string) []eventsub.Condition {
	var conditions = []eventsub.Condition{
		&eventsub.ChannelChatMessageCondition{
			BroadcasterUserId: userId,
			UserId:            userId,
		},
		&eventsub.ChannelSubscribeCondition{
			BroadcasterUserId: userId,
		},
		&eventsub.ChannelFollowCondition{
			BroadcasterUserId: userId,
			ModeratorUserId:   userId,
		},
		&eventsub.ChannelCheerCondition{
			BroadcasterUserId: userId,
		},
		&eventsub.ChannelRaidCondition{
			ToBroadcasterUserId: userId,
		},
		&eventsub.ChannelPointsCustomRewardRedemptionAddCondition{
			BroadcasterUserId: userId,
		},
	}

	return conditions
}
