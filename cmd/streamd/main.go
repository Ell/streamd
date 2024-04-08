package main

import (
	"flag"
	"github.com/ell/streamd/server"
	"github.com/ell/streamd/twitch"
	"github.com/ell/streamd/twitch/eventsub"
	"github.com/ell/streamd/twitch/helix"
	"log"
	"os"
)

func main() {
	accessTokenEnv := os.Getenv("TWITCH_ACCESS_TOKEN")
	clientIdEnv := os.Getenv("TWITCH_CLIENT_ID")
	listenAddressEnv := os.Getenv("STREAMD_LISTEN_ADDRESS")
	apiKeyEnv := os.Getenv("STREAMD_API_KEY")

	if listenAddressEnv == "" {
		listenAddressEnv = "localhost:8065"
	}

	accessToken := flag.String("access_token", accessTokenEnv, "twitch access token")
	clientId := flag.String("client_id", clientIdEnv, "twitch client id")
	debugFlag := flag.Bool("debug", false, "run in debug mode")
	listenAddress := flag.String("listen_address", listenAddressEnv, "service listen address")
	apiKey := flag.String("api_key", apiKeyEnv, "service api key")

	if *accessToken == "" {
		log.Fatalf("access_token not provided")
	}

	if *clientId == "" {
		log.Fatalf("client_id not provided")
	}

	if *listenAddress == "" {
		log.Fatalf("listen_address not provided")
	}

	if *apiKey == "" {
		log.Fatalf("api_key not provided")
	}

	flag.Parse()

	var conditions []eventsub.Condition
	var client *twitch.Client
	var err error

	if *debugFlag {
		log.Println("Running in debug mode")
		client, err = CreateDebugClient("ws://127.0.0.1:8080/ws", "http://127.0.0.1:9090")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		client = twitch.NewClient(*clientId, *accessToken)

		user, err := client.Helix.GetUsers(&helix.GetUsersRequest{})
		if err != nil {
			log.Fatalf("could not fetch user for conditions subs %s\n", err)
		}

		conditions = createConditionList(user.Id)
	}

	go client.Listen(conditions...)

	log.Println("Server listening on", *listenAddress)

	s, err := server.NewServer(client, *apiKey)
	if err != nil {
		log.Fatalf("Unable to run server %s\n", err)
	}

	log.Fatal(s.Listen(*listenAddress))
}
