package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/itzmanish/go-micro/v2/errors"
	"github.com/itzmanish/go-micro/v2/logger"
	"github.com/itzmanish/slatomate/internal/auth"
	"github.com/itzmanish/slatomate/internal/entity"
	slatomatepb "github.com/itzmanish/slatomate/proto/gen/slatomate/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *slatomateHandler) CreateUser(ctx context.Context, in *slatomatepb.CreateUserRequest, out *slatomatepb.User) error {

	if len(in.GetEmail()) == 0 {
		return errors.BadRequest("CREATE_USER_HANDLER", "Email is required but not provided.")
	}
	if len(in.GetPassword()) == 0 {
		return errors.BadRequest("CREATE_USER_HANDLER", "Password is required but not provided.")
	}

	user := &entity.User{
		Name:  in.Name,
		Email: in.Email,
	}

	err := user.SetPassword(in.GetPassword())
	if err != nil {
		return err
	}
	err = user.GenerateAPIKey()
	if err != nil {
		return err
	}

	resUser, err := h.userRepo.CreateUser(user)
	if err != nil {
		return err
	}
	*out = entity.DeserializeUser(resUser)
	return nil
}

func (h *slatomateHandler) GetUser(ctx context.Context, in *slatomatepb.GetUserRequest, out *slatomatepb.User) error {
	if len(in.GetId()) == 0 && len(in.GetEmail()) == 0 && len(in.GetApiKey()) == 0 {
		return errors.BadRequest("GET_USER_HANDLER", "user id or email is required.")
	}
	filter := &entity.User{}
	if len(in.GetId()) != 0 {
		id, err := uuid.Parse(in.GetId())
		if err != nil {
			return err
		}
		filter.ID = id
	} else if len(in.GetApiKey()) != 0 {
		filter.APIKey = in.ApiKey
	} else {
		filter.Email = in.GetEmail()
	}
	userOut, err := h.userRepo.GetUser(filter)
	if err != nil {
		return err
	}

	*out = entity.DeserializeUser(userOut)
	return nil
}

func (h *slatomateHandler) DeleteUser(ctx context.Context, in *slatomatepb.DeleteUserRequest, out *emptypb.Empty) error {
	if len(in.GetId()) == 0 {
		return errors.BadRequest("DELETE_USER_HANDLER", "user id is required!")
	}
	id, err := uuid.Parse(in.GetId())
	if err != nil {
		return err
	}
	return h.userRepo.DeleteUser(&entity.User{ID: id})
}

func (h *slatomateHandler) UpdateUser(ctx context.Context, in *slatomatepb.UpdateUserRequest, out *slatomatepb.User) error {
	user, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("DELETE_USER_HANDLER", "Unauthorized access.")
	}
	updatedUser, err := h.userRepo.UpdateUser(&entity.User{ID: user.ID, Name: in.GetName()})
	if err != nil {
		return err
	}
	*out = entity.DeserializeUser(updatedUser)
	return nil
}

func (h *slatomateHandler) GetAllUser(ctx context.Context, in *emptypb.Empty, out *slatomatepb.GetAllUserResponse) error {
	users, err := h.userRepo.GetAllUser()
	if err != nil {
		return err
	}

	out.Users = make([]*slatomatepb.User, len(users))
	out.Count = int32(len(users))
	for i, user := range users {
		puser := entity.DeserializeUser(user)
		out.Users[i] = &puser
	}
	return nil
}

func (h *slatomateHandler) LoginUser(ctx context.Context, in *slatomatepb.User, out *slatomatepb.User) error {
	logger.Debug("LoginUser request: ", in)
	if len(in.Email) == 0 || len(in.Password) == 0 {
		return errors.BadRequest("LoginUser", "Email/Password both are required!")
	}
	user, err := h.userRepo.GetUser(&entity.User{Email: in.Email})
	if err != nil {
		return err
	}
	ok, err := user.ValidatePassword(in.Password)
	if err != nil {
		return err
	}
	if !ok {
		return errors.Unauthorized("LoginUser", "Email/password is not correct.")
	}
	*out = entity.DeserializeUser(user)
	return nil
}

func (h *slatomateHandler) Me(ctx context.Context, in *slatomatepb.APIKeyRequest, out *slatomatepb.User) error {
	logger.Debug("ValidateAPIKey request: ", in)
	if len(in.GetApiKey()) == 0 {
		return errors.BadRequest("ValidateAPIKey", "API Key is required!")
	}
	user, err := h.userRepo.GetUser(&entity.User{APIKey: in.ApiKey})
	if err != nil {
		return err
	}
	*out = entity.DeserializeUser(user)
	return nil
}

// Not for now
func (h *slatomateHandler) GenerateAPIKey(ctx context.Context, in *slatomatepb.GenerateAPIKeyRequest, out *slatomatepb.GenerateAPIKeyResponse) error {
	return nil
}
