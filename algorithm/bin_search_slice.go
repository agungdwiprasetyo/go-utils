package algorithm

import (
	"fmt"
	"reflect"

	utils "github.com/agungdwiprasetyo/go-utils"
)

// SearchInSlice func
func SearchInSlice(slice, searchValue interface{}, fieldName string) int {
	var idx = -1

	ref := reflect.ValueOf(slice)
	if ref.Kind() == reflect.Ptr {
		ref = reflect.ValueOf(slice).Elem()
	}
	if ref.Kind() != reflect.Slice {
		return idx
	}

	n := ref.Len()
	first, last := 0, n-1
	for first <= last {
		mid := (first + last) / 2

		var valueInData int
		isExist := false
		for i := 0; i < ref.Index(mid).NumField(); i++ {
			key := ref.Index(mid).Type().Field(i).Name
			if key == fieldName {
				valueI := ref.Index(mid).Field(i).Interface()
				valueInData = utils.ParseInt(fmt.Sprint(valueI))
				isExist = true
				break
			}
		}
		if !isExist {
			return -1
		}

		inSearch := utils.ParseInt(fmt.Sprint(searchValue))
		if inSearch > valueInData {
			first = mid + 1
		} else if inSearch < valueInData {
			last = mid - 1
		} else {
			idx = mid
			break
		}
	}

	return idx
}
