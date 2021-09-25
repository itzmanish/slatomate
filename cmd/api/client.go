package api

import (
	"github.com/itzmanish/go-micro/v2/client/grpc"
	slatomate "github.com/itzmanish/slatomate/proto/slatomate/v1"
)

var APIClient slatomate.SlatomateService

func init() {
	APIClient = NewClient()
}

func NewClient() slatomate.SlatomateService {
	c := grpc.NewClient()
	return slatomate.NewSlatomateService("github.itzmanish.service.slatomate", c)
}
