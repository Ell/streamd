package eventsub

type ChannelFollowPayload NotificationPayload[
	ChannelFollowEvent,
	ChannelFollowCondition,
]

type ChannelAdBreakBeginsPayload NotificationPayload[
	ChannelAdBreakBeginsEvent,
	ChannelAdBreakBeginsCondition,
]

type ChannelChatMessagePayload NotificationPayload[
	ChannelChatMessageEvent,
	ChannelChatMessageCondition,
]

type ChannelSubscribePayload NotificationPayload[
	ChannelSubscribeEvent,
	ChannelSubscribeCondition,
]

type ChannelSubscriptionGiftPayload NotificationPayload[
	ChannelSubscriptionGiftEvent,
	ChannelSubscriptionGiftCondition,
]

type ChannelCheerPayload NotificationPayload[
	ChannelCheerEvent,
	ChannelCheerCondition,
]

type ChannelRaidPayload NotificationPayload[
	ChannelRaidEvent,
	ChannelRaidCondition,
]

type ChannelPointsCustomRewardRedemptionAddPayload NotificationPayload[
	ChannelPointsCustomRewardRedemptionAddEvent,
	ChannelPointsCustomRewardRedemptionAddCondition,
]

type ChannelPollBeginPayload NotificationPayload[
	ChannelPollBeginEvent,
	ChannelPollBeginCondition,
]

type ChannelPollProgressPayload NotificationPayload[
	ChannelPollProgressEvent,
	ChannelPollProgressCondition,
]

type ChannelPollEndPayload NotificationPayload[
	ChannelPollEndEvent,
	ChannelPollEndCondition,
]

type ChannelPredictionBeginPayload NotificationPayload[
	ChannelPredictionBeginEvent,
	ChannelPredictionBeginCondition,
]

type ChannelPredictionProgressPayload NotificationPayload[
	ChannelPredictionProgressEvent,
	ChannelPredictionProgressCondition,
]

type ChannelPredictionEndEventPayload NotificationPayload[
	ChannelPredictionEndEvent,
	ChannelPredictionEndCondition,
]
