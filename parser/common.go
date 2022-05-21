package parser

import (
	"reflect"
	"regexp"
)

//根据正则进行匹配单个字段
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}


func ConvertStructToMap(i interface{})  map[string]interface{}{
	if _,ok :=i.(struct{}); !ok {
		return nil
	}
	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}