package subscriber

import (
	"context"

	log "github.com/itzmanish/go-micro/v2/logger"

	slatomate "github.com/itzmanish/slatomate/proto/slatomate"
)

type Slatomate struct{}

func (e *Slatomate) Handle(ctx context.Context, msg *slatomate.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *slatomate.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
