package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/itzmanish/slatomate/proto/slatomate"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Status int

const (
	ActiveStatus Status = iota
	InActive
)

// Job represent a Job
type Job struct {
	ID             uuid.UUID `json:"uuid" gorm:"primary_key; unique; type:uuid;"`
	Name           string    `json:"name" gorm:"type:varchar(100)"`
	ScheduleAt     string    `json:"schedule_at" gorm:"type:varchar(100)"`
	Task           string    `json:"task" gorm:"type:varchar(100)"`
	Status         Status    `json:"status"`
	Data           map[string]string
	OrganizationID uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// BeforeCreate will set a UUID rather than numeric ID.
func (job *Job) BeforeCreate(tx *gorm.DB) error {
	u := uuid.New()
	job.ID = u
	return nil
}

//SerializeJob converts proto Job to Job struct
func SerializeJob(in *slatomate.Job) Job {
	if in == nil {
		return Job{}
	}
	job := Job{
		Name:       in.Name,
		ScheduleAt: in.ScheduleAt,
		Task:       in.Task.String(),
		Data:       in.Data,
	}

	return job
}

//DeserializeJob converts Job to proto Job
func DeserializeJob(in *Job) slatomate.Job {
	return slatomate.Job{
		Id:         in.ID.String(),
		Name:       in.Name,
		OrgId:      in.OrganizationID.String(),
		ScheduleAt: in.ScheduleAt,
		Task:       GetTask(in.Task),
		Status:     GetStatusToString(in.Status),
		Data:       in.Data,
		CreatedAt:  timestamppb.New(in.CreatedAt),
	}
}

func GetTask(task string) slatomate.Task {
	switch task {
	case "status_update":
		return slatomate.Task_STATUS_UPDATE
	default:
		return slatomate.Task_DEFAULT
	}
}

func GetStatusToString(status Status) string {
	switch status {
	case ActiveStatus:
		return "active"
	case InActive:
		return "inactive"

	default:
		return "unknown"
	}
}
