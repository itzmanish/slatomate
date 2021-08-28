package main

import (
	"github.com/itzmanish/go-micro/v2"
	log "github.com/itzmanish/go-micro/v2/logger"
	"github.com/itzmanish/slatomate/handler"
	"github.com/itzmanish/slatomate/internal/db"
	"github.com/itzmanish/slatomate/internal/repository"
	"github.com/itzmanish/slatomate/subscriber"
	"github.com/joho/godotenv"

	slatomate "github.com/itzmanish/slatomate/proto/slatomate"
)

var (
	SERVICE_NAME    = "github.itzmanish.service.slatomate"
	SERVICE_VERSION = "0.1.0"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Info("No .env files present in the root or Error loading .env")
	}
}

func main() {
	// New Service
	service := micro.NewService(
		micro.Name(SERVICE_NAME),
		micro.Version(SERVICE_VERSION),
	)

	// Initialise service
	service.Init()

	pdb, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	// Register Handler
	slatomate.RegisterSlatomateHandler(service.Server(),
		handler.NewHandler(repository.NewUserRepository(pdb),
			repository.NewOrganizationRepository(pdb),
			repository.NewJobRepository(pdb), micro.NewEvent(SERVICE_NAME, service.Client())))

	// Register Struct as Subscriber
	micro.RegisterSubscriber(SERVICE_NAME, service.Server(), subscriber.NewSubscriber(repository.NewOrganizationRepository(pdb)))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
