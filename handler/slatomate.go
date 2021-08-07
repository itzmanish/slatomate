package handler

import (
	"context"

	slatomatepb "github.com/itzmanish/slatomate/proto/slatomate"
)

type Slatomate struct{}

func (h *Slatomate) CreateProject(ctx context.Context, in *slatomatepb.CreateProjectRequest, out *slatomatepb.Project) error {
	return nil
}

func (h *Slatomate) GetAllProjct(ctx context.Context, in *slatomatepb.GetAllProjectRequest, out *slatomatepb.GetAllProjectResponse) error {
	return nil
}

func (h *Slatomate) GetProject(ctx context.Context, in *slatomatepb.GetProjectRequest, out *slatomatepb.Project) error {
	return nil
}
