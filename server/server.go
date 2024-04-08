package server

import (
	"connectrpc.com/connect"
	"github.com/ell/streamd/rpc/twitch/v1/twitchv1connect"
	"github.com/ell/streamd/twitch"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
	"sync"
)

type Server struct {
	TwitchClient        *twitch.Client
	mux                 *http.ServeMux
	eventListenersMutex sync.Mutex
	apiKey              string
}

func NewServer(twitchClient *twitch.Client, apiKey string) (*Server, error) {
	interceptors := connect.WithInterceptors(NewAuthInterceptor(apiKey))

	mux := http.NewServeMux()

	server := &Server{
		TwitchClient: twitchClient,
		mux:          mux,
		apiKey:       apiKey,
	}

	mux.Handle(twitchv1connect.NewTwitchServiceHandler(server, interceptors))

	return server, nil
}

func (s *Server) Listen(addr string) error {
	return http.ListenAndServe(addr, h2c.NewHandler(s.mux, &http2.Server{}))
}
