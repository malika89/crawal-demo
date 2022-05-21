package scheduler

import (
	"crawl/enginee"
	"crawl/handler/models"
	"log"
)

type QueueScheduler struct {
	requestChan chan models.Task
	workerChan  chan chan models.Task
}

func Schedule() {
	DBSaverWorker, _ := enginee.DBSaverWorker()
	e := enginee.Engine{
		Scheduler:   &QueueScheduler{},
		WorkerCount: 10,
		DBChan:      DBSaverWorker,
	}
	e.Run()
}

func (s *QueueScheduler) WorkerChan() chan models.Task {
	return make(chan models.Task, 10)
}

func (s *QueueScheduler) Submit(t models.Task) {
	log.Println("request channel received new tasks")
	s.requestChan <- t

}

func (s *QueueScheduler) WorkerReady(w chan models.Task) {
	s.workerChan <- w
}

//GMP 调度
func (s *QueueScheduler) Run() {
	log.Println("QueueScheduler  start to schedule.....")
	s.workerChan = make(chan chan models.Task, 10)
	s.requestChan = make(chan models.Task, 10)
	go func() {
		var requestQ []models.Task
		var workerQ []chan models.Task
		for {
			var activeRequest models.Task
			var activeWorker chan models.Task
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
