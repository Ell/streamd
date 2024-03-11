package main

import (
	"connectrpc.com/connect"
	"context"
	twitchrpc "github.com/ell/streamd/rpc/twitch/v1"
	twitchrpcconnect "github.com/ell/streamd/rpc/twitch/v1/twitchv1connect"
	"log"
	"net/http"
)

func main() {
	client := twitchrpcconnect.NewTwitchServiceClient(http.DefaultClient, "http://localhost:9090")

	res, err := client.GetCurrentUserInfo(context.Background(), connect.NewRequest(&twitchrpc.GetCurrentUserInfoRequest{}))
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res.Msg)
}
