package operator

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/itzmanish/slatomate/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestProcessTask(t *testing.T) {
	token := os.Getenv("SLACK_API_TOKEN")
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
