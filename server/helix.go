package server

import (
	"context"
	"strconv"

	"connectrpc.com/connect"
	twitchv1 "github.com/ell/streamd/rpc/twitch/v1"
	"github.com/ell/streamd/twitch/helix"
)

func (s *Server) SendChatMessage(
	ctx context.Context, req *connect.Request[twitchv1.SendChatMessageRequest],
) (*connect.Response[twitchv1.SendChatMessageResponse], error) {
	helixSendChatMessageRequest := &helix.SendChatMessageRequest{
		BroadcasterId:        req.Msg.GetBroadcasterId(),
		SenderId:             req.Msg.GetSenderId(),
		Message:              req.Msg.GetMessage(),
		ReplyParentMessageId: req.Msg.GetReplyParentMessageId(),
	}

	r, err := s.TwitchClient.Helix.SendChatMessage(helixSendChatMessageRequest)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&twitchv1.SendChatMessageResponse{
		MessageId: r.MessageId,
		IsSent:    r.IsSent,
	})

	res.Header().Set("SendChatMessage-Version", "v1")

	return res, nil
}

func (s *Server) GetUsers(
	ctx context.Context,
	req *connect.Request[twitchv1.GetUsersRequest],
) (*connect.Response[twitchv1.GetUsersResponse], error) {
	helixGetUserRequest := &helix.GetUsersRequest{}

	switch req.Msg.User.(type) {
	case *twitchv1.GetUsersRequest_Id:
		{
			helixGetUserRequest.Id = req.Msg.GetId()
		}
	case *twitchv1.GetUsersRequest_Login:
		{
			helixGetUserRequest.Login = req.Msg.GetLogin()
		}
	}

	r, err := s.TwitchClient.Helix.GetUsers(helixGetUserRequest)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&twitchv1.GetUsersResponse{
		Id:              r.Id,
		Login:           r.Login,
		DisplayName:     r.DisplayName,
		UserType:        r.Type,
		BroadcasterType: r.BroadcasterType,
		Description:     r.Description,
		ProfileImageUrl: r.ProfileImageUrl,
		OfflineImageUrl: r.OfflineImageUrl,
		ViewCount:       uint64(r.ViewCount),
		Email:           r.Email,
		CreatedAt:       r.CreatedAt,
	})

	res.Header().Set("GetUsers-Version", "v1")
	return res, nil
}

func (s *Server) CreateCustomReward(
	ctx context.Context,
	req *connect.Request[twitchv1.CreateCustomRewardRequest],
) (*connect.Response[twitchv1.CreateCustomRewardResponse], error) {
	helixCreateCustomRewardRequest := &helix.CreateCustomRewardsRequest{
		Title: req.Msg.Title,
		Cost:  strconv.FormatUint(req.Msg.Cost, 10),
	}

	r, err := s.TwitchClient.Helix.CreateCustomReward(helixCreateCustomRewardRequest)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&twitchv1.CreateCustomRewardResponse{
		Reward: &twitchv1.CustomRewardResponse{
			BroadcasterName:     r.BroadcasterName,
			BroadcasterLogin:    r.BroadcasterLogin,
			BroadcasterId:       r.BroadcasterId,
			Id:                  r.Id,
			BackgroundColor:     r.BackgroundColor,
			IsEnabled:           r.IsEnabled,
			Cost:                uint32(r.Cost),
			Title:               r.Title,
			Prompt:              r.Prompt,
			IsUserInputRequired: r.IsUserInputRequired,
			MaxPerStreamSetting: &twitchv1.MaxPerStreamSetting{
				IsEnabled:    r.MaxPerStreamSetting.IsEnabled,
				MaxPerStream: uint32(r.MaxPerStreamSetting.MaxPerStream),
			},
			MaxPerUserPerStreamSetting: &twitchv1.MaxPerUserPerStreamSetting{
				IsEnabled:           r.MaxPerUserPerStreamSetting.IsEnabled,
				MaxPerUserPerStream: uint32(r.MaxPerUserPerStreamSetting.MaxPerUserPerStream),
			},
			GlobalCooldownSetting: &twitchv1.GlobalCooldownSetting{
				IsEnabled:             r.GlobalCooldownSetting.IsEnabled,
				GlobalCooldownSeconds: uint32(r.GlobalCooldownSetting.GlobalCooldownSeconds),
			},
			IsPaused:  r.IsPaused,
			IsInStock: r.IsInStock,
			DefaultImage: &twitchv1.DefaultImage{
				Url_1X: r.DefaultImage.Url1X,
				Url_2X: r.DefaultImage.Url2X,
				Url_4X: r.DefaultImage.Url4X,
			},
			ShouldRedemptionsSkipRequestQueue: r.ShouldRedemptionsSkipRequestQueue,
			RedemptionsRedeemedCurrentStream:  strconv.FormatInt(int64(r.RedemptionsRedeemedCurrentStream), 10),
			CooldownExpiresAt:                 r.CooldownExpiresAt,
		},
	})

	res.Header().Set("CreateCustomReward-Version", "v1")

	return res, nil
}

