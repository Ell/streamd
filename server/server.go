package server

import (
	"connectrpc.com/connect"
	"github.com/ell/streamd/rpc/twitch/v1/twitchv1connect"
	"github.com/ell/streamd/twitch"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	TwitchClient        *twitch.Client
	mux                 *http.ServeMux
	eventListenersMutex sync.Mutex
	apiKey              string
}

func newCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOriginFunc: func(origin string) bool {
			// Allow all origins, which effectively disables CORS.
			return true
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
		},
		// Let browsers cache CORS information for longer, which reduces the number
		// of preflight requests. Any changes to ExposedHeaders won't take effect
		// until the cached data expires. FF caps this value at 24h, and modern
		// Chrome caps it at 2h.
		MaxAge: int(2 * time.Hour / time.Second),
	})
}

func NewServer(twitchClient *twitch.Client, apiKey string) (*Server, error) {
	server := &Server{
		TwitchClient: twitchClient,
		apiKey:       apiKey,
	}

	return server, nil
}

func (s *Server) Listen(addr string) error {
	interceptors := connect.WithInterceptors(NewAuthInterceptor(s.apiKey))

	mux := http.NewServeMux()
	mux.Handle(twitchv1connect.NewTwitchServiceHandler(s, interceptors))

	srv := &http.Server{
		Addr: addr,
		Handler: h2c.NewHandler(
			newCORS().Handler(mux),
			&http2.Server{},
		),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       0,
		WriteTimeout:      0,
		MaxHeaderBytes:    8 * 1024,
	}

	return srv.ListenAndServe()
}
