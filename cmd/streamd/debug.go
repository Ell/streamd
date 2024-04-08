package main

import (
	"encoding/json"
	"github.com/ell/streamd/twitch"
	"github.com/ell/streamd/twitch/eventsub"
	"github.com/ell/streamd/twitch/helix"
	"github.com/google/go-querystring/query"
	"io"
	"log"
	"net/http"
	"time"
)

type accessTokenResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int      `json:"expires_in"`
	Scope        []string `json:"scope"`
	TokenType    string   `json:"token_type"`
}

type accessTokenRequest struct {
	ClientId     string   `url:"client_id"`
	ClientSecret string   `url:"client_secret"`
	GrantType    string   `url:"grant_type"`
	UserId       string   `url:"user_id"`
	Scope        []string `url:"scope"`
}

type unitResponse[D interface{}] struct {
	Cursor string `json:"cursor"`
	Total  int    `json:"total"`
	Data   []D    `json:"data"`
}

type unitClient struct {
	Id          string `json:"ID"`
	Secret      string `json:"Secret"`
	Name        string `json:"Name"`
	IsExtension bool   `json:"IsExtension"`
}

type unitUser struct {
	Id              string    `json:"id"`
	Login           string    `json:"login"`
	DisplayName     string    `json:"display_name"`
	Email           string    `json:"email"`
	Type            string    `json:"type"`
	BroadcasterType string    `json:"broadcaster_type"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	ProfileImageUrl string    `json:"profile_image_url"`
	OfflineImageUrl string    `json:"offline_image_url"`
	ViewCount       int       `json:"view_count"`
	GameId          struct {
		String string `json:"String"`
		Valid  bool   `json:"Valid"`
	} `json:"game_id"`
	GameName struct {
		String string `json:"String"`
		Valid  bool   `json:"Valid"`
	} `json:"game_name"`
	Title            string `json:"title"`
	StreamLanguage   string `json:"stream_language"`
	Delay            int    `json:"delay"`
	IsBrandedContent bool   `json:"is_branded_content"`
}

func getUnitResponse[D interface{}](baseUrl, endpoint string, r *unitResponse[D]) error {
	c := http.Client{Timeout: 5 * time.Second}

	url := baseUrl + "/units/" + endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("unable to make unit request")
		return err
	}

	res, err := c.Do(req)
	if err != nil {
		log.Println("unable to do unit request", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("unable to close unit response body")
		}
	}(res.Body)

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("unable to unmarshal unit response", err)
		return err
	}

	return json.Unmarshal(respBody, &r)
}

func getAccessToken(baseUrl, clientId, clientSecret, userId string) (string, error) {
	params := &accessTokenRequest{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		GrantType:    "user_token",
		UserId:       userId,
		Scope: []string{
			"bits:read",
			"channel:read:ads",
			"channel:edit:commercial",
			"channel:manage:broadcast",
			"channel:read:hype_train",
			"channel:manage:moderators",
			"channel:read:polls",
			"channel:manage:polls",
			"channel:read:predictions",
			"channel:manage:predictions",
			"channel:manage:raids",
			"channel:read:redemptions",
			"channel:manage:redemptions",
			"channel:read:subscriptions",
			"channel:read:vips",
			"channel:manage:vips",
			"clips:edit",
			"moderation:read",
			"moderator:manage:announcements",
			"moderator:manage:automod",
			"moderator:manage:banned_users",
			"moderator:read:blocked_terms",
			"moderator:manage:blocked_terms",
			"moderator:manage:chat_messages",
			"moderator:read:chat_settings",
			"moderator:manage:chat_settings",
			"moderator:read:chatters",
			"moderator:read:followers",
			"moderator:read:shoutouts",
			"moderator:manage:shoutouts",
			"moderator:read:unban_requests",
			"moderator:manage:unban_requests",
			"channel:bot",
			"channel:moderate",
			"chat:edit",
			"chat:read",
			"user:bot",
			"user:read:chat",
			"user:write:chat",
		},
	}

	resp := &accessTokenResponse{}

	v, err := query.Values(params)
	if err != nil {
		log.Println("unable to serialize access token values")
		return "", err
	}

	url := baseUrl + "/auth/authorize?" + v.Encode()
	log.Println("access token url", url)

	c := http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Println("unable to create access token request")
		return "", err
	}

	res, err := c.Do(req)
	if err != nil {
		log.Println("unable to do access token request")
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Unable to close helix body reader %s", err)
		}
	}(res.Body)

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("unable to read access token body")
		return "", err
	}

	err = json.Unmarshal(respBody, resp)
	if err != nil {
		log.Println("unable to unmarshal access token response")
		return "", err
	}

	return resp.AccessToken, nil
}

func CreateDebugClient(eventsubUrl, helixUrl string) (*twitch.Client, error) {
	users := &unitResponse[unitUser]{}

	err := getUnitResponse(helixUrl, "users", users)
	if err != nil {
		return nil, err
	}

	clients := &unitResponse[unitClient]{}
	err = getUnitResponse(helixUrl, "clients", clients)
	if err != nil {
		return nil, err
	}

	user := users.Data[0]
	client := clients.Data[0]

	accessToken, err := getAccessToken(helixUrl, client.Id, client.Secret, user.Id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("got access token", accessToken)

	helixClient := helix.NewTestClient(client.Id, accessToken, helixUrl+"/mock")
	eventsubClient := eventsub.NewTestClient(eventsubUrl)

	twitchClient := &twitch.Client{
		Helix:       *helixClient,
		Eventsub:    eventsubClient,
		ClientId:    client.Id,
		AccessToken: accessToken,
	}

	log.Println("created new debug client")

	return twitchClient, nil
}