func (s *Server) GetCustomRewards(
	ctx context.Context,
	req *connect.Request[twitchv1.GetCustomRewardsRequest],
) (*connect.Response[twitchv1.GetCustomRewardsResponse], error) {
	helixGetCustomRewardsRequest := &helix.GetCustomRewardsRequest{
		BroadcasterId: req.Msg.BroadcasterId,
	}

	r, err := s.TwitchClient.Helix.GetCustomRewards(helixGetCustomRewardsRequest)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	customRewards := make([]*twitchv1.CustomRewardResponse, 0, len(r))
	for _, reward := range r {
		customRewards = append(customRewards, &twitchv1.CustomRewardResponse{
			BroadcasterName:     reward.BroadcasterName,
			BroadcasterLogin:    reward.BroadcasterLogin,
			BroadcasterId:       reward.BroadcasterId,
			Id:                  reward.Id,
			BackgroundColor:     reward.BackgroundColor,
			IsEnabled:           reward.IsEnabled,
			Cost:                uint32(reward.Cost),
			Title:               reward.Title,
			Prompt:              reward.Prompt,
			IsUserInputRequired: reward.IsUserInputRequired,
			MaxPerStreamSetting: &twitchv1.MaxPerStreamSetting{
				IsEnabled:    reward.MaxPerStreamSetting.IsEnabled,
				MaxPerStream: uint32(reward.MaxPerStreamSetting.MaxPerStream),
			},
			MaxPerUserPerStreamSetting: &twitchv1.MaxPerUserPerStreamSetting{
				IsEnabled:           reward.MaxPerUserPerStreamSetting.IsEnabled,
				MaxPerUserPerStream: uint32(reward.MaxPerUserPerStreamSetting.MaxPerUserPerStream),
			},
			GlobalCooldownSetting: &twitchv1.GlobalCooldownSetting{
				IsEnabled:             reward.GlobalCooldownSetting.IsEnabled,
				GlobalCooldownSeconds: uint32(reward.GlobalCooldownSetting.GlobalCooldownSeconds),
			},
			IsPaused:  reward.IsPaused,
			IsInStock: reward.IsInStock,
			DefaultImage: &twitchv1.DefaultImage{
				Url_1X: reward.DefaultImage.Url1X,
				Url_2X: reward.DefaultImage.Url2X,
				Url_4X: reward.DefaultImage.Url4X,
			},
			ShouldRedemptionsSkipRequestQueue: reward.ShouldRedemptionsSkipRequestQueue,
			RedemptionsRedeemedCurrentStream:  strconv.FormatInt(int64(reward.RedemptionsRedeemedCurrentStream), 10),
			CooldownExpiresAt:                 reward.CooldownExpiresAt,
		})
	}

	res := connect.NewResponse(&twitchv1.GetCustomRewardsResponse{
		Rewards: customRewards,
	})

	res.Header().Set("GetCustomRewards-Version", "v1")

	return res, nil
}

