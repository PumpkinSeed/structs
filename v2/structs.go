package structs

import (
	"errors"
	"reflect"
)

/*
	-------------------------------------------
	Try to optimize things with Go concurrency,
	but as seen it's worse. Bench:
	BenchmarkContains-4 300000	     4012 ns/op
	-------------------------------------------
*/

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
	check := make(chan bool, 1)
	go containsMain(check, v, value)
	for {
		select {
		case c := <-check:
			return c
		}
	}
}

/*
	Contains helper functions
*/

func containsMain(c chan bool, v reflect.Value, value interface{}) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		typeOfField := field.Type()
		switch typeOfField {
		case stringType:
			go containsString(c, field.String(), value)
		case intType, int8Type, int16Type, int32Type, int64Type:
			go containsInt(c, field.Int(), value)
		case boolType:
			go containsBool(c, field.Bool(), value)
		case float32Type:
			go containsFloat32(c, field.Interface().(float32), value)
		case float64Type:
			go containsFloat64(c, field.Interface().(float64), value)
		case complex64Type:
			go containsComplex64(c, field.Interface().(complex64), value)
		case complex128Type:
			go containsComplex128(c, field.Interface().(complex128), value)
		}
	}
}

func containsString(c chan bool, s string, v interface{}) {
	switch v.(type) {
	case string:
		if v.(string) == s {
			c <- true
		}
	}
	return
}

func containsInt(c chan bool, s int64, v interface{}) {
	switch v.(type) {
	case int64:
		if v.(int64) == s {
			c <- true
		}
	case int32:
		if int64(v.(int32)) == s {
			c <- true
		}
	case int16:
		if int64(v.(int16)) == s {
			c <- true
		}
	case int8:
		if int64(v.(int8)) == s {
			c <- true
		}
	case int:
		if int64(v.(int)) == s {
			c <- true
		}
	}
	return
}

func containsBool(c chan bool, s bool, v interface{}) {
	switch v.(type) {
	case bool:
		if v.(bool) == s {
			c <- true
		}
	}
	return
}

func containsFloat32(c chan bool, s float32, v interface{}) {
	switch v.(type) {
	case float32:
		if v.(float32) == s {
			c <- true
		}
	}
	return
}

func containsFloat64(c chan bool, s float64, v interface{}) {
	switch v.(type) {
	case float64:
		if v.(float64) == s {
			c <- true
		}
	}
	return
}

func containsComplex64(c chan bool, s complex64, v interface{}) {
	switch v.(type) {
	case complex64:
		if v.(complex64) == s {
			c <- true
		}
	}
	return
}

func containsComplex128(c chan bool, s complex128, v interface{}) {
	switch v.(type) {
	case complex128:
		if v.(complex128) == s {
			c <- true
		}
	}
	return
}
