package slack

import (
	"os"

	"github.com/slack-go/slack"
)

type SlackClient struct {
	*slack.Client
}

func NewSlackClient(apikey string) *SlackClient {
	sc := slack.New(apikey, slack.OptionDebug(os.Getenv("SLATOMATE_DEBUG") == "true"))
	return &SlackClient{Client: sc}
}
