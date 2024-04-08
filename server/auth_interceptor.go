package server

import (
	"connectrpc.com/connect"
	"context"
	"errors"
)

const tokenHeader = "Streamd-Auth"

var noTokenError = errors.New("no token provided")
var wrongTokenError = errors.New("wrong token provided")

type AuthInterceptor struct {
	authToken string
}

func NewAuthInterceptor(authToken string) *AuthInterceptor {
	return &AuthInterceptor{
		authToken: authToken,
	}
}

func (i *AuthInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(
		ctx context.Context,
		req connect.AnyRequest,
	) (connect.AnyResponse, error) {
		if req.Spec().IsClient {
			req.Header().Set(tokenHeader, i.authToken)
		} else if req.Header().Get(tokenHeader) == "" {
			return nil, connect.NewError(connect.CodeUnauthenticated, noTokenError)
		} else if req.Header().Get(tokenHeader) != i.authToken {
			return nil, connect.NewError(connect.CodeUnauthenticated, wrongTokenError)
		}

		return next(ctx, req)
	}
}

func (i *AuthInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		conn := next(ctx, spec)
		conn.RequestHeader().Set(tokenHeader, i.authToken)

		return conn
	}
}

func (i *AuthInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(ctx context.Context, conn connect.StreamingHandlerConn) error {
		if conn.RequestHeader().Get(tokenHeader) == "" {
			return connect.NewError(connect.CodeUnauthenticated, noTokenError)
		} else if conn.RequestHeader().Get(tokenHeader) != i.authToken {
			return connect.NewError(connect.CodeUnauthenticated, wrongTokenError)
		}

		return next(ctx, conn)
	}
}
