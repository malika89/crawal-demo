package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Errno struct {
	Code    int
	Message string
}


type ResponseJSON struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"success"`
	Object  interface{} `json:"result"`
	TotalCount interface{} `json:"total_count"`
}
var (
	OK                 = 200
	Error              = 400
	NotFound           = 404
	ServerError        = 500
)


func NormalResponse(c *gin.Context, code int, message string, object interface{},total int) {
	var resp ResponseJSON
	resp = ResponseJSON{Code: code, Message: message, Object: object,TotalCount: total}
	c.JSON(200, resp)
}

func BadResponse(c *gin.Context, code int, message string) {
	var resp ResponseJSON
	resp = ResponseJSON{Code: code, Message: message,Object: nil,TotalCount: 0}
	c.JSON(200, resp)
}


func GetPageInfo(c *gin.Context) (int,int) {
	var (
		page int
		pageSize int
		err error
	)
	if page,err = strconv.Atoi(c.DefaultQuery("page","1")); err!=nil{
		page = 1
	}
	if pageSize,err = strconv.Atoi(c.DefaultQuery("page_size","20")); err!=nil{
		pageSize = 20
	}
	return page,pageSize
}
