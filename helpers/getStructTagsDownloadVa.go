package helpers

import "reflect"

func GetStructTags(data interface{}) []string {
	v := reflect.TypeOf(data)
	headers := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		headers[i] = v.Field(i).Tag.Get("xlsx")
	}
	return headers
}
