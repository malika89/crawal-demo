package utils

import "reflect"

func ConvertStructToMap(i interface{}) map[string]interface{} {
	if _, ok := i.(struct{}); !ok {
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
