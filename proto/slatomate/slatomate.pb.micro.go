// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/slatomate/slatomate.proto

package slatomate

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

import (
	context "context"
	api "github.com/itzmanish/go-micro/v2/api"
	client "github.com/itzmanish/go-micro/v2/client"
	server "github.com/itzmanish/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Slatomate service

func NewSlatomateEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "Slatomate.CreateOrganization",
			Path:    []string{"/v1/slatomate/org"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.AuthorizeOrganization",
			Path:    []string{"/v1/slatomate/org/authorize"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.ValidateOrgAccess",
			Path:    []string{"/v1/slatomate/org/{org_id}"},
			Method:  []string{"POST"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.GetAllOrganization",
			Path:    []string{"/v1/slatomate/orgs"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.GetOrganization",
			Path:    []string{"/v1/slatomate/org/{id}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.DeleteOrganization",
			Path:    []string{"/v1/slatomate/org/{id}"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.DeleteAllOrganization",
			Path:    []string{"/v1/slatomate/orgs"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.CreateUser",
			Path:    []string{"/v1/slatomate/user"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.GetUser",
			Path:    []string{"/v1/slatomate/user/{id}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.DeleteUser",
			Path:    []string{"/v1/slatomate/user/{id}"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.UpdateUser",
			Path:    []string{"/v1/slatomate/user/{id}"},
			Method:  []string{"PATCH"},
			Body:    "name",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.LoginUser",
			Path:    []string{"/v1/slatomate/login"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.Me",
			Path:    []string{"/v1/slatomate/user/me"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.GenerateAPIKey",
			Path:    []string{"/v1/slatomate/user/api_key"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.CreateJob",
			Path:    []string{"/v1/slatomate/org/{org_id}/job"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.GetJob",
			Path:    []string{"/v1/slatomate/org/{org_id}/job/{id}{name}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.DeleteJob",
			Path:    []string{"/v1/slatomate/org/{org_id}/job/{id}"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Slatomate.GetAllJob",
			Path:    []string{"/v1/slatomate/org/{org_id}/job/{id}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
	}
}

// Client API for Slatomate service

type SlatomateService interface {
	CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...client.CallOption) (*Organization, error)
	AuthorizeOrganization(ctx context.Context, in *AuthorizeOrganizationRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	ValidateOrgAccess(ctx context.Context, in *ValidateOrgAccessRequest, opts ...client.CallOption) (*ValidateOrgAccessResponse, error)
	GetAllOrganization(ctx context.Context, in *GetAllOrganizationRequest, opts ...client.CallOption) (*GetAllOrganizationResponse, error)
	GetOrganization(ctx context.Context, in *GetOrganizationRequest, opts ...client.CallOption) (*Organization, error)
	DeleteOrganization(ctx context.Context, in *DeleteOrganizationRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	DeleteAllOrganization(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*emptypb.Empty, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*User, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*User, error)
	LoginUser(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
	Me(ctx context.Context, in *APIKeyRequest, opts ...client.CallOption) (*User, error)
	// Not for now
	GenerateAPIKey(ctx context.Context, in *GenerateAPIKeyRequest, opts ...client.CallOption) (*GenerateAPIKeyResponse, error)
	// Admin only
	GetAllUser(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetAllUserResponse, error)
	// Jobs
	CreateJob(ctx context.Context, in *Job, opts ...client.CallOption) (*Job, error)
	GetJob(ctx context.Context, in *GetJobRequest, opts ...client.CallOption) (*Job, error)
	DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	GetAllJob(ctx context.Context, in *GetAllJobRequset, opts ...client.CallOption) (*GetAllJobResponse, error)
}

type slatomateService struct {
	c    client.Client
	name string
}

func NewSlatomateService(name string, c client.Client) SlatomateService {
	return &slatomateService{
		c:    c,
		name: name,
	}
}

func (c *slatomateService) CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...client.CallOption) (*Organization, error) {
	req := c.c.NewRequest(c.name, "Slatomate.CreateOrganization", in)
	out := new(Organization)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) AuthorizeOrganization(ctx context.Context, in *AuthorizeOrganizationRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Slatomate.AuthorizeOrganization", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) ValidateOrgAccess(ctx context.Context, in *ValidateOrgAccessRequest, opts ...client.CallOption) (*ValidateOrgAccessResponse, error) {
	req := c.c.NewRequest(c.name, "Slatomate.ValidateOrgAccess", in)
	out := new(ValidateOrgAccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) GetAllOrganization(ctx context.Context, in *GetAllOrganizationRequest, opts ...client.CallOption) (*GetAllOrganizationResponse, error) {
	req := c.c.NewRequest(c.name, "Slatomate.GetAllOrganization", in)
	out := new(GetAllOrganizationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) GetOrganization(ctx context.Context, in *GetOrganizationRequest, opts ...client.CallOption) (*Organization, error) {
	req := c.c.NewRequest(c.name, "Slatomate.GetOrganization", in)
	out := new(Organization)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) DeleteOrganization(ctx context.Context, in *DeleteOrganizationRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Slatomate.DeleteOrganization", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) DeleteAllOrganization(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Slatomate.DeleteAllOrganization", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Slatomate.CreateUser", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Slatomate.GetUser", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Slatomate.DeleteUser", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Slatomate.UpdateUser", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) LoginUser(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Slatomate.LoginUser", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) Me(ctx context.Context, in *APIKeyRequest, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Slatomate.Me", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) GenerateAPIKey(ctx context.Context, in *GenerateAPIKeyRequest, opts ...client.CallOption) (*GenerateAPIKeyResponse, error) {
	req := c.c.NewRequest(c.name, "Slatomate.GenerateAPIKey", in)
	out := new(GenerateAPIKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) GetAllUser(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*GetAllUserResponse, error) {
	req := c.c.NewRequest(c.name, "Slatomate.GetAllUser", in)
	out := new(GetAllUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) CreateJob(ctx context.Context, in *Job, opts ...client.CallOption) (*Job, error) {
	req := c.c.NewRequest(c.name, "Slatomate.CreateJob", in)
	out := new(Job)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) GetJob(ctx context.Context, in *GetJobRequest, opts ...client.CallOption) (*Job, error) {
	req := c.c.NewRequest(c.name, "Slatomate.GetJob", in)
	out := new(Job)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Slatomate.DeleteJob", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) GetAllJob(ctx context.Context, in *GetAllJobRequset, opts ...client.CallOption) (*GetAllJobResponse, error) {
	req := c.c.NewRequest(c.name, "Slatomate.GetAllJob", in)
	out := new(GetAllJobResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Slatomate service

type SlatomateHandler interface {
	CreateOrganization(context.Context, *CreateOrganizationRequest, *Organization) error
	AuthorizeOrganization(context.Context, *AuthorizeOrganizationRequest, *emptypb.Empty) error
	ValidateOrgAccess(context.Context, *ValidateOrgAccessRequest, *ValidateOrgAccessResponse) error
	GetAllOrganization(context.Context, *GetAllOrganizationRequest, *GetAllOrganizationResponse) error
	GetOrganization(context.Context, *GetOrganizationRequest, *Organization) error
	DeleteOrganization(context.Context, *DeleteOrganizationRequest, *emptypb.Empty) error
	DeleteAllOrganization(context.Context, *emptypb.Empty, *emptypb.Empty) error
	CreateUser(context.Context, *CreateUserRequest, *User) error
	GetUser(context.Context, *GetUserRequest, *User) error
	DeleteUser(context.Context, *DeleteUserRequest, *emptypb.Empty) error
	UpdateUser(context.Context, *UpdateUserRequest, *User) error
	LoginUser(context.Context, *User, *User) error
	Me(context.Context, *APIKeyRequest, *User) error
	// Not for now
	GenerateAPIKey(context.Context, *GenerateAPIKeyRequest, *GenerateAPIKeyResponse) error
	// Admin only
	GetAllUser(context.Context, *emptypb.Empty, *GetAllUserResponse) error
	// Jobs
	CreateJob(context.Context, *Job, *Job) error
	GetJob(context.Context, *GetJobRequest, *Job) error
	DeleteJob(context.Context, *DeleteJobRequest, *emptypb.Empty) error
	GetAllJob(context.Context, *GetAllJobRequset, *GetAllJobResponse) error
}

func RegisterSlatomateHandler(s server.Server, hdlr SlatomateHandler, opts ...server.HandlerOption) error {
	type slatomate interface {
		CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, out *Organization) error
		AuthorizeOrganization(ctx context.Context, in *AuthorizeOrganizationRequest, out *emptypb.Empty) error
		ValidateOrgAccess(ctx context.Context, in *ValidateOrgAccessRequest, out *ValidateOrgAccessResponse) error
		GetAllOrganization(ctx context.Context, in *GetAllOrganizationRequest, out *GetAllOrganizationResponse) error
		GetOrganization(ctx context.Context, in *GetOrganizationRequest, out *Organization) error
		DeleteOrganization(ctx context.Context, in *DeleteOrganizationRequest, out *emptypb.Empty) error
		DeleteAllOrganization(ctx context.Context, in *emptypb.Empty, out *emptypb.Empty) error
		CreateUser(ctx context.Context, in *CreateUserRequest, out *User) error
		GetUser(ctx context.Context, in *GetUserRequest, out *User) error
		DeleteUser(ctx context.Context, in *DeleteUserRequest, out *emptypb.Empty) error
		UpdateUser(ctx context.Context, in *UpdateUserRequest, out *User) error
		LoginUser(ctx context.Context, in *User, out *User) error
		Me(ctx context.Context, in *APIKeyRequest, out *User) error
		GenerateAPIKey(ctx context.Context, in *GenerateAPIKeyRequest, out *GenerateAPIKeyResponse) error
		GetAllUser(ctx context.Context, in *emptypb.Empty, out *GetAllUserResponse) error
		CreateJob(ctx context.Context, in *Job, out *Job) error
		GetJob(ctx context.Context, in *GetJobRequest, out *Job) error
		DeleteJob(ctx context.Context, in *DeleteJobRequest, out *emptypb.Empty) error
		GetAllJob(ctx context.Context, in *GetAllJobRequset, out *GetAllJobResponse) error
	}
	type Slatomate struct {
		slatomate
	}
	h := &slatomateHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.CreateOrganization",
		Path:    []string{"/v1/slatomate/org"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.AuthorizeOrganization",
		Path:    []string{"/v1/slatomate/org/authorize"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.ValidateOrgAccess",
		Path:    []string{"/v1/slatomate/org/{org_id}"},
		Method:  []string{"POST"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.GetAllOrganization",
		Path:    []string{"/v1/slatomate/orgs"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.GetOrganization",
		Path:    []string{"/v1/slatomate/org/{id}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.DeleteOrganization",
		Path:    []string{"/v1/slatomate/org/{id}"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.DeleteAllOrganization",
		Path:    []string{"/v1/slatomate/orgs"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.CreateUser",
		Path:    []string{"/v1/slatomate/user"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.GetUser",
		Path:    []string{"/v1/slatomate/user/{id}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.DeleteUser",
		Path:    []string{"/v1/slatomate/user/{id}"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.UpdateUser",
		Path:    []string{"/v1/slatomate/user/{id}"},
		Method:  []string{"PATCH"},
		Body:    "name",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.LoginUser",
		Path:    []string{"/v1/slatomate/login"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.Me",
		Path:    []string{"/v1/slatomate/user/me"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.GenerateAPIKey",
		Path:    []string{"/v1/slatomate/user/api_key"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.CreateJob",
		Path:    []string{"/v1/slatomate/org/{org_id}/job"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.GetJob",
		Path:    []string{"/v1/slatomate/org/{org_id}/job/{id}{name}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.DeleteJob",
		Path:    []string{"/v1/slatomate/org/{org_id}/job/{id}"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Slatomate.GetAllJob",
		Path:    []string{"/v1/slatomate/org/{org_id}/job/{id}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&Slatomate{h}, opts...))
}

type slatomateHandler struct {
	SlatomateHandler
}

func (h *slatomateHandler) CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, out *Organization) error {
	return h.SlatomateHandler.CreateOrganization(ctx, in, out)
}

func (h *slatomateHandler) AuthorizeOrganization(ctx context.Context, in *AuthorizeOrganizationRequest, out *emptypb.Empty) error {
	return h.SlatomateHandler.AuthorizeOrganization(ctx, in, out)
}

func (h *slatomateHandler) ValidateOrgAccess(ctx context.Context, in *ValidateOrgAccessRequest, out *ValidateOrgAccessResponse) error {
	return h.SlatomateHandler.ValidateOrgAccess(ctx, in, out)
}

func (h *slatomateHandler) GetAllOrganization(ctx context.Context, in *GetAllOrganizationRequest, out *GetAllOrganizationResponse) error {
	return h.SlatomateHandler.GetAllOrganization(ctx, in, out)
}

func (h *slatomateHandler) GetOrganization(ctx context.Context, in *GetOrganizationRequest, out *Organization) error {
	return h.SlatomateHandler.GetOrganization(ctx, in, out)
}

func (h *slatomateHandler) DeleteOrganization(ctx context.Context, in *DeleteOrganizationRequest, out *emptypb.Empty) error {
	return h.SlatomateHandler.DeleteOrganization(ctx, in, out)
}

func (h *slatomateHandler) DeleteAllOrganization(ctx context.Context, in *emptypb.Empty, out *emptypb.Empty) error {
	return h.SlatomateHandler.DeleteAllOrganization(ctx, in, out)
}

func (h *slatomateHandler) CreateUser(ctx context.Context, in *CreateUserRequest, out *User) error {
	return h.SlatomateHandler.CreateUser(ctx, in, out)
}

func (h *slatomateHandler) GetUser(ctx context.Context, in *GetUserRequest, out *User) error {
	return h.SlatomateHandler.GetUser(ctx, in, out)
}

func (h *slatomateHandler) DeleteUser(ctx context.Context, in *DeleteUserRequest, out *emptypb.Empty) error {
	return h.SlatomateHandler.DeleteUser(ctx, in, out)
}

func (h *slatomateHandler) UpdateUser(ctx context.Context, in *UpdateUserRequest, out *User) error {
	return h.SlatomateHandler.UpdateUser(ctx, in, out)
}

func (h *slatomateHandler) LoginUser(ctx context.Context, in *User, out *User) error {
	return h.SlatomateHandler.LoginUser(ctx, in, out)
}

func (h *slatomateHandler) Me(ctx context.Context, in *APIKeyRequest, out *User) error {
	return h.SlatomateHandler.Me(ctx, in, out)
}

func (h *slatomateHandler) GenerateAPIKey(ctx context.Context, in *GenerateAPIKeyRequest, out *GenerateAPIKeyResponse) error {
	return h.SlatomateHandler.GenerateAPIKey(ctx, in, out)
}

func (h *slatomateHandler) GetAllUser(ctx context.Context, in *emptypb.Empty, out *GetAllUserResponse) error {
	return h.SlatomateHandler.GetAllUser(ctx, in, out)
}

func (h *slatomateHandler) CreateJob(ctx context.Context, in *Job, out *Job) error {
	return h.SlatomateHandler.CreateJob(ctx, in, out)
}

func (h *slatomateHandler) GetJob(ctx context.Context, in *GetJobRequest, out *Job) error {
	return h.SlatomateHandler.GetJob(ctx, in, out)
}

func (h *slatomateHandler) DeleteJob(ctx context.Context, in *DeleteJobRequest, out *emptypb.Empty) error {
	return h.SlatomateHandler.DeleteJob(ctx, in, out)
}

func (h *slatomateHandler) GetAllJob(ctx context.Context, in *GetAllJobRequset, out *GetAllJobResponse) error {
	return h.SlatomateHandler.GetAllJob(ctx, in, out)
}
