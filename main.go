package main

import (
	"github.com/itzmanish/go-micro/v2"
	log "github.com/itzmanish/go-micro/v2/logger"
	"github.com/itzmanish/slatomate/handler"
	"github.com/itzmanish/slatomate/subscriber"

	slatomate "github.com/itzmanish/slatomate/proto/slatomate"
)

var (
	SERVICE_NAME    = "github.itzmanish.service.slatomate"
	SERVICE_VERSION = "0.1.0"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name(SERVICE_NAME),
		micro.Version(SERVICE_VERSION),
	)

	// Initialise service
	service.Init()

	// Register Handler
	slatomate.RegisterSlatomateHandler(service.Server(), new(handler.Slatomate))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("github.itzmanish.service.slatomate", service.Server(), new(subscriber.Slatomate))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
