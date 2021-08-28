package operator

import (
	"testing"

	"github.com/google/uuid"
	"github.com/itzmanish/slatomate/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestProcessTask(t *testing.T) {
	token := "xoxp-1001856848789-2040990405655-2381966723552-c657a0c560681078666336217006abd2"
	job := &entity.Job{
		ID:   uuid.New(),
		Name: "Test status update",
		Task: entity.StatusUpdate,
		Data: map[string]string{"emoji": ":fire:", "status": "Heya I am slatomate"},
	}
	err := Process(token, job)
	assert.Nil(t, err)
	// clear status
	err = Process(token, &entity.Job{Task: entity.StatusUpdate})
	assert.Nil(t, err)
}
