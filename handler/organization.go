package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/itzmanish/go-micro/v2/errors"
	"github.com/itzmanish/go-micro/v2/logger"
	"github.com/itzmanish/slatomate/internal/auth"
	"github.com/itzmanish/slatomate/internal/entity"
	slatomatepb "github.com/itzmanish/slatomate/proto/slatomate"
	"github.com/slack-go/slack"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *slatomateHandler) CreateOrganization(ctx context.Context, in *slatomatepb.CreateOrganizationRequest, out *slatomatepb.Organization) error {
	user, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("CREATE_ORG_HANDLER", "Unauthorized access.")
	}
	if len(in.GetName()) == 0 {
		return errors.BadRequest("CREATE_ORG_HANDLER", "Name is required.")
	}
	orgin := &entity.Organization{Name: in.GetName()}
	err := h.orgRepo.CreateOrganization(user.ID, orgin)
	if err != nil {
		return err
	}
	*out = entity.DeserializeOrganization(orgin)
	return nil
}

func (h *slatomateHandler) GetAllOrganization(ctx context.Context, in *slatomatepb.GetAllOrganizationRequest, out *slatomatepb.GetAllOrganizationResponse) error {
	user, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("GET_ALL_ORG_HANDLER", "Unauthorized access.")
	}

	res, err := h.orgRepo.GetAllOrganization(user.ID)
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
	user, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("GET_ORG_HANDLER", "Unauthorized access.")
	}

	if len(in.Id) == 0 && len(in.Name) == 0 {
		return errors.BadRequest("GET_ORG_HANDLER", "Organization id/name is required!")
	}
	org := &entity.Organization{UserID: user.ID}
	if len(in.GetId()) != 0 {
		oid, err := uuid.Parse(in.GetId())
		if err != nil {
			return errors.BadRequest("GET_ORG_HANDLER", "Organization id is invalid!")
		}
		org.ID = oid
	} else {
		org.Name = in.GetName()
	}

	org, err := h.orgRepo.GetOrganization(org)
	if err != nil {
		return err
	}
	*out = entity.DeserializeOrganization(org)
	return nil
}

func (h *slatomateHandler) DeleteOrganization(ctx context.Context, in *slatomatepb.DeleteOrganizationRequest, out *emptypb.Empty) error {
	user, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("DELETE_ORG_HANDLER", "Unauthorized access.")
	}

	if len(in.Id) == 0 {
		return errors.BadRequest("DELETE_ORG_HANDLER", "Organization id is required!")
	}

	oid, err := uuid.Parse(in.GetId())
	if err != nil {
		return errors.BadRequest("DELETE_ORG_HANDLER", "Organization id is invalid!")
	}

	return h.orgRepo.DeleteOrganization(&entity.Organization{ID: oid, UserID: user.ID})
}

func (h *slatomateHandler) DeleteAllOrganization(ctx context.Context, in *emptypb.Empty, out *emptypb.Empty) error {
	user, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("DELETE_ORG_HANDLER", "Unauthorized access.")
	}

	return h.orgRepo.DeleteOrganization(&entity.Organization{UserID: user.ID})
}

func (h *slatomateHandler) ValidateOrgAccess(ctx context.Context, in *slatomatepb.ValidateOrgAccessRequest, out *slatomatepb.ValidateOrgAccessResponse) error {
	logger.Debugf("Validate Organization request: %v", in)
	if len(in.OrgId) == 0 {
		return errors.BadRequest("AUTHORIZE_ORG", "Organization id is required!")
	}

	oid, err := uuid.Parse(in.OrgId)
	if err != nil {
		return errors.BadRequest("AUTHORIZE_ORG", "Organization id is invalid!")
	}
	_, err = validateOrganizationAccess(ctx, oid)
	if err != nil {
		return err
	}
	*out = slatomatepb.ValidateOrgAccessResponse{HasAccess: true}
	return nil
}

func (h *slatomateHandler) AuthorizeOrganization(ctx context.Context, in *slatomatepb.AuthorizeOrganizationRequest, out *emptypb.Empty) error {
	logger.Debugf("Authorize Organization request: %v", in)
	if len(in.OrgId) == 0 {
		return errors.BadRequest("AUTHORIZE_ORG", "Organization id is required!")
	}

	oid, err := uuid.Parse(in.OrgId)
	if err != nil {
		return errors.BadRequest("AUTHORIZE_ORG", "Organization id is invalid!")
	}
	_, err = validateOrganizationAccess(ctx, oid)
	if err != nil {
		return err
	}

	if len(in.Code) == 0 {
		return errors.BadRequest("AUTHORIZE_ORG", "code is invalid!")
	}

	sres, err := slack.GetOAuthV2Response(http.DefaultClient, os.Getenv("SLACK_CLIENT_ID"), os.Getenv("SLACK_CLIENT_SECRET"), in.Code, os.Getenv("SLACK_REDIRECT_URI"))
	if err != nil {
		return err
	}
	_, err = h.orgRepo.UpdateOrganization(&entity.Organization{ID: oid, SlackAPIKey: sres.AuthedUser.AccessToken})
	if err != nil {
		return err
	}
	return nil
}

func validateOrganizationAccess(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, ok := auth.AccountFromContext(ctx)
	if !ok {
		return user, errors.Unauthorized("VALIDATE_ORG_ACCESS", "Unauthorized access.")
	}

	if ok := user.HaveOrg(id); !ok {
		return user, errors.Unauthorized("VALIDATE_ORG_ACCESS", "Unauthorized access to this organization.")
	}
	return user, nil
}
