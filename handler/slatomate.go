package handler

import (
	"context"

	"github.com/itzmanish/slatomate/internal/repository"
	slatomatepb "github.com/itzmanish/slatomate/proto/slatomate"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SlatomateHandler interface {
	CreateOrganization(context.Context, *slatomatepb.CreateOrganizationRequest, *slatomatepb.Organization) error
	GetAllOrganization(context.Context, *slatomatepb.GetAllOrganizationRequest, *slatomatepb.GetAllOrganizationResponse) error
	GetOrganization(context.Context, *slatomatepb.GetOrganizationRequest, *slatomatepb.Organization) error
	DeleteOrganization(context.Context, *slatomatepb.DeleteOrganizationRequest, *emptypb.Empty) error
	CreateUser(context.Context, *slatomatepb.CreateUserRequest, *slatomatepb.User) error
	GetUser(context.Context, *slatomatepb.GetUserRequest, *slatomatepb.User) error
	DeleteUser(context.Context, *slatomatepb.DeleteUserRequest, *emptypb.Empty) error
	UpdateUser(context.Context, *slatomatepb.UpdateUserRequest, *slatomatepb.User) error
	GenerateAPIKey(context.Context, *slatomatepb.GenerateAPIKeyRequest, *slatomatepb.GenerateAPIKeyResponse) error
	// Admin only
	GetAllUser(context.Context, *emptypb.Empty, *slatomatepb.GetAllUserResponse) error
}

type slatomateHandler struct {
	userRepo repository.UserRepository
	orgRepo  repository.OrganizationRepository
}

func NewHandler(userRepo repository.UserRepository, orgRepo repository.OrganizationRepository) SlatomateHandler {
	return &slatomateHandler{userRepo, orgRepo}
}
