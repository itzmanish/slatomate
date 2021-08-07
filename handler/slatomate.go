package handler

import (
	"context"

	"github.com/itzmanish/slatomate/internal/repository"
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

type slatomateHandler struct {
	userRepo    repository.UserRepository
	projectRepo repository.ProjectRepository
}

func NewHandler(userRepo repository.UserRepository, projectRepo repository.ProjectRepository) SlatomateHandler {
	return &slatomateHandler{userRepo, projectRepo}
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
