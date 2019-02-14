package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var mapDate = map[string]string{
	"yyyy-mm-ddThh:mm:ssZ": time.RFC3339, // default
	"yyyy-mm-dd":           "2006-01-02",
}

// Validator common model/domain validator
type Validator struct {
	errs   *MultiError
	prefix string
}

// NewValidator init validator with prefix error
func NewValidator(prefix ...string) *Validator {
	validator := new(Validator)
	validator.errs = NewMultiError()
	validator.prefix = strings.Join(prefix, ".")
	return validator
}

// IsAlphanumeric function for check valid alphanumeric, string must contains alphabet & numeric format
func (v *Validator) IsAlphanumeric(str string) bool {
	var uppercase, lowercase, num int
	for _, r := range str {
		if r >= 65 && r <= 90 { //code ascii for [A-Z]
			uppercase++
		} else if r >= 97 && r <= 122 { //code ascii for [a-z]
			lowercase++
		} else if r >= 48 && r <= 57 { //code ascii for [0-9]
			num++
		} else {
			return false
		}
	}

	return (uppercase >= 1 || lowercase >= 1) && num >= 1
}

// IsNumeric function for check valid numeric
func (v *Validator) IsNumeric(str string) bool {
	for _, r := range str {
		if r < 48 || r > 57 {
			return false
		}
	}

	return str != ""
}

// IsAlphabet function for check valid numeric
func (v *Validator) IsAlphabet(str string) bool {
	var uppercase, lowercase int
	for _, r := range str {
		if r >= 65 && r <= 90 { //code ascii for [A-Z]
			uppercase++
		} else if r >= 97 && r <= 122 { //code ascii for [a-z]
			lowercase++
		} else { // except alphabet (symbol)
			return false
		}
	}

	return uppercase >= 1 || lowercase >= 1
}

// IsStringInSlice function for check valid numeric
func (v *Validator) IsStringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if str == s {
			return true
		}
	}

	return false
}

func (v *Validator) processSlice(fields reflect.Value, prefix string) {
	// iterate array of objects
	for i := 0; i < fields.Len(); i++ {
		v.validate(fields.Index(i), fmt.Sprintf("%s[%d]", prefix, i))
	}
}

func (v *Validator) validate(obj reflect.Value, prefix string) {
	// iterate fields from object
	for i := 0; i < obj.NumField(); i++ {
		field := obj.Field(i)
		fieldValue := field.Interface()
		fieldType := field.Kind()

		requiredTag := obj.Type().Field(i).Tag.Get("required")
		formatTag := obj.Type().Field(i).Tag.Get("format")
		maxLengthTag := obj.Type().Field(i).Tag.Get("maxlength")
		minLengthTag := obj.Type().Field(i).Tag.Get("minlength")
		inTag := obj.Type().Field(i).Tag.Get("in")
		jsonTag := strings.Replace(obj.Type().Field(i).Tag.Get("json"), ",omitempty", "", -1)
		if prefix != "" {
			jsonTag = fmt.Sprintf("%s.%s", prefix, jsonTag)
		}

		isRequire, _ := strconv.ParseBool(requiredTag)
		isEmptyValue := reflect.DeepEqual(fieldValue, reflect.Zero(reflect.TypeOf(fieldValue)).Interface())
		propertyLevel := "field"

		// get element if field is pointer
		if fieldType == reflect.Ptr {
			if field.IsNil() {
				field = reflect.ValueOf(&fieldValue).Elem()
				field = reflect.New(field.Elem().Type().Elem()) // create from new domain model type of field
			}
			field = field.Elem()
			fieldValue = field.Interface()
			fieldType = field.Kind()
		}

		switch fieldType {
		case reflect.Struct:
			if !isEmptyValue {
				v.validate(field, jsonTag)
			}
			propertyLevel = fmt.Sprintf("object %s{}", jsonTag)
		case reflect.Slice:
			v.processSlice(field, jsonTag)
			propertyLevel = fmt.Sprintf("array of object %s[{}]", jsonTag)
		}

		if isRequire && isEmptyValue {
			v.errs.Append(jsonTag, fmt.Errorf("%s is mandatory", propertyLevel))
		}

		fieldValueStr := fmt.Sprint(fieldValue)
		if !isEmptyValue && formatTag != "" {
			isValid := false

			switch {
			case formatTag == "alphanumeric":
				isValid = v.IsAlphanumeric(fieldValueStr)
			case formatTag == "numeric":
				isValid = v.IsNumeric(fieldValueStr)
			case formatTag == "alphabet":
				isValid = v.IsAlphabet(fieldValueStr)
			case strings.Contains(formatTag, "date"):
				dates := strings.Split(formatTag, ",")
				formatTime := "yyyy-mm-ddThh:mm:ssZ"
				if len(dates) > 1 {
					formatTime = strings.ToLower(dates[1])
				}

				if f, ok := mapDate[formatTime]; ok {
					_, err := time.Parse(f, fieldValueStr)
					isValid = (err == nil)
					formatTag = fmt.Sprintf("'%s: %s'", dates[0], formatTime)
				} else {
					panic(fmt.Errorf("invalid datetime format (%s) in field %s", formatTime, jsonTag))
				}
			default:
				panic(fmt.Errorf("undefined format '%s' in %s '%s'", formatTag, propertyLevel, jsonTag))
			}

			if !isValid {
				v.errs.Append(jsonTag, fmt.Errorf("field must in %s format", formatTag))
			}
		}

		// maxlength & minlength valid only if field is string
		if fieldType == reflect.String {
			maxLength, err := strconv.Atoi(maxLengthTag)
			if err == nil && len(fieldValueStr) > maxLength {
				v.errs.Append(jsonTag, fmt.Errorf("maximum length is %d", maxLength))
			}
			minLength, err := strconv.Atoi(minLengthTag)
			if err == nil && len(fieldValueStr) < minLength {
				v.errs.Append(jsonTag, fmt.Errorf("minimum length is %d", minLength))
			}
		}

		if !isEmptyValue && inTag != "" && !v.IsStringInSlice(fieldValueStr, strings.Split(inTag, ",")) {
			v.errs.Append(jsonTag, fmt.Errorf("field must in [ %s ]", inTag))
		}
	}
}

// Validate for all domain/model validator
func (v *Validator) Validate(data interface{}) (multiError *MultiError) {
	multiError = v.errs
	defer func() {
		if r := recover(); r != nil {
			v.errs.Clear()
			v.errs.Append("panic", fmt.Errorf("%v", r))
		}
	}()

	refValue := reflect.ValueOf(data)
	if refValue.Kind() == reflect.Ptr {
		refValue = refValue.Elem()
	}

	switch refValue.Kind() {
	case reflect.Slice:
		v.processSlice(refValue, refValue.Type().String())
	case reflect.Struct:
		v.validate(refValue, v.prefix)
	}

	if !v.errs.IsNil() {
		return v.errs
	}

	return nil
}
