package helpers

import "reflect"

// TODO: investigate https://stackoverflow.com/questions/61217817/why-does-reflecting-the-name-or-package-path-of-the-error-type-cause-a-panic-i
func GetType(v interface{}) string {
	if v == nil {
		return "nil"
	}
	return reflect.TypeOf(v).String()
}
