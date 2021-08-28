package subscriber

import (
	"context"
	"encoding/json"

	"github.com/itzmanish/go-micro/v2/errors"
	log "github.com/itzmanish/go-micro/v2/logger"

	"github.com/itzmanish/slatomate/internal/entity"
	"github.com/itzmanish/slatomate/internal/repository"
	slatomate "github.com/itzmanish/slatomate/proto/slatomate"
)

type Slatomate struct {
	orgRepo repository.OrganizationRepository
}

func NewSubscriber(orgRepo repository.OrganizationRepository) *Slatomate {
	return &Slatomate{orgRepo: orgRepo}
}

func (e *Slatomate) Handle(ctx context.Context, msg *slatomate.Message) error {
	log.Debug("Handler Received message: ", string(msg.GetBody()))
	evType, ok := msg.GetHeader()["type"]
	if !ok {
		return errors.BadRequest("EVENT", "Event type is not set")
	}
	switch evType {
	case "JOB":
		var data entity.Job
		err := json.Unmarshal(msg.Body, &data)
		if err != nil {
			return err
		}
		log.Infof("%+v", data)
	default:
		return errors.BadRequest("EVENT", "Event type is unknown.")
	}
	return nil
}
