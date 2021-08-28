package worker

import (
	"sync"

	"github.com/itzmanish/go-micro/v2/logger"
	"github.com/itzmanish/slatomate/internal/entity"
	"github.com/itzmanish/slatomate/internal/operator"
	"github.com/robfig/cron/v3"
)

type Worker interface {
	Add(token string, job *entity.Job) error
	Remove(jobID string)
	Start()
	Stop() error
}

type worker struct {
	jobs sync.Map
	cron cron.Cron
}

func NewWorker() Worker {
	w := &worker{
		jobs: sync.Map{},
		cron: *cron.New(),
	}
	w.Start()
	return w
}

func (w *worker) Add(token string, job *entity.Job) error {
	entryID, err := w.cron.AddFunc(job.ScheduleAt, func() {
		err := operator.Process(token, job)
		if err != nil {
			logger.Error(err)
		}
	})
	if err != nil {
		return err
	}

	w.jobs.Store(entryID, job)
	logger.Debugf("Job added for schedule_at: %s with job id: %s and entry id: %v", job.ScheduleAt, job.ID, entryID)
	return nil
}

func (w *worker) Remove(jobID string) {
	w.jobs.Range(func(key, value interface{}) bool {
		if value != nil {
			if job, ok := value.(*entity.Job); ok {
				if job.ID.String() == jobID {
					w.remove(key.(cron.EntryID))
					return false
				}
			}
		}
		return true
	})
}

func (w *worker) remove(entryID cron.EntryID) {
	w.cron.Remove(entryID)
	w.jobs.Delete(entryID)
	logger.Debugf("Job with entry id: %s removed", entryID)
}

func (w *worker) Start() {
	w.cron.Start()
	logger.Debug("Worker started.")
}

func (w *worker) Stop() error {
	return w.cron.Stop().Err()

}
