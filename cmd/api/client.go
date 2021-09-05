package api

import (
	"github.com/itzmanish/go-micro/v2/client/grpc"
	"github.com/itzmanish/slatomate/proto/slatomate"
)

var APIClient slatomate.SlatomateService

func init() {
	APIClient = NewClient()
}

func NewClient() slatomate.SlatomateService {
	c := grpc.NewClient()
	return slatomate.NewSlatomateService("github.itzmanish.service.slatomate", c)
}
