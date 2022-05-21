package enginee

import (
	"crawl/handler/models"
	"log"
)

type Engine struct {
	WorkerCount int
	Scheduler   Scheduler
	DBChan      chan ParseResult //数据处理通道
}

type Processor func(task models.Task) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(task models.Task)
	WorkerChan() chan models.Task
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan models.Task)
}

func (e *Engine) Run(tasks ...models.Task) {
	log.Println("engine run tasks.....")
	respChan := make(chan ParseResult, 10)
	//调度算法,scheduler(reqchan,workerchan)
	e.Scheduler.Run()

	//new 处理器：processor( workchan<-data <-reqchan)
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), respChan, e.Scheduler)
	}

	//提交新请求任务到调度器
	for _, task := range tasks {
		if isDuplicate(task) {
			continue
		}
		e.Scheduler.Submit(task)
	}

	for {
		task, ok := <-TasksChan
		if !ok {
			break
		}
		e.Scheduler.Submit(task)
	}
	// readDa

	// 保存结果
	for {
		result := <-respChan
		if isDuplicate(result.Task) {
			continue
		}
		//结果数据放入数据处理worker通道进行更新
		e.DBChan <- result
		e.Scheduler.Submit(result.Task)
	}

}

func (e *Engine) createWorker(req chan models.Task, resp chan ParseResult, ready ReadyNotifier) {
	log.Println("worker created")
	go func() {
		for {
			ready.WorkerReady(req)
			task := <-req
			result, err := WorkParse(task)
			if err != nil {
				continue
			}
			resp <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(task models.Task) bool {
	if visitedUrls[task.Url] {
		return true
	}

	visitedUrls[task.Url] = true
	return false
}
