package main

import (
	"os"

	"github.com/itzmanish/go-micro-plugins/wrapper/trace/opentracing/v2"
	"github.com/itzmanish/go-micro/v2"
	log "github.com/itzmanish/go-micro/v2/logger"
	"github.com/itzmanish/slatomate/handler"
	"github.com/itzmanish/slatomate/internal/auth"
	"github.com/itzmanish/slatomate/internal/db"
	"github.com/itzmanish/slatomate/internal/repository"
	"github.com/itzmanish/slatomate/internal/types"
	"github.com/itzmanish/slatomate/subscriber"
	"github.com/itzmanish/slatomate/utils"
	"github.com/itzmanish/slatomate/wrapper"
	"github.com/joho/godotenv"

	slatomate "github.com/itzmanish/slatomate/proto/gen/slatomate/v1"
)

var (
	SERVICE_NAME                      = "github.itzmanish.service.slatomate"
	SERVICE_VERSION                   = "0.1.0"
	NoAuthEndpoint  types.ArrayString = []string{
		"Slatomate.CreateUser",
		"Slatomate.LoginUser",
		"Slatomate.AuthorizeOrganization",
		"Slatomate.Me",
		"Slatomate.Health",
	}
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Info("No .env files present in the root or Error loading .env")
	}
}

func main() {
	pdb, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	tracer, closer, err := utils.InitTracer(SERVICE_NAME, os.Getenv("JAEGER_AGENT_HOST"))
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()

	// New Service
	service := micro.NewService(
		micro.Name(SERVICE_NAME),
		micro.Version(SERVICE_VERSION),
		micro.WrapHandler(wrapper.AuthHandler(auth.NewAPIKeyAuth(repository.NewUserRepository(pdb)), NoAuthEndpoint)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(tracer)),
	)

	// Initialise service
	service.Init()

	// Register Handler
	slatomate.RegisterSlatomateServiceHandler(service.Server(),
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
