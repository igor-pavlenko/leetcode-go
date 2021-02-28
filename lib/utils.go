package lib

import (
	"fmt"
	"reflect"
)

func PrintList(head interface{}) string {
	result := ""
	for !reflect.ValueOf(head).IsNil() {
		v := reflect.ValueOf(head)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		val := v.FieldByName("Val").Int()
		result = fmt.Sprintf("%s->%d", result, val)
		nodePtr := v.FieldByName("Next")
		if nodePtr.Kind() == reflect.Ptr {
			if nodePtr.IsNil() {
				return result
			}
			head = nodePtr.Interface()
		} else {
			break
		}
	}
	return result
}
