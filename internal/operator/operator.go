package operator

import (
	"strconv"

	"github.com/itzmanish/slatomate/internal/entity"
	"github.com/mitchellh/mapstructure"
	"github.com/slack-go/slack"
)

type StatusUpdateData struct {
	Emoji      string `mapstructure:"emoji"`
	Status     string `mapstructure:"status"`
	ClearAfter string `mapstructure:"clear_after"`
}

func Process(token string, job *entity.Job) error {
	switch job.Task {
	case entity.NoOp:
		return nil
	case entity.StatusUpdate:
		return updateStatus(token, job)
	default:
		return nil
	}
}

func updateStatus(token string, job *entity.Job) error {
	data := StatusUpdateData{}
	err := mapstructure.Decode(job.Data, &data)
	if err != nil {
		return err
	}
	var clear_after int
	if len(data.ClearAfter) != 0 {
		clear_after, err = strconv.Atoi(data.ClearAfter)
		if err != nil {
			return err
		}
	}
	sc := slack.New(token)
	return sc.SetUserCustomStatus(data.Status, data.Emoji, int64(clear_after))
}
