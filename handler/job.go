package handler

import (
	"context"

	slatomatepb "github.com/itzmanish/slatomate/proto/slatomate"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *slatomateHandler) CreateJob(ctx context.Context, in *slatomatepb.CreateJobRequest, out *slatomatepb.Job) error {
	return nil
}

func (h *slatomateHandler) GetJob(ctx context.Context, in *slatomatepb.GetJobRequest, out *slatomatepb.Job) error {
	return nil
}

func (h *slatomateHandler) DeleteJob(ctx context.Context, in *slatomatepb.DeleteJobRequest, out *emptypb.Empty) error {
	return nil
}

func (h *slatomateHandler) GetAllJob(ctx context.Context, in *slatomatepb.GetAllJobRequset, out *slatomatepb.GetAllJobResponse) error {
	return nil
}
