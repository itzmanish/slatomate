package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/itzmanish/go-micro/v2/errors"
	"github.com/itzmanish/go-micro/v2/logger"
	"github.com/itzmanish/slatomate/internal/entity"
	slatomatepb "github.com/itzmanish/slatomate/proto/slatomate"
	"github.com/slack-go/slack"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *slatomateHandler) CreateOrganization(ctx context.Context, in *slatomatepb.CreateOrganizationRequest, out *slatomatepb.Organization) error {
	if len(in.GetName()) == 0 {
		return errors.BadRequest("CREATE_ORG_HANDLER", "Name is required.")
	}
	uid, err := uuid.Parse(in.UserId)
	if err != nil {
		return errors.BadRequest("CREATE_ORG_HANDLER", "user id is wrong.")
	}
	cp, err := h.orgRepo.CreateOrganization(&entity.Organization{Name: in.GetName(), UserID: uid})
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

func (h *slatomateHandler) AuthorizeOrganization(ctx context.Context, in *slatomatepb.AuthorizeOrganizationRequest, out *emptypb.Empty) error {
	logger.Info("Authorize Organization request: %v", in)
	if len(in.Code) == 0 {
		return errors.BadRequest("AUTHORIZE_ORG", "code is invalid!")
	}
	sres, err := slack.GetOAuthV2Response(http.DefaultClient, os.Getenv("SLACK_CLIENT_ID"), os.Getenv("SLACK_CLIENT_SECRET"), in.Code, os.Getenv("SLACK_REDIRECT_URI"))
	if err != nil {
		return err
	}
	logger.Info(sres)
	return nil
}
