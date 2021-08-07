package handler

import (
	"context"

	slatomatepb "github.com/itzmanish/slatomate/proto/slatomate"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SlatomateHandler interface {
	CreateProject(context.Context, *slatomatepb.CreateProjectRequest, *slatomatepb.Project) error
	GetAllProjct(context.Context, *slatomatepb.GetAllProjectRequest, *slatomatepb.GetAllProjectResponse) error
	GetProject(context.Context, *slatomatepb.GetProjectRequest, *slatomatepb.Project) error
	DeleteProject(context.Context, *slatomatepb.DeleteProjectRequest, *emptypb.Empty) error
	CreateUser(context.Context, *slatomatepb.CreateUserRequest, *slatomatepb.User) error
	GetUser(context.Context, *slatomatepb.GetUserRequest, *slatomatepb.User) error
	DeleteUser(context.Context, *slatomatepb.DeleteUserRequest, *emptypb.Empty) error
	// Admin only
	GetAllUser(context.Context, *emptypb.Empty, *slatomatepb.GetAllUserResponse) error
}

type slatomateHandler struct{}

func NewHandler() SlatomateHandler {
	return &slatomateHandler{}
}

func (h *slatomateHandler) CreateProject(ctx context.Context, in *slatomatepb.CreateProjectRequest, out *slatomatepb.Project) error {
	return nil
}

func (h *slatomateHandler) GetAllProjct(ctx context.Context, in *slatomatepb.GetAllProjectRequest, out *slatomatepb.GetAllProjectResponse) error {
	return nil
}

func (h *slatomateHandler) GetProject(ctx context.Context, in *slatomatepb.GetProjectRequest, out *slatomatepb.Project) error {
	return nil
}

func (h *slatomateHandler) DeleteProject(ctx context.Context, in *slatomatepb.DeleteProjectRequest, out *emptypb.Empty) error {
	return nil
}

func (h *slatomateHandler) CreateUser(ctx context.Context, in *slatomatepb.CreateUserRequest, out *slatomatepb.User) error {
	return nil
}

func (h *slatomateHandler) GetUser(ctx context.Context, in *slatomatepb.GetUserRequest, out *slatomatepb.User) error {
	return nil
}

func (h *slatomateHandler) DeleteUser(ctx context.Context, in *slatomatepb.DeleteUserRequest, out *emptypb.Empty) error {
	return nil
}

func (h *slatomateHandler) GetAllUser(ctx context.Context, in *emptypb.Empty, out *slatomatepb.GetAllUserResponse) error {
	return nil
}
