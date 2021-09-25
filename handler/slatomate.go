package handler

import (
	"context"

	"github.com/itzmanish/go-micro/v2"
	"github.com/itzmanish/go-micro/v2/logger"
	"github.com/itzmanish/slatomate/internal/repository"
	slatomatepb "github.com/itzmanish/slatomate/proto/gen/slatomate/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SlatomateHandler interface {
	Health(context.Context, *emptypb.Empty, *slatomatepb.HealthResponse) error
	CreateOrganization(context.Context, *slatomatepb.CreateOrganizationRequest, *slatomatepb.Organization) error
	AuthorizeOrganization(ctx context.Context, in *slatomatepb.AuthorizeOrganizationRequest, out *slatomatepb.GenericResponse) error
	ValidateOrgAccess(ctx context.Context, in *slatomatepb.ValidateOrgAccessRequest, out *slatomatepb.ValidateOrgAccessResponse) error
	DeleteAllOrganization(context.Context, *emptypb.Empty, *emptypb.Empty) error
	GetAllOrganization(context.Context, *slatomatepb.GetAllOrganizationRequest, *slatomatepb.GetAllOrganizationResponse) error
	GetOrganization(context.Context, *slatomatepb.GetOrganizationRequest, *slatomatepb.Organization) error
	DeleteOrganization(context.Context, *slatomatepb.DeleteOrganizationRequest, *emptypb.Empty) error
	CreateUser(context.Context, *slatomatepb.CreateUserRequest, *slatomatepb.User) error
	GetUser(context.Context, *slatomatepb.GetUserRequest, *slatomatepb.User) error
	DeleteUser(context.Context, *slatomatepb.DeleteUserRequest, *emptypb.Empty) error
	UpdateUser(context.Context, *slatomatepb.UpdateUserRequest, *slatomatepb.User) error
	LoginUser(ctx context.Context, in *slatomatepb.User, out *slatomatepb.User) error
	Me(ctx context.Context, in *slatomatepb.APIKeyRequest, out *slatomatepb.User) error
	GenerateAPIKey(context.Context, *slatomatepb.GenerateAPIKeyRequest, *slatomatepb.GenerateAPIKeyResponse) error
	// Admin only
	GetAllUser(context.Context, *emptypb.Empty, *slatomatepb.GetAllUserResponse) error
	// Jobs
	CreateJob(context.Context, *slatomatepb.Job, *slatomatepb.Job) error
	GetJob(context.Context, *slatomatepb.GetJobRequest, *slatomatepb.Job) error
	DeleteJob(context.Context, *slatomatepb.DeleteJobRequest, *emptypb.Empty) error
	GetAllJob(context.Context, *slatomatepb.GetAllJobRequset, *slatomatepb.GetAllJobResponse) error
}

type slatomateHandler struct {
	userRepo  repository.UserRepository
	orgRepo   repository.OrganizationRepository
	jobRepo   repository.JobRepository
	publisher micro.Event
}

func NewHandler(userRepo repository.UserRepository, orgRepo repository.OrganizationRepository, jobRepo repository.JobRepository, event micro.Event) SlatomateHandler {
	return &slatomateHandler{userRepo, orgRepo, jobRepo, event}
}

func (s *slatomateHandler) Health(ctx context.Context, in *emptypb.Empty, out *slatomatepb.HealthResponse) error {
	logger.Debug("Health request")
	out.Status = "up"
	out.Version = "v1.0.1"
	return nil
}
