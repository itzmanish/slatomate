package subscriber

import (
	"context"
	"encoding/json"

	"github.com/itzmanish/go-micro/v2/errors"
	log "github.com/itzmanish/go-micro/v2/logger"

	"github.com/itzmanish/slatomate/internal/entity"
	"github.com/itzmanish/slatomate/internal/repository"
	"github.com/itzmanish/slatomate/internal/worker"
	slatomate "github.com/itzmanish/slatomate/proto/slatomate"
)

type Slatomate struct {
	orgRepo repository.OrganizationRepository
	worker  worker.Worker
}

func NewSubscriber(orgRepo repository.OrganizationRepository) *Slatomate {
	return &Slatomate{orgRepo: orgRepo, worker: worker.NewWorker()}
}

func (e *Slatomate) Handle(ctx context.Context, msg *slatomate.Message) error {
	log.Debug("Handler Received message: ", string(msg.GetBody()))
	evType, ok := msg.GetHeader()["type"]
	if !ok {
		return errors.BadRequest("EVENT_HANDLER", "Event type is not set")
	}
	switch evType {
	case "JOB_CREATED":
		log.Debug("Got event with type JOB_CREATED")
		var data entity.Job
		err := json.Unmarshal(msg.Body, &data)
		if err != nil {
			return err
		}
		org, err := e.orgRepo.GetOrganization(&entity.Organization{ID: data.OrganizationID})
		if err != nil {
			return err
		}
		err = e.worker.Add(org.SlackAPIKey, &data)
		if err != nil {
			return err
		}
	case "JOB_DELETED":
		log.Debug("Got event with type JOB_DELETED")
		e.worker.Remove(string(msg.GetBody()))

	default:
		return errors.BadRequest("EVENT_HANDLER", "Event type is unknown.")
	}
	return nil
}
