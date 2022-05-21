package views

import (
	"crawl/handler/services"
	"crawl/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)


func QueryUserByNameHandler(c *gin.Context) {
	// 获取参数
	name := c.Query("name")
	userService := services.GetService()
	userDao, err := userService.Query("name",name)
	if err !=nil {
		utils.BadResponse(c,utils.Error, fmt.Sprintf("请求数据库错误:%v",err))
		return
	}
	utils.NormalResponse(c, utils.OK, "success",userDao,len(userDao))
}


func UpdateUserHandler(c *gin.Context) {
	// 获取参数
	userService := services.GetService()
	if err :=c.BindJSON(&userService.Model);err !=nil{
		utils.BadResponse(c,utils.Error,"传参错误")
		return
	}
	err := userService.Update()
	if err !=nil {
		utils.BadResponse(c,utils.Error, fmt.Sprintf("修改失败:%v",err))
		return
	}
	utils.NormalResponse(c, utils.OK, "修改成功",nil,0)
}
func AddUserHandler(c *gin.Context) {
	// 获取参数
	userService := services.GetService()
	if err :=c.BindJSON(&userService.Model);err !=nil{
		utils.BadResponse(c,utils.Error,"传参错误")
		return
	}
	err := userService.Add()
	if err !=nil {
		utils.BadResponse(c,utils.Error, fmt.Sprintf("新增记录失败:%v",err))
		return
	}
	utils.NormalResponse(c, utils.OK, "新增记录成功",nil,0)
}
func DeleteUserHandler(c *gin.Context) {
	// 获取参数
	userService := services.GetService()
	if err :=c.BindJSON(&userService.Model);err !=nil{
		utils.BadResponse(c,utils.Error,"传参错误")
		return
	}
	err := userService.Delete()
	if err !=nil {
		utils.BadResponse(c,utils.Error, fmt.Sprintf("删除失败:%v",err))
		return
	}
	utils.NormalResponse(c, utils.OK, "删除成功",nil,0)
}


