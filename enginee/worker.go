package enginee

import (
	"crawl/handler/models"
	"time"
)

var TasksChan = make(chan models.Task, 10)

func WorkParse(task models.Task) (ParseResult, error) {
	res, err := Fetch(task.Url)
	if err != nil {
		task.Status = "failed"
		task.FinishTime = time.Now().String()
		models.DBX.Table("task").Where("task_id=?", task.Id).Update(task)
		return ParseResult{}, err
	}
	parseResult := ParseResult{
		Task:  task,
		Value: nil,
	}
	parseResult.Parse(res)
	return parseResult, nil
}
func DBSaverWorker() (chan ParseResult, error) {
	out := make(chan ParseResult, 10)
	go func() {
		for {
			result := <-out
			models.DBX.Table("task").Where("task_id=?", result.Task.Id).Update(result.Task)
			models.DBX.Table("task_result").Insert(result)

		}
	}()
	return out, nil
}
