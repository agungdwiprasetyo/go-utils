package utils

import (
	"reflect"
)

func SliceCut(data interface{}, a, b int) []interface{} {
	ref := reflect.ValueOf(data)
	if ref.Kind() != reflect.Slice {
		return nil
	}
	if a > b || a > ref.Len() || b > ref.Len() {
		return nil
	}

	var res []interface{}
	for i := 0; i < ref.Len(); i++ {
		if i > a && i < b {
			continue
		}
		res = append(res, ref.Index(i).Interface())
	}
	return res
}

func SliceDelete(data interface{}, a int) []interface{} {
	ref := reflect.ValueOf(data)
	if ref.Kind() != reflect.Slice {
		return nil
	}

	var res []interface{}
	for i := 0; i < ref.Len(); i++ {
		if i == a {
			continue
		}
		res = append(res, ref.Index(i).Interface())
	}
	return res
}
