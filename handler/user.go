package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/itzmanish/go-micro/v2/errors"
	"github.com/itzmanish/slatomate/internal/entity"
	slatomatepb "github.com/itzmanish/slatomate/proto/slatomate"
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

	user.SetPassword(in.GetPassword())

	resUser, err := h.userRepo.CreateUser(user)
	if err != nil {
		return err
	}
	*out = entity.DeserializeUser(resUser)
	return nil
}

func (h *slatomateHandler) GetUser(ctx context.Context, in *slatomatepb.GetUserRequest, out *slatomatepb.User) error {
	if len(in.GetId()) == 0 && len(in.GetEmail()) == 0 {
		return errors.BadRequest("GET_USER_HANDLER", "user id or email is required.")
	}
	filter := &entity.User{}
	if len(in.GetId()) != 0 {
		id, err := uuid.Parse(in.GetId())
		if err != nil {
			return err
		}
		filter.ID = id
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