func (s *Server) UpdateCustomReward(
	ctx context.Context,
	req *connect.Request[twitchv1.UpdateCustomRewardRequest],
) (*connect.Response[twitchv1.UpdateCustomRewardResponse], error) {
	helixUpdateCustomRewardRequest := &helix.UpdateCustomRewardsRequest{}

	if req.Msg.Title != nil {
		helixUpdateCustomRewardRequest.Title = *req.Msg.Title
	}

	if req.Msg.Prompt != nil {
		helixUpdateCustomRewardRequest.Prompt = *req.Msg.Prompt
	}

	if req.Msg.Cost != nil {
		helixUpdateCustomRewardRequest.Cost = *req.Msg.Cost
	}

	if req.Msg.BackgroundColor != nil {
		helixUpdateCustomRewardRequest.BackgroundColor = *req.Msg.BackgroundColor
	}

	if req.Msg.IsEnabled != nil {
		helixUpdateCustomRewardRequest.IsEnabled = *req.Msg.IsEnabled
	}

	if req.Msg.IsUserInputRequired != nil {
		helixUpdateCustomRewardRequest.IsUserInputRequired = *req.Msg.IsUserInputRequired
	}

	if req.Msg.IsMaxPerStreamEnabled != nil {
		helixUpdateCustomRewardRequest.IsMaxPerStreamEnabled = *req.Msg.IsMaxPerStreamEnabled
	}

	if req.Msg.MaxPerStream != nil {
		helixUpdateCustomRewardRequest.MaxPerStream = *req.Msg.MaxPerStream
	}

	if req.Msg.IsMaxPerUserPerStreamEnabled != nil {
		helixUpdateCustomRewardRequest.IsMaxPerUserPerStreamEnabled = *req.Msg.IsMaxPerUserPerStreamEnabled
	}

	if req.Msg.MaxPerUserPerStream != nil {
		helixUpdateCustomRewardRequest.MaxPerUserPerStream = *req.Msg.MaxPerUserPerStream
	}

	if req.Msg.IsGlobalCooldownEnabled != nil {
		helixUpdateCustomRewardRequest.IsGlobalCooldownEnabled = *req.Msg.IsGlobalCooldownEnabled
	}

	if req.Msg.GlobalCooldownSeconds != nil {
		helixUpdateCustomRewardRequest.GlobalCooldownSeconds = *req.Msg.GlobalCooldownSeconds
	}

	if req.Msg.IsPaused != nil {
		helixUpdateCustomRewardRequest.IsPaused = *req.Msg.IsPaused
	}

	if req.Msg.ShouldRedemptionsSkipRequestQueue != nil {
		helixUpdateCustomRewardRequest.ShouldRedemptionsSkipRequestQueue = *req.Msg.ShouldRedemptionsSkipRequestQueue
	}

	params := &helix.UpdateCustomRewardsParams{
		BroadcasterId: req.Msg.BroadcasterId,
		Id:            req.Msg.Id,
	}

	// make request
	r, err := s.TwitchClient.Helix.UpdateCustomReward(params, helixUpdateCustomRewardRequest)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&twitchv1.UpdateCustomRewardResponse{
		Reward: &twitchv1.CustomRewardResponse{
			BroadcasterName:     r.BroadcasterId,
			BroadcasterLogin:    r.BroadcasterLogin,
			BroadcasterId:       r.BroadcasterId,
			Id:                  r.Id,
			BackgroundColor:     r.BackgroundColor,
			IsEnabled:           r.IsEnabled,
			Cost:                uint32(r.Cost),
			Title:               r.Title,
			Prompt:              r.Prompt,
			IsUserInputRequired: r.IsUserInputRequired,
			MaxPerStreamSetting: &twitchv1.MaxPerStreamSetting{
				IsEnabled:    r.MaxPerStreamSetting.IsEnabled,
				MaxPerStream: uint32(r.MaxPerStreamSetting.MaxPerStream),
			},
			MaxPerUserPerStreamSetting: &twitchv1.MaxPerUserPerStreamSetting{
				IsEnabled:           r.MaxPerUserPerStreamSetting.IsEnabled,
				MaxPerUserPerStream: uint32(r.MaxPerUserPerStreamSetting.MaxPerUserPerStream),
			},
			GlobalCooldownSetting: &twitchv1.GlobalCooldownSetting{
				IsEnabled:             r.GlobalCooldownSetting.IsEnabled,
				GlobalCooldownSeconds: uint32(r.GlobalCooldownSetting.GlobalCooldownSeconds),
			},
			IsPaused:  r.IsPaused,
			IsInStock: r.IsInStock,
			DefaultImage: &twitchv1.DefaultImage{
				Url_1X: r.DefaultImage.Url1X,
				Url_2X: r.DefaultImage.Url2X,
				Url_4X: r.DefaultImage.Url4X,
			},
			ShouldRedemptionsSkipRequestQueue: r.ShouldRedemptionsSkipRequestQueue,
			CooldownExpiresAt:                 r.CooldownExpiresAt,
		},
	})

	res.Header().Set("UpdateCustomReward-Version", "v1")

	return res, nil
}

func (s *Server) DeleteCustomReward(
	ctx context.Context,
	req *connect.Request[twitchv1.DeleteCustomRewardRequest],
) (*connect.Response[twitchv1.DeleteCustomRewardResponse], error) {
	helixDeleteCustomRewardRequest := &helix.DeleteCustomRewardsRequest{
		BroadcasterUserId: req.Msg.BroadcasterUserId,
		Id:                req.Msg.Id,
	}

	err := s.TwitchClient.Helix.DeleteCustomReward(helixDeleteCustomRewardRequest)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&twitchv1.DeleteCustomRewardResponse{})

	res.Header().Set("DeleteCustomReward-Version", "v1")

	return res, nil
}

