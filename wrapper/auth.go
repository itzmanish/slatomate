package wrapper

import (
	"context"
	"strings"

	"github.com/itzmanish/go-micro/v2/errors"
	"github.com/itzmanish/go-micro/v2/metadata"
	"github.com/itzmanish/go-micro/v2/server"
	"github.com/itzmanish/slatomate/internal/auth"
)

type Array []string

var NoAuthEndpoint Array = []string{
	"Slatomate.Login",
	"Slatomate.CreateUser",
	"Slatomate.LoginUser",
}

func (arr Array) Has(value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

// AuthHandler wraps a server handler to perform auth
func AuthHandler(a auth.Auth) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {

			// Check for debug endpoints which should be excluded from auth
			if NoAuthEndpoint.Has(req.Endpoint()) {
				return h(ctx, req, rsp)
			}

			// Extract the token if present. Note: if noop is being used
			// then the token can be blank without erroring
			header, ok := metadata.Get(ctx, "Authorization")
			if !ok {
				errors.Unauthorized("NO_APIKEY", "API key is required!")
			}
			// Ensure the correct scheme is being used
			if !strings.HasPrefix(header, auth.Scheme) {
				return errors.Unauthorized(req.Service(), "invalid authorization header. expected APIKEY schema")
			}

			// Strip the prefix and inspect the resulting token
			account, err := a.Inspect(strings.TrimPrefix(header, auth.Scheme))
			if err != nil {
				return errors.Unauthorized("INVALID_APIKEY", err.Error())
			}

			// // There is an account, set it in the context
			if account != nil {
				ctx = auth.ContextWithAccount(ctx, account)
			}

			// The user is authorised, allow the call
			return h(ctx, req, rsp)
		}
	}
}
