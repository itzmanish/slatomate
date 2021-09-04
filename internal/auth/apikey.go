package auth

import (
	"context"

	"github.com/itzmanish/go-micro/v2/auth"
	"github.com/itzmanish/slatomate/internal/entity"
	"github.com/itzmanish/slatomate/internal/repository"
)

var Scheme = "APIKEY "

type apiKeyAuth struct {
	userRepo repository.UserRepository
}

type Auth interface {
	Generate(id string, opts ...auth.GenerateOption) (*auth.Account, error)
	Inspect(token string) (*entity.User, error)
	String() string
}

func NewAPIKeyAuth(userRepo repository.UserRepository) Auth {
	return &apiKeyAuth{
		userRepo: userRepo,
	}
}

func (apiAuth *apiKeyAuth) Generate(id string, opts ...auth.GenerateOption) (*auth.Account, error) {
	return nil, nil
}

func (apiAuth *apiKeyAuth) Inspect(token string) (*entity.User, error) {
	return apiAuth.userRepo.GetUser(&entity.User{APIKey: token})
}

func (apiAuth *apiKeyAuth) String() string {
	return "API Key Auth"
}

type accountKey struct{}

// AccountFromContext gets the account from the context, which
// is set by the auth wrapper at the start of a call. If the account
// is not set, a nil account will be returned. The error is only returned
// when there was a problem retrieving an account
func AccountFromContext(ctx context.Context) (*entity.User, bool) {
	acc, ok := ctx.Value(accountKey{}).(*entity.User)
	return acc, ok
}

// ContextWithAccount sets the account in the context
func ContextWithAccount(ctx context.Context, account *entity.User) context.Context {
	return context.WithValue(ctx, accountKey{}, account)
}
