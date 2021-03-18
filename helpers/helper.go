package helpers

import "reflect"

func GetType(v interface{}) string {
	return reflect.TypeOf(v).String()
}
