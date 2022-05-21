package views

import (
	"crawl/enginee"
	"crawl/handler/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func AddTaskHandler(c *gin.Context) {
	var data models.Task
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	task := models.Task{
		Url:       data.Url,
		StartTime: time.Now().String(),
		Parsers:   data.Parsers,
		Status:    "pending",
	}
	if _, err := models.DBX.Table("task").Insert(task); err != nil {
		c.JSON(400, gin.H{"message": fmt.Sprintf("任务写入数据库失败:%s", err.Error())})
		return
	}
	enginee.TasksChan <- task
	c.JSON(200, gin.H{"message": "任务添加成功"})
}
