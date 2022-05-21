package parser

import (
	"crawl/handler/models"
	"github.com/goinggo/mapstructure"
	"log"
	"regexp"
)

//result : 任务url,id,正则表达式,结果body
type ParseResult struct {
	Task  models.Task
	Value map[string]string
}

type OrmExample struct {
	name   string
	age    int
	active bool
}

type Parser interface {
	Parse(content []byte)
	ToOrm(model interface{}) interface{}
}

// 传递到item
func (res *ParseResult) Parse(content []byte) {
	for field, reg := range res.Task.Parsers {
		resRe := regexp.MustCompile(reg)
		val := extractString(content, resRe)
		res.Value[field] = val
	}

}

//转换成orm;example model=OrmExample
func (res *ParseResult) ToOrm(model interface{}) interface{} {
	itemMap := res.Value
	if err := mapstructure.Decode(itemMap, &model); err != nil {
		log.Println(err)
	}
	return model
}
