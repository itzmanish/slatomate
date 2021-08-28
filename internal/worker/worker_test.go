package worker

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itzmanish/slatomate/internal/entity"
	"github.com/stretchr/testify/assert"
)

var testWorker = NewWorker()

func TestAdd(t *testing.T) {
	token := "xoxp-1001856848789-2040990405655-2381966723552-c657a0c560681078666336217006abd2"
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
