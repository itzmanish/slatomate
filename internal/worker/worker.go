package worker

import (
	"sync"

	"github.com/itzmanish/go-micro/v2/logger"
	"github.com/itzmanish/slatomate/internal/entity"
	"github.com/itzmanish/slatomate/internal/operator"
	"github.com/robfig/cron/v3"
)

type Worker struct {
	jobs sync.Map
	cron cron.Cron
}

func NewWorker() *Worker {
	w := &Worker{
		jobs: sync.Map{},
		cron: *cron.New(),
	}
	w.Start()
	return w
}

func (w *Worker) Add(token string, job *entity.Job) error {
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
	return nil
}

func (w *Worker) Remove(entryID cron.EntryID) {
	w.cron.Remove(entryID)
	w.jobs.Delete(entryID)
}

func (w *Worker) Start() {
	w.cron.Start()
}

func (w *Worker) Stop() error {
	return w.cron.Stop().Err()

}
