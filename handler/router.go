package handler


import (
	"crawl/handler/views"
	"github.com/gin-gonic/gin"
)


func setupRoute(engine *gin.Engine) {

	engine.GET("/", func(c *gin.Context) {
		c.JSON(200,gin.H{"ping":"pong"})
	})

	test :=engine.Group("mock")
	{
		test.POST("addTask", views.AddTaskHandler)
	}

	v1 :=engine.Group("/api/tasks")
	{
		v1.GET("/search",views.QueryUserByNameHandler)
		v1.POST("/",views.AddTaskHandler)
	}
}