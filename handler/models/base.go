package models

import (
	"crawl/conf"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/goinggo/mapstructure"
	"log"
)

var DBX *xorm.Engine

func Init() {
	if DBX == nil {
		var err error
		//"root:123456@tcp(127.0.0.1:3306)/beego?charset=utf8&parseTime=true&loc=Local"
		connStr := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local`,
			conf.Conf.DataBase.Username, conf.Conf.DataBase.Password, conf.Conf.DataBase.Host, conf.Conf.DataBase.Port, conf.Conf.DataBase.Name)
		DBX, err = xorm.NewEngine(conf.Conf.DataBase.Driver, connStr)
		if err != nil {
			log.Println("init database error", err)
		}
		DBX.ShowSQL(true)
	}
}

type BaseModel struct {
	Table string
	Model interface{} //struct结构体
}

func (t *BaseModel) TableName() string {
	return "base_model"
}
func (t *BaseModel) GetSession() *xorm.Session {
	if t.Table == "" {
		t.Table = t.TableName()
	}
	dbSession := DBX.Table(t.Table)
	return dbSession
}
func (t *BaseModel) Exists(keyword string, value string) (bool, error) {
	infStruct := t.Model
	dbSession := t.GetSession()
	filterMap := map[string]string{keyword: value}
	if err := mapstructure.Decode(filterMap, infStruct); err != nil {
		return false, err
	}
	return dbSession.Exist(&infStruct)
}
func (t *BaseModel) Query(keyword string, value string) ([]map[string]string, error) {
	dbSession := t.GetSession()
	infLst, err := dbSession.Where(keyword+" = ?", value).QueryString()
	return infLst, err
}
func (t *BaseModel) Add() error {
	dbSession := t.GetSession()
	if _, err := dbSession.Insert(t.Model); err != nil {
		return err
	}
	return nil
}
func (t *BaseModel) Update() error {
	dbSession := t.GetSession()
	modelMap := t.Model.(map[string]interface{})
	if _, err := dbSession.ID(modelMap["id"].(string)).Update(modelMap); err != nil {
		return err
	}
	return nil
}

//Delete 支持批量删除
func (t *BaseModel) Delete() error {
	dbSession := t.GetSession()
	if _, err := dbSession.Delete(t.Model); err != nil {
		return err
	}
	return nil
}
