package worker

import (
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itzmanish/slatomate/internal/entity"
	"github.com/stretchr/testify/assert"
)

var testWorker = NewWorker()

func TestAdd(t *testing.T) {
	token := os.Getenv("SLACK_API_TOKEN")
	job := &entity.Job{
		ID:         uuid.New(),
		ScheduleAt: "* * * * *",
		Name:       "test job",
		Task:       entity.StatusUpdate,
		Data:       map[string]string{"emoji": ":fire:", "status": "Heya I am slatomate"},
	}
	err := testWorker.Add(token, job)
	assert.Nil(t, err)
	time.Sleep(70 * time.Second)
}