func (s *Server) GetChannelInformation(
	ctx context.Context,
	req *connect.Request[twitchv1.GetChannelInformationRequest],
) (*connect.Response[twitchv1.GetChannelInformationResponse], error) {
	helixGetChannelInformationRequest := &helix.GetChannelInformationRequest{
		BroadcasterId: req.Msg.BroadcasterId,
	}

	r, err := s.TwitchClient.Helix.GetChannelInformation(helixGetChannelInformationRequest)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&twitchv1.GetChannelInformationResponse{
		BroadcasterId:               r.BroadcasterId,
		BroadcasterLogin:            r.BroadcasterLogin,
		BroadcasterName:             r.BroadcasterName,
		BroadcasterLanguage:         r.BroadcasterLanguage,
		GameId:                      r.GameId,
		GameName:                    r.GameName,
		Title:                       r.Title,
		Delay:                       uint32(r.Delay),
		Tags:                        r.Tags,
		ContentClassificationLabels: r.ContentClassificationLabels,
		IsBrandedContent:            r.IsBrandedContent,
	})

	res.Header().Set("GetChannelInformation-Version", "v1")

	return res, nil
}

func (s *Server) UpdateRedemptionStatus(
	ctx context.Context,
	req *connect.Request[twitchv1.UpdateRedemptionStatusRequest],
) (*connect.Response[twitchv1.UpdateRedemptionStatusResponse], error) {
	helixUpdateRedemptionStatusRequest := &helix.UpdateRedemptionStatusRequest{
		Status: req.Msg.Status,
	}

	params := &helix.UpdateRedemptionStatusParams{
		Id:            req.Msg.RewardId,
		BroadcasterId: req.Msg.BroadcasterId,
		RewardId:      req.Msg.RewardId,
	}

	_, err := s.TwitchClient.Helix.UpdateRedemptionStatus(params, helixUpdateRedemptionStatusRequest)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&twitchv1.UpdateRedemptionStatusResponse{})

	res.Header().Set("UpdateRedemptionStatus-Version", "v1")

	return res, nil
}

func (s *Server) SendShoutout(
	ctx context.Context,
	req *connect.Request[twitchv1.SendShoutoutRequest],
) (*connect.Response[twitchv1.SendShoutoutResponse], error) {
	helixSendShoutoutRequest := &helix.SendShoutoutRequest{
		FromBroadcasterId: req.Msg.FromBroadcasterId,
		ToBroadcasterId:   req.Msg.ToBroadcasterId,
		ModeratorId:       req.Msg.ModeratorId,
	}

	err := s.TwitchClient.Helix.SendChatShoutout(helixSendShoutoutRequest)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&twitchv1.SendShoutoutResponse{})

	res.Header().Set("SendShoutout-Version", "v1")

	return res, nil
}

func (s *Server) SendAnnouncement(
	ctx context.Context,
	req *connect.Request[twitchv1.SendAnnouncementRequest],
) (*connect.Response[twitchv1.SendAnnouncementResponse], error) {
	helixSendAnnouncementRequest := &helix.SendAnnouncementRequest{
		Message: req.Msg.Message,
	}

	if req.Msg.Color != nil {
		helixSendAnnouncementRequest.Color = *req.Msg.Color
	}

	params := &helix.SendAnnouncementParams{
		BroadcasterId: req.Msg.BroadcasterId,
		ModeratorId:   req.Msg.ModeratorId,
	}

	err := s.TwitchClient.Helix.SendChatAnnouncement(params, helixSendAnnouncementRequest)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&twitchv1.SendAnnouncementResponse{})

	res.Header().Set("SendAnnouncement-Version", "v1")

	return res, nil
}

func (s *Server) GetChatters(
	ctx context.Context,
	req *connect.Request[twitchv1.GetChattersRequest],
) (*connect.Response[twitchv1.GetChattersResponse], error) {
	helixGetChattersRequest := &helix.GetChattersParams{
		BroadcasterId: req.Msg.BroadcasterId,
		ModeratorId:   req.Msg.ModeratorId,
		First:         1000,
	}

	chattersResponse, err := s.TwitchClient.Helix.GetChatters(helixGetChattersRequest)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var chatters []*twitchv1.Chatter
	for _, chatterResponse := range chattersResponse {
		chatter := &twitchv1.Chatter{
			UserId:    chatterResponse.UserId,
			UserLogin: chatterResponse.UserLogin,
			Username:  chatterResponse.UserName,
		}

		chatters = append(chatters, chatter)
	}

	res := connect.NewResponse(&twitchv1.GetChattersResponse{Chatters: chatters})

	res.Header().Set("GetChatters-Version", "v1")

	return res, nil
}
