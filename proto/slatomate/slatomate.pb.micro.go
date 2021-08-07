// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/slatomate/slatomate.proto

package slatomate

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	return []*api.Endpoint{}
}

// Client API for Slatomate service

type SlatomateService interface {
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...client.CallOption) (*Project, error)
	GetAllProjct(ctx context.Context, in *GetAllProjectRequest, opts ...client.CallOption) (*GetAllProjectResponse, error)
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...client.CallOption) (*Project, error)
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

func (c *slatomateService) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...client.CallOption) (*Project, error) {
	req := c.c.NewRequest(c.name, "Slatomate.CreateProject", in)
	out := new(Project)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) GetAllProjct(ctx context.Context, in *GetAllProjectRequest, opts ...client.CallOption) (*GetAllProjectResponse, error) {
	req := c.c.NewRequest(c.name, "Slatomate.GetAllProjct", in)
	out := new(GetAllProjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slatomateService) GetProject(ctx context.Context, in *GetProjectRequest, opts ...client.CallOption) (*Project, error) {
	req := c.c.NewRequest(c.name, "Slatomate.GetProject", in)
	out := new(Project)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Slatomate service

type SlatomateHandler interface {
	CreateProject(context.Context, *CreateProjectRequest, *Project) error
	GetAllProjct(context.Context, *GetAllProjectRequest, *GetAllProjectResponse) error
	GetProject(context.Context, *GetProjectRequest, *Project) error
}

func RegisterSlatomateHandler(s server.Server, hdlr SlatomateHandler, opts ...server.HandlerOption) error {
	type slatomate interface {
		CreateProject(ctx context.Context, in *CreateProjectRequest, out *Project) error
		GetAllProjct(ctx context.Context, in *GetAllProjectRequest, out *GetAllProjectResponse) error
		GetProject(ctx context.Context, in *GetProjectRequest, out *Project) error
	}
	type Slatomate struct {
		slatomate
	}
	h := &slatomateHandler{hdlr}
	return s.Handle(s.NewHandler(&Slatomate{h}, opts...))
}

type slatomateHandler struct {
	SlatomateHandler
}

func (h *slatomateHandler) CreateProject(ctx context.Context, in *CreateProjectRequest, out *Project) error {
	return h.SlatomateHandler.CreateProject(ctx, in, out)
}

func (h *slatomateHandler) GetAllProjct(ctx context.Context, in *GetAllProjectRequest, out *GetAllProjectResponse) error {
	return h.SlatomateHandler.GetAllProjct(ctx, in, out)
}

func (h *slatomateHandler) GetProject(ctx context.Context, in *GetProjectRequest, out *Project) error {
	return h.SlatomateHandler.GetProject(ctx, in, out)
}
