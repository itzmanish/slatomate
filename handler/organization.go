package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/itzmanish/go-micro/v2/errors"
	"github.com/itzmanish/slatomate/internal/entity"
	slatomatepb "github.com/itzmanish/slatomate/proto/slatomate"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *slatomateHandler) CreateOrganization(ctx context.Context, in *slatomatepb.CreateOrganizationRequest, out *slatomatepb.Organization) error {
	if len(in.GetName()) == 0 || len(in.GetSlackApikey()) == 0 {
		return errors.BadRequest("CREATE_ORG_HANDLER", "Name and slack api key both are required.")
	}
	uid, err := uuid.Parse(in.UserId)
	if err != nil {
		return errors.BadRequest("CREATE_ORG_HANDLER", "user id is wrong.")
	}
	cp, err := h.orgRepo.CreateOrganization(&entity.Organization{Name: in.GetName(), SlackAPIKey: in.SlackApikey, UserID: uid})
	if err != nil {
		return err
	}
	*out = entity.DeserializeOrganization(cp)
	return nil
}

func (h *slatomateHandler) GetAllOrganization(ctx context.Context, in *slatomatepb.GetAllOrganizationRequest, out *slatomatepb.GetAllOrganizationResponse) error {
	if len(in.GetUserId()) == 0 {
		return errors.BadRequest("GET_ALL_ORG_HANDLER", "User id is required!")
	}
	id, err := uuid.Parse(in.GetUserId())
	if err != nil {
		return errors.BadRequest("GET_ALL_ORG_HANDLER", "User id is wrong!")
	}
	res, err := h.orgRepo.GetAllOrganization(id)
	if err != nil {
		return err
	}

	out.Count = int32(len(res))
	out.Organizations = make([]*slatomatepb.Organization, len(res))
	for i, v := range res {
		dorg := entity.DeserializeOrganization(v)
		out.Organizations[i] = &dorg
	}
	return nil
}

func (h *slatomateHandler) GetOrganization(ctx context.Context, in *slatomatepb.GetOrganizationRequest, out *slatomatepb.Organization) error {
	if len(in.Id) == 0 || len(in.UserId) == 0 {
		return errors.BadRequest("GET_ORG_HANDLER", "Organization id and User id both are required!")
	}
	oid, err := uuid.Parse(in.GetId())
	if err != nil {
		return errors.BadRequest("GET_ORG_HANDLER", "Organization id is invalid!")
	}
	uid, err := uuid.Parse(in.GetUserId())
	if err != nil {
		return errors.BadRequest("GET_ORG_HANDLER", "User id is invalid!")
	}
	org, err := h.orgRepo.GetOrganization(&entity.Organization{ID: oid, UserID: uid})
	if err != nil {
		return err
	}
	*out = entity.DeserializeOrganization(org)
	return nil
}

func (h *slatomateHandler) DeleteOrganization(ctx context.Context, in *slatomatepb.DeleteOrganizationRequest, out *emptypb.Empty) error {
	if len(in.Id) == 0 || len(in.UserId) == 0 {
		return errors.BadRequest("DELETE_ORG_HANDLER", "Organization id and User id both are required!")
	}

	oid, err := uuid.Parse(in.GetId())
	if err != nil {
		return errors.BadRequest("DELETE_ORG_HANDLER", "Organization id is invalid!")
	}
	uid, err := uuid.Parse(in.GetUserId())
	if err != nil {
		return errors.BadRequest("DELETE_ORG_HANDLER", "User id is invalid!")
	}

	return h.orgRepo.DeleteOrganization(&entity.Organization{ID: oid, UserID: uid})
}
