// Package structs implements simple functions to manipulate structs in Golang.
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
	complex64Type  = reflect.TypeOf(complex64(0))
	complex128Type = reflect.TypeOf(complex128(0))

	// add byte array type

	// add rune type

	boolType = reflect.TypeOf(true)

	errorType = reflect.TypeOf(errors.New(""))
)

// Contains reports whether value is within s
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
		case complex64Type:
			if containsComplex64(field.Interface().(complex64), value) {
				return true
			}
		case complex128Type:
			if containsComplex128(field.Interface().(complex128), value) {
				return true
			}
		}
	}
	return false
}

// Compare returns a boolean comparing two struct
func Compare(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// Index returns the index of the first instance of the value in s
func Index(s, value interface{}) int {
	v1 := reflect.ValueOf(s)
	v2 := reflect.ValueOf(value)

	for i := 0; i < v1.NumField(); i++ {
		f := v1.Field(i)
		if f.Type() == v2.Type() {
			switch f.Type() {
			case stringType:
				if f.String() == v2.String() {
					return i
				}
			case intType, int8Type, int16Type, int32Type, int64Type:
				if f.Int() == v2.Int() {
					return i
				}
			case boolType:
				if f.Bool() == v2.Bool() {
					return i
				}
			case float32Type:
				if f.Interface().(float32) == v2.Interface().(float32) {
					return i
				}
			case float64Type:
				if f.Float() == v2.Float() {
					return i
				}
			case complex64Type:
				if f.Interface().(complex64) == v2.Interface().(complex64) {
					return i
				}
			case complex128Type:
				if f.Complex() == v2.Complex() {
					return i
				}
			}
		}
	}
	return -1
}

// FieldNameByValue returns the field's name of the first instance of the value in s
func FieldNameByValue(s, value interface{}) string {
	v1 := reflect.ValueOf(s)
	t1 := reflect.TypeOf(s)
	v2 := reflect.ValueOf(value)

	for i := 0; i < v1.NumField(); i++ {
		f := v1.Field(i)
		if f.Type() == v2.Type() {
			switch f.Type() {
			case stringType:
				if f.String() == v2.String() {
					return t1.Field(i).Name
				}
			case intType, int8Type, int16Type, int32Type, int64Type:
				if f.Int() == v2.Int() {
					return t1.Field(i).Name
				}
			case boolType:
				if f.Bool() == v2.Bool() {
					return t1.Field(i).Name
				}
			case float32Type:
				if f.Interface().(float32) == v2.Interface().(float32) {
					return t1.Field(i).Name
				}
			case float64Type:
				if f.Float() == v2.Float() {
					return t1.Field(i).Name
				}
			case complex64Type:
				if f.Interface().(complex64) == v2.Interface().(complex64) {
					return t1.Field(i).Name
				}
			case complex128Type:
				if f.Complex() == v2.Complex() {
					return t1.Field(i).Name
				}
			}
		}
	}
	return ""
}

func Map(s interface{}, handler func(reflect.Value) error) (interface{}, error) {
	v := reflect.Indirect(reflect.ValueOf(s))

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.CanSet() {
			err := handler(f)
			if err != nil {
				return nil, err
			}
		}

	}
	return s, nil
}

// Replace returns a copy of the struct s with the first n non-overlapping instance of old replaced by new
func Replace(s, old, new interface{}, n int) (interface{}, error) {
	v1 := reflect.Indirect(reflect.ValueOf(s))
	//t1 := reflect.TypeOf(s)
	oldV := reflect.ValueOf(old)
	newV := reflect.ValueOf(new)

	if oldV.Type() != newV.Type() {
		return nil, errReplaceTypesNotMatching
	}

	var c = 0

	for i := 0; i < v1.NumField(); i++ {
		if n != -1 && c >= n {
			return s, nil
		}
		f := v1.Field(i)
		switch f.Type() {
		case stringType:
			if oldV.Type() == stringType &&
				f.String() == oldV.String() {
				f.SetString(newV.String())
				c++
			}
		case intType, int8Type, int16Type, int32Type, int64Type:
			if oldV.Type() == intType &&
				f.Int() == oldV.Int() {
				f.SetInt(newV.Int())
				c++
			}
		case boolType:
			if oldV.Type() == boolType &&
				f.Bool() == oldV.Bool() {
				f.SetBool(newV.Bool())
				c++
			}
		case float32Type, float64Type:
			if oldV.Type() == float32Type &&
				f.Float() == oldV.Float() {
				f.SetFloat(newV.Float())
				c++
			}
			// @todo add float64
		case complex64Type, complex128Type:
			if oldV.Type() == complex64Type &&
				f.Complex() == oldV.Complex() {
				f.SetComplex(newV.Complex())
				c++
			}
			// @todo add complex128
		}
	}
	return s, nil
}

/*
	Contains helper functions
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

func containsComplex64(s complex64, v interface{}) bool {
	switch v.(type) {
	case complex64:
		if v.(complex64) == s {
			return true
		}
	}
	return false
}

func containsComplex128(s complex128, v interface{}) bool {
	switch v.(type) {
	case complex128:
		if v.(complex128) == s {
			return true
		}
	}
	return false
}
