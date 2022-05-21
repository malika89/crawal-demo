package models

//worker 需要抓取的内容
type Task struct {
	Id         int64             `json:"id"  xorm:"id pk autoincr"`
	Url        string            `json:"url"  xorm:"url"`         //enginee
	Parsers    map[string]string `json:"parsers"  xorm:"parsers"` //parser
	StartTime  string            `json:"start_time"  xorm:"start_time"`
	FinishTime string            `json:"finish_time"  xorm:"finish_time"`
	Status     string            `json:"status"  xorm:"status"`
}

func (t *Task) TableName() string {
	return "task"
}

type TaskResult struct {
	Task  Task              `json:"task_id"  xorm:"task_id"`
	Value map[string]string `json:"result"  xorm:"result"`
}

func (r *TaskResult) TableName() string {
	return "task_result"
}
