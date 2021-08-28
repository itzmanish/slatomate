package handler

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/itzmanish/go-micro/v2/errors"
	"github.com/itzmanish/go-micro/v2/logger"
	"github.com/itzmanish/slatomate/internal/entity"
	slatomatepb "github.com/itzmanish/slatomate/proto/slatomate"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *slatomateHandler) CreateJob(ctx context.Context, in *slatomatepb.Job, out *slatomatepb.Job) error {
	logger.Debug("Create job request: ", in)
	if len(in.GetOrgId()) == 0 || len(in.GetScheduleAt()) == 0 || len(in.GetTask().String()) == 0 || len(in.GetName()) == 0 {
		return errors.BadRequest("CREATE_JOB_HANDLER", "org_id,schedule_at,task,name all fields are required!")
	}
	serialized := entity.SerializeJob(in)
	orgID, err := uuid.Parse(in.GetOrgId())
	if err != nil {
		return err
	}
	err = h.jobRepo.CreateJob(orgID, serialized)
	if err != nil {
		return err
	}
	data, err := json.Marshal(serialized)
	if err != nil {
		return err
	}
	err = h.publisher.Publish(context.TODO(), &slatomatepb.Message{Header: map[string]string{"type": "JOB"}, Body: data})
	if err != nil {
		return err
	}
	*out = entity.DeserializeJob(serialized)
	return nil
}

func (h *slatomateHandler) GetJob(ctx context.Context, in *slatomatepb.GetJobRequest, out *slatomatepb.Job) error {
	logger.Debug("Get job request: ", in)
	if len(in.GetId()) == 0 || len(in.GetOrgId()) == 0 {
		return errors.BadRequest("GET_JOB_HANDLER", "org_id and job id both fields are required!")
	}
	oid, err := uuid.Parse(in.GetOrgId())
	if err != nil {
		return errors.BadRequest("GET_JOB_HANDLER", "Organization id is invalid!")
	}
	jid, err := uuid.Parse(in.GetId())
	if err != nil {
		return errors.BadRequest("GET_JOB_HANDLER", "Job id is invalid!")
	}
	job, err := h.jobRepo.GetJob(&entity.Job{ID: jid, OrganizationID: oid, Name: in.Name})
	if err != nil {
		return err
	}
	*out = entity.DeserializeJob(job)
	return nil
}

func (h *slatomateHandler) DeleteJob(ctx context.Context, in *slatomatepb.DeleteJobRequest, out *emptypb.Empty) error {
	logger.Debug("Delete job request: ", in)
	if len(in.Id) == 0 || len(in.OrgId) == 0 {
		return errors.BadRequest("DELETE_JOB_HANDLER", "Organization id and JOb id both are required!")
	}

	oid, err := uuid.Parse(in.GetId())
	if err != nil {
		return errors.BadRequest("DELETE_JOB_HANDLER", "Organization id is invalid!")
	}
	jid, err := uuid.Parse(in.GetId())
	if err != nil {
		return errors.BadRequest("DELETE_JOB_HANDLER", "Job id is invalid!")
	}
	return h.jobRepo.DeleteJob(&entity.Job{ID: jid, OrganizationID: oid})
}

func (h *slatomateHandler) GetAllJob(ctx context.Context, in *slatomatepb.GetAllJobRequset, out *slatomatepb.GetAllJobResponse) error {
	logger.Debug("GetAllJob request: ", in)
	if len(in.GetOrgId()) == 0 {
		return errors.BadRequest("GET_ALL_JOB_HANDLER", "Organization id is required!")
	}
	id, err := uuid.Parse(in.OrgId)
	if err != nil {
		return errors.BadRequest("GET_ALL_JOB_HANDLER", "Organization id is wrong!")
	}
	res, err := h.jobRepo.GetAllJob(id)
	if err != nil {
		return err
	}

	out.Count = int32(len(res))
	out.Jobs = make([]*slatomatepb.Job, len(res))
	for i, v := range res {
		djob := entity.DeserializeJob(v)
		out.Jobs[i] = &djob
	}

	return nil
}
