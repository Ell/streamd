package server

import (
	"context"
	"strconv"

	"connectrpc.com/connect"
	twitchv1 "github.com/ell/streamd/rpc/twitch/v1"
	"github.com/ell/streamd/twitch/eventsub"
)

func messageToEventMessage(m eventsub.Message) (*twitchv1.EventMessage, error) {
	p := &eventsub.NotificationPayload[interface{}, interface{}]{}

	err := eventsub.UnmarshalMessagePayload(&m, p)
	if err != nil {
		return nil, err
	}

	resp := &twitchv1.EventMessage{
		Payload: &twitchv1.EventPayload{
			Subscriptions: &twitchv1.EventSubscription{
				Id:        p.Subscription.Id,
				Status:    p.Subscription.Status,
				Type:      p.Subscription.Type,
				Version:   p.Subscription.Version,
				Cost:      uint64(p.Subscription.Cost),
				Condition: nil,
				Transport: &twitchv1.EventTransport{},
				CreatedAt: "",
			},
			Event: nil,
		},
		Metadata: &twitchv1.EventMetadata{
			MessageId:           m.Metadata.MessageId,
			MessageType:         m.Metadata.MessageType,
			MessageTimeStamp:    m.Metadata.MessageTimestamp,
			SubscriptionType:    m.Metadata.SubscriptionType,
			SubscriptionVersion: m.Metadata.SubscriptionVersion,
		},
	}

	cond, err := eventsub.GetConditionForConditionName(m.Metadata.SubscriptionType)
	if err != nil {
		return nil, err
	}

	switch cond.(type) {
	case *eventsub.ChannelFollowCondition:
		{
			p := &eventsub.ChannelFollowPayload{}

			err := eventsub.UnmarshalMessagePayload(&m, p)
			if err != nil {
				return nil, err
			}

			resp.Payload.Subscriptions.Condition = &twitchv1.EventCondition{
				Condition: &twitchv1.EventCondition_ChannelFollow{
					ChannelFollow: &twitchv1.ChannelFollowCondition{
						BroadcasterUserId: p.Subscription.Condition.BroadcasterUserId,
						ModeratorUserId:   p.Subscription.Condition.ModeratorUserId,
					},
				},
			}

			resp.Payload.Event = &twitchv1.Event{
				Event: &twitchv1.Event_ChannelFollow{
					ChannelFollow: &twitchv1.ChannelFollowEvent{
						User: &twitchv1.EventUser{
							Id:       p.Event.UserId,
							Login:    p.Event.UserLogin,
							Username: p.Event.UserName,
						},
						Broadcaster: &twitchv1.EventUser{
							Id:       p.Event.BroadcasterUserId,
							Login:    p.Event.BroadcasterUserLogin,
							Username: p.Event.BroadcasterUserName,
						},
						FollowedAt: p.Event.FollowedAt.String(),
					},
				},
			}
		}
	case *eventsub.ChannelAdBreakBeginsCondition:
		{
			p := &eventsub.ChannelAdBreakBeginsPayload{}

			err := eventsub.UnmarshalMessagePayload(&m, p)
			if err != nil {
				return nil, err
			}

			resp.Payload.Subscriptions.Condition = &twitchv1.EventCondition{
				Condition: &twitchv1.EventCondition_ChannelAdBreak{
					ChannelAdBreak: &twitchv1.ChannelAdBreakCondition{
						BroadcasterUserId: p.Subscription.Condition.BroadcasterUserId,
					},
				},
			}

			resp.Payload.Event = &twitchv1.Event{
				Event: &twitchv1.Event_ChannelAdBreak{
					ChannelAdBreak: &twitchv1.ChannelAdBreakEvent{
						DurationSeconds: p.Event.DurationSeconds,
						StartedAt:       p.Event.StartedAt.String(),
						IsAutomatic:     p.Event.IsAutomatic,
						Broadcaster: &twitchv1.EventUser{
							Id:       p.Event.BroadcasterUserId,
							Login:    p.Event.BroadcasterUserLogin,
							Username: p.Event.BroadcasterUserName,
						},
						Requester: &twitchv1.EventUser{
							Id:       p.Event.RequesterUserId,
							Login:    p.Event.RequesterUserLogin,
							Username: p.Event.RequesterUserName,
						},
					},
				},
			}
		}
	case *eventsub.ChannelChatMessageCondition:
		{
			p := &eventsub.ChannelChatMessagePayload{}

			err := eventsub.UnmarshalMessagePayload(&m, p)
			if err != nil {
				return nil, err
			}

			resp.Payload.Subscriptions.Condition = &twitchv1.EventCondition{
				Condition: &twitchv1.EventCondition_ChannelChatMessage{
					ChannelChatMessage: &twitchv1.ChannelChatMessageCondition{
						BroadcasterUserId: p.Subscription.Condition.BroadcasterUserId,
						UserId:            p.Subscription.Condition.UserId,
					},
				},
			}

			badges := make([]*twitchv1.EventBadge, 0)
			for _, badge := range p.Event.Badges {
				eventBadge := &twitchv1.EventBadge{
					SetId: badge.SetId,
					Id:    badge.Id,
					Info:  badge.Info,
				}

				badges = append(badges, eventBadge)
			}

			resp.Payload.Event = &twitchv1.Event{
				Event: &twitchv1.Event_ChannelChatMessage{
					ChannelChatMessage: &twitchv1.ChannelChatMessageEvent{
						Broadcaster: &twitchv1.EventUser{
							Id:       p.Event.BroadcasterUserId,
							Login:    p.Event.BroadcasterUserLogin,
							Username: p.Event.BroadcasterUserName,
						},
						Chatter: &twitchv1.EventUser{
							Id:       p.Event.ChatterUserId,
							Login:    p.Event.ChatterUserLogin,
							Username: p.Event.ChatterUserName,
						},
						MessageId:   p.Event.MessageId,
						Color:       p.Event.Color,
						Badges:      badges,
						MessageType: p.Event.MessageType,
						Cheer: &twitchv1.EventCheer{
							Bits: uint64(p.Event.Cheer.Bits),
						},
						ChannelPointsCustomRewardId: p.Event.ChannelPointsCustomRewardId,
						Reply: &twitchv1.EventReply{
							ParentMessageId:   p.Event.Reply.ParentMessageId,
							ParentMessageBody: p.Event.Reply.ParentMessageBody,
							ParentUser: &twitchv1.EventUser{
								Id:       p.Event.Reply.ParentUserId,
								Login:    p.Event.Reply.ParentUserLogin,
								Username: p.Event.Reply.ParentUserName,
							},
							ThreadMessageId:   p.Event.Reply.ThreadMessageId,
							ThreadMessageBody: p.Event.Reply.ThreadMessageBody,
							ThreadUser: &twitchv1.EventUser{
								Id:       p.Event.Reply.ThreadUserId,
								Login:    p.Event.Reply.ThreadUserLogin,
								Username: p.Event.Reply.ThreadUserName,
							},
						},
					},
				},
			}
		}
	case *eventsub.ChannelSubscribeCondition:
		{
			p := &eventsub.ChannelSubscribePayload{}

			err := eventsub.UnmarshalMessagePayload(&m, p)
			if err != nil {
				return nil, err
			}

			resp.Payload.Subscriptions.Condition = &twitchv1.EventCondition{
				Condition: &twitchv1.EventCondition_ChannelSubscribe{
					ChannelSubscribe: &twitchv1.ChannelSubscribeCondition{
						BroadcasterUserId: p.Subscription.Condition.BroadcasterUserId,
					},
				},
			}

			resp.Payload.Event = &twitchv1.Event{
				Event: &twitchv1.Event_ChannelSubscribe{
					ChannelSubscribe: &twitchv1.ChannelSubscribeEvent{
						User: &twitchv1.EventUser{
							Id:       p.Event.UserId,
							Login:    p.Event.UserLogin,
							Username: p.Event.UserName,
						},
						Broadcaster: &twitchv1.EventUser{
							Id:       p.Event.BroadcasterUserId,
							Login:    p.Event.BroadcasterUserLogin,
							Username: p.Event.BroadcasterUserName,
						},
						Tier:   p.Event.Tier,
						IsGift: strconv.FormatBool(p.Event.IsGift),
					},
				},
			}
		}
	case *eventsub.ChannelCheerCondition:
		{
			p := &eventsub.ChannelCheerPayload{}

			err := eventsub.UnmarshalMessagePayload(&m, p)
			if err != nil {
				return nil, err
			}

			resp.Payload.Subscriptions.Condition = &twitchv1.EventCondition{
				Condition: &twitchv1.EventCondition_ChannelCheer{
					ChannelCheer: &twitchv1.ChannelCheerCondition{
						BroadcasterUserId: p.Subscription.Condition.BroadcasterUserId,
					},
				},
			}

			resp.Payload.Event = &twitchv1.Event{
				Event: &twitchv1.Event_ChannelCheer{
					ChannelCheer: &twitchv1.ChannelCheerEvent{
						User: &twitchv1.EventUser{
							Id:       p.Event.UserId,
							Login:    p.Event.UserLogin,
							Username: p.Event.UserName,
						},
						Broadcaster: &twitchv1.EventUser{
							Id:       p.Event.BroadcasterUserId,
							Login:    p.Event.BroadcasterUserLogin,
							Username: p.Event.BroadcasterUserName,
						},
						Bits:        uint64(p.Event.Bits),
						IsAnonymous: p.Event.IsAnonymous,
						Message:     p.Event.Message,
					},
				},
			}
		}
	case *eventsub.ChannelRaidCondition:
		{
			p := &eventsub.ChannelRaidPayload{}

			err := eventsub.UnmarshalMessagePayload(&m, p)
			if err != nil {
				return nil, err
			}

			resp.Payload.Subscriptions.Condition = &twitchv1.EventCondition{
				Condition: &twitchv1.EventCondition_ChannelRaid{
					ChannelRaid: &twitchv1.ChannelRaidCondition{
						ToBroadcasterUserId: p.Subscription.Condition.ToBroadcasterUserId,
					},
				},
			}

			resp.Payload.Event = &twitchv1.Event{
				Event: &twitchv1.Event_ChannelRaid{
					ChannelRaid: &twitchv1.ChannelRaidEvent{
						FromBroadcaster: &twitchv1.EventUser{
							Id:       p.Event.FromBroadcasterUserId,
							Login:    p.Event.FromBroadcasterUserLogin,
							Username: p.Event.FromBroadcasterUserName,
						},
						ToBroadcaster: &twitchv1.EventUser{
							Id:       p.Event.ToBroadcasterUserId,
							Login:    p.Event.ToBroadcasterUserLogin,
							Username: p.Event.ToBroadcasterUserName,
						},
						Viewers: uint64(p.Event.Viewers),
					},
				},
			}
		}
	case *eventsub.ChannelPointsCustomRewardRedemptionAddCondition:
		{
			p := &eventsub.ChannelPointsCustomRewardRedemptionAddPayload{}

			err := eventsub.UnmarshalMessagePayload(&m, p)
			if err != nil {
				return nil, err
			}

			resp.Payload.Subscriptions.Condition = &twitchv1.EventCondition{
				Condition: &twitchv1.EventCondition_ChannelPointsCustomRewardRedemption{
					ChannelPointsCustomRewardRedemption: &twitchv1.ChannelPointsCustomRewardRedemptionCondition{
						BroadcasterUserId: p.Subscription.Condition.BroadcasterUserId,
					},
				},
			}

			resp.Payload.Event = &twitchv1.Event{
				Event: &twitchv1.Event_ChannelPointsCustomRewardRedemption{
					ChannelPointsCustomRewardRedemption: &twitchv1.ChannelPointsCustomRewardRedemptionEvent{
						Id: p.Event.Id,
						User: &twitchv1.EventUser{
							Id:       p.Event.UserId,
							Login:    p.Event.UserLogin,
							Username: p.Event.UserName,
						},
						Broadcaster: &twitchv1.EventUser{
							Id:       p.Event.BroadcasterUserId,
							Login:    p.Event.BroadcasterUserLogin,
							Username: p.Event.BroadcasterUserName,
						},
						Status: p.Event.Status,
						Reward: &twitchv1.EventChannelPointsCustomReward{
							Id:     p.Event.Reward.Id,
							Title:  p.Event.Reward.Title,
							Prompt: p.Event.Reward.Prompt,
							Cost:   uint64(p.Event.Reward.Cost),
						},
						RedeemedAt: p.Event.RedeemedAt.String(),
					},
				},
			}
		}
	}

	return resp, nil
}

func (s *Server) SubscribeToEvents(
	ctx context.Context,
	req *connect.Request[twitchv1.SubscribeToEventsRequest],
	resp *connect.ServerStream[twitchv1.SubscribeToEventsResponse],
) error {
	events := s.TwitchClient.AddListener()
	defer close(*events)

	for {
		select {
		case event := <-*events:
			{
				message, err := messageToEventMessage(event)
				if err != nil {
					return connect.NewError(connect.CodeInternal, err)
				}

				response := &twitchv1.SubscribeToEventsResponse{
					Message: message,
				}

				err = resp.Send(response)
				if err != nil {
					return connect.NewError(connect.CodeInternal, err)
				}
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
