package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var XormDb *xorm.Engine

func Init() {
	// 1.连接数据库
	sqlStr := "root:123456@tcp(127.0.0.1:3306)/beego?charset=utf8&parseTime=true&loc=Local"
	var err error
	XormDb, err = xorm.NewEngine("mysql", sqlStr)
	if err != nil {
		log.Println("连接数据库失败", err)
		return
	}
	// 会在控制台打印执行的sql
	XormDb.ShowSQL(true)
}
