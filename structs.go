package structs

import (
	"errors"
	"reflect"
)

var (
	float32Type = reflect.TypeOf(float32(0))
	float64Type = reflect.TypeOf(float64(0))

	stringType = reflect.TypeOf("")

	intType   = reflect.TypeOf(int(0))
	int8Type  = reflect.TypeOf(int8(0))
	int16Type = reflect.TypeOf(int16(0))
	int32Type = reflect.TypeOf(int32(0))
	int64Type = reflect.TypeOf(int64(0))

	uintType   = reflect.TypeOf(uint(0))
	uint8Type  = reflect.TypeOf(uint8(0))
	uint16Type = reflect.TypeOf(uint16(0))
	uint32Type = reflect.TypeOf(uint32(0))
	uint64Type = reflect.TypeOf(uint64(0))
	// add uintptr

	// add complex types

	// add byte array type

	// add rune type

	boolType = reflect.TypeOf(true)

	errorType = reflect.TypeOf(errors.New(""))
)

func Contains(s, value interface{}) bool {
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		typeOfField := field.Type()
		switch typeOfField {
		case stringType:
			if containsString(field.String(), value) {
				return true
			}
		case intType, int8Type, int16Type, int32Type, int64Type:
			if containsInt(field.Int(), value) {
				return true
			}
		case boolType:
			if containsBool(field.Bool(), value) {
				return true
			}
		case float32Type:
			if containsFloat32(field.Interface().(float32), value) {
				return true
			}
		case float64Type:
			if containsFloat64(field.Interface().(float64), value) {
				return true
			}
		}
	}
	return false
}

/*
func Compare(a, b interface{}) bool {
	return false
}

func Index(s, value interface{}) string {
	return ""
}

func Map(s interface{}, f func(interface{}) error) error {
	return nil
}

func Replace(s, old, new interface{}) error {
	return nil
}
*/
func containsString(s string, v interface{}) bool {
	switch v.(type) {
	case string:
		if v.(string) == s {
			return true
		}
	}
	return false
}

func containsInt(s int64, v interface{}) bool {
	switch v.(type) {
	case int64:
		if v.(int64) == s {
			return true
		}
	case int32:
		if int64(v.(int32)) == s {
			return true
		}
	case int16:
		if int64(v.(int16)) == s {
			return true
		}
	case int8:
		if int64(v.(int8)) == s {
			return true
		}
	case int:
		if int64(v.(int)) == s {
			return true
		}
	}
	return false
}

func containsBool(s bool, v interface{}) bool {
	switch v.(type) {
	case bool:
		if v.(bool) == s {
			return true
		}
	}
	return false
}

func containsFloat32(s float32, v interface{}) bool {
	switch v.(type) {
	case float32:
		if v.(float32) == s {
			return true
		}
	}
	return false
}

func containsFloat64(s float64, v interface{}) bool {
	switch v.(type) {
	case float64:
		if v.(float64) == s {
			return true
		}
	}
	return false
}
