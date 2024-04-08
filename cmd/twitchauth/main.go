package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var listenAddress = "localhost:5790"
var scopes = []string{
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
}

func main() {
	clientIdEnv := os.Getenv("TWITCH_CLIENT_ID")
	clientId := flag.String("client_id", clientIdEnv, "twitch client id")

	flag.Parse()

	if clientId == nil || *clientId == "" {
		panic("Missing TWITCH_CLIENT_ID / client_id argument")
	}

	authUrl := buildUrl(*clientId)

	println(authUrl)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := `
			<!DOCTYPE html>
			<body>
				<div id="key"></div>
				<script>
					const hashValue = document.location.hash.split("&")[0].split("=")[1];
					document.getElementById("key").innerHTML = "<h1> Access Token: " + hashValue + "</h1>";
				</script>
			</body>
		`

		_, err := w.Write([]byte(response))

		if err != nil {
			panic(err)
		}
	})

	fmt.Printf("Webserver running on http://%s\n", listenAddress)

	log.Fatal(http.ListenAndServe(listenAddress, nil))
}

func buildUrl(clientId string) string {
	redirectUri := "http://" + listenAddress + "/"

	queryValues := url.Values{}
	queryValues.Add("response_type", "token")
	queryValues.Add("client_id", clientId)
	queryValues.Add("redirect_uri", redirectUri)
	queryValues.Add("scope", strings.Join(scopes, " "))
	queryValues.Add("token_type", "bearer")

	queryParams := queryValues.Encode()

	return "https://id.twitch.tv/oauth2/authorize?" + queryParams
}
