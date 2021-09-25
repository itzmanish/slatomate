package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	slatomate "github.com/itzmanish/slatomate/proto/gen/slatomate/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Status int
type Task int
type JSONData map[string]string

const (
	ActiveStatus Status = iota
	InActive
)
const (
	NoOp Task = iota
	StatusUpdate
)

// Job represent a Job
type Job struct {
	ID             uuid.UUID `json:"id" gorm:"primary_key; unique; type:uuid;"`
	Name           string    `json:"name" gorm:"type:varchar(100)"`
	ScheduleAt     string    `json:"schedule_at" gorm:"type:varchar(100)"`
	Task           Task      `json:"task" gorm:"type:int"`
	Status         Status    `json:"status"`
	Data           JSONData  `json:"data" gorm:"type:json"`
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

func (jsonData *JSONData) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := JSONData{}
	err := json.Unmarshal(bytes, &result)
	*jsonData = result
	return err
}

func (jsonData *JSONData) Value() (driver.Value, error) {
	if len(*jsonData) == 0 {
		return nil, nil
	}
	return json.Marshal(jsonData)
}
func (JSONData) GormDataType() string {
	return "json"
}

//SerializeJob converts proto Job to Job struct
func SerializeJob(in *slatomate.Job) *Job {
	if in == nil {
		return &Job{}
	}
	job := Job{
		Name:       in.Name,
		ScheduleAt: in.ScheduleAt,
		Task:       GetTaskFromProtoTask(in.Task),
		Data:       in.Data,
	}

	return &job
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

func GetTask(task Task) slatomate.Task {
	switch task {
	case StatusUpdate:
		return slatomate.Task_TASK_STATUS_UPDATE
	default:
		return slatomate.Task_TASK_UNSPECIFIED
	}
}

func GetTaskFromProtoTask(task slatomate.Task) Task {
	switch task {
	case slatomate.Task_TASK_STATUS_UPDATE:
		return StatusUpdate
	case slatomate.Task_TASK_UNSPECIFIED:
		return NoOp
	default:
		return NoOp
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
