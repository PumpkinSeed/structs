package structs

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	testStruct := struct {
		TestInt        int
		TestInt8       int8
		TestInt16      int16
		TestInt32      int32
		TestInt64      int64
		TestString     string
		TestBool       bool
		TestFloat32    float32
		TestFloat64    float64
		TestComplex64  complex64
		TestComplex128 complex128
	}{
		TestInt:        12,
		TestInt8:       42,
		TestInt16:      55,
		TestInt32:      33,
		TestInt64:      78,
		TestString:     "test",
		TestBool:       false,
		TestFloat32:    13.444,
		TestFloat64:    16.444,
		TestComplex64:  12333,
		TestComplex128: 123444455,
	}

	result := Contains(testStruct, complex64(12333))
	if !result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, complex128(123444455))
	if !result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, float64(16.444))
	if !result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, float32(13.444))
	if !result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, 12)
	if !result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, int8(42))
	if !result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, int16(55))
	if !result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, int32(33))
	if !result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, int64(78))
	if !result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, "test")
	if !result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, false)
	if !result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, float64(16.445))
	if result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, float32(14.444))
	if result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, 14)
	if result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, 44)
	if result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, 54)
	if result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, 34)
	if result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, 79)
	if result {
		t.Error("Should be true instead of false")
	}

	result = Contains(testStruct, "tests")
	if result {
		t.Error("Should be true instead of false")
	}
}

func TestCompare(t *testing.T) {
	testStructA := struct {
		TestInt   int
		TestInt8  int8
		TestInt16 int16
	}{
		TestInt:   12,
		TestInt8:  42,
		TestInt16: 55,
	}

	testStructB := struct {
		TestInt   int
		TestInt8  int8
		TestInt16 int16
	}{
		TestInt:   12,
		TestInt8:  42,
		TestInt16: 55,
	}

	if !Compare(testStructA, testStructB) {
		t.Errorf("Should return true")
	}

}

func TestIndex(t *testing.T) {
	testStruct := struct {
		TestInt        int
		TestInt8       int8
		TestInt16      int16
		TestInt32      int32
		TestInt64      int64
		TestString     string
		TestBool       bool
		TestFloat32    float32
		TestFloat64    float64
		TestComplex64  complex64
		TestComplex128 complex128
	}{
		TestInt:        12,
		TestInt8:       42,
		TestInt16:      55,
		TestInt32:      33,
		TestInt64:      78,
		TestString:     "test",
		TestBool:       false,
		TestFloat32:    13.444,
		TestFloat64:    16.444,
		TestComplex64:  12333,
		TestComplex128: 123444455,
	}

	value := Index(testStruct, "test")
	if value != 5 {
		t.Errorf("Position should be 5, instead of %d", value)
	}

	value = Index(testStruct, 12)
	if value != 0 {
		t.Errorf("Position should be 0, instead of %d", value)
	}

	value = Index(testStruct, false)
	if value != 6 {
		t.Errorf("Position should be 6, instead of %d", value)
	}

	value = Index(testStruct, float32(13.444))
	if value != 7 {
		t.Errorf("Position should be 7, instead of %d", value)
	}

	value = Index(testStruct, 16.444)
	if value != 8 {
		t.Errorf("Position should be 8, instead of %d", value)
	}

	value = Index(testStruct, complex64(12333))
	if value != 9 {
		t.Errorf("Position should be 9, instead of %d", value)
	}

	value = Index(testStruct, complex128(123444455))
	if value != 10 {
		t.Errorf("Position should be 10, instead of %d", value)
	}

	value = Index(testStruct, 42)
	if value != -1 {
		t.Errorf("Position should be -1, instead of %d", value)
	}
}

func TestFieldNameByValue(t *testing.T) {
	testStruct := struct {
		TestInt        int
		TestInt8       int8
		TestInt16      int16
		TestInt32      int32
		TestInt64      int64
		TestString     string
		TestBool       bool
		TestFloat32    float32
		TestFloat64    float64
		TestComplex64  complex64
		TestComplex128 complex128
	}{
		TestInt:        12,
		TestInt8:       42,
		TestInt16:      55,
		TestInt32:      33,
		TestInt64:      78,
		TestString:     "test",
		TestBool:       false,
		TestFloat32:    13.444,
		TestFloat64:    16.444,
		TestComplex64:  12333,
		TestComplex128: 123444455,
	}

	value := FieldNameByValue(testStruct, "test")
	if value != "TestString" {
		t.Errorf("Position should be 'TestString', instead of %s", value)
	}

	value = FieldNameByValue(testStruct, 12)
	if value != "TestInt" {
		t.Errorf("Position should be 'TestInt', instead of %s", value)
	}

	value = FieldNameByValue(testStruct, false)
	if value != "TestBool" {
		t.Errorf("Position should be 'TestBool', instead of %s", value)
	}

	value = FieldNameByValue(testStruct, float32(13.444))
	if value != "TestFloat32" {
		t.Errorf("Position should be 'TestFloat32', instead of %s", value)
	}

	value = FieldNameByValue(testStruct, 16.444)
	if value != "TestFloat64" {
		t.Errorf("Position should be 'TestFloat64', instead of %s", value)
	}

	value = FieldNameByValue(testStruct, complex64(12333))
	if value != "TestComplex64" {
		t.Errorf("Position should be 'TestComplex64', instead of %s", value)
	}

	value = FieldNameByValue(testStruct, complex128(123444455))
	if value != "TestComplex128" {
		t.Errorf("Position should be 'TestComplex128', instead of %s", value)
	}

	value = FieldNameByValue(testStruct, 42)
	if value != "" {
		t.Errorf("Position should be '', instead of %s", value)
	}
}

func TestReplace(t *testing.T) {
	type testStruct struct {
		TestInt        int
		TestInt8       int8
		TestInt16      int16
		TestInt32      int32
		TestInt64      int64
		TestString1    string
		TestString2    string
		TestString3    string
		TestString4    string
		TestBool       bool
		TestFloat32    float32
		TestFloat64    float64
		TestComplex64  complex64
		TestComplex128 complex128
	}
	ts := testStruct{
		TestInt:        12,
		TestInt8:       42,
		TestInt16:      55,
		TestInt32:      33,
		TestInt64:      78,
		TestString1:    "test",
		TestString2:    "test",
		TestString3:    "test",
		TestString4:    "test",
		TestBool:       false,
		TestFloat32:    13.444,
		TestFloat64:    16.444,
		TestComplex64:  12333,
		TestComplex128: 123444455,
	}

	value, err := Replace(&ts, "test", "new", 2)
	if err != nil {
		t.Errorf("Error should be nil, instead of %s", err.Error())
	}
	if value.(*testStruct).TestString1 != "new" {
		t.Errorf("TestString1 should be 'new', instead of %s", value.(*testStruct).TestString1)
	}
	if value.(*testStruct).TestString3 != "test" {
		t.Errorf("TestString1 should be 'test', instead of %s", value.(*testStruct).TestString3)
	}

	value, err = Replace(&ts, "test", "new", -1)
	if err != nil {
		t.Errorf("Error should be nil, instead of %s", err.Error())
	}
	if value.(*testStruct).TestString1 != "new" {
		t.Errorf("TestString1 should be 'new', instead of %s", value.(*testStruct).TestString1)
	}
	if value.(*testStruct).TestString3 != "new" {
		t.Errorf("TestString1 should be 'new', instead of %s", value.(*testStruct).TestString3)
	}

	value, err = Replace(&ts, 12, 42, -1)
	if err != nil {
		t.Errorf("Error should be nil, instead of %s", err.Error())
	}
	if value.(*testStruct).TestInt != 42 {
		t.Errorf("TestInt should be 42, instead of %d", value.(*testStruct).TestInt)
	}

	value, err = Replace(&ts, false, true, -1)
	if err != nil {
		t.Errorf("Error should be nil, instead of %s", err.Error())
	}
	if value.(*testStruct).TestBool != true {
		t.Errorf("TestInt should be true, instead of %t", value.(*testStruct).TestBool)
	}

	value, err = Replace(&ts, float32(13.444), float32(42.444), -1)
	if err != nil {
		t.Errorf("Error should be nil, instead of %s", err.Error())
	}
	if value.(*testStruct).TestFloat32 != 42.444 {
		t.Errorf("TestFloat32 should be 42.444, instead of %f", value.(*testStruct).TestFloat32)
	}

	value, err = Replace(&ts, complex64(12333), complex64(12334), -1)
	if err != nil {
		t.Errorf("Error should be nil, instead of %s", err.Error())
	}
	if value.(*testStruct).TestComplex64 != 12334 {
		t.Errorf("TestFloat32 should be 42.444, instead of %f", value.(*testStruct).TestComplex64)
	}

	value, err = Replace(&ts, complex64(12333), float32(12334), -1)
	if err != errReplaceTypesNotMatching {
		t.Errorf("Error should be %s, instead of nil", errReplaceTypesNotMatching.Error())
	}
}

func TestMap(t *testing.T) {
	type testStruct struct {
		Username string
		Title    string
		Content  string
	}
	ts := testStruct{
		Username: "PumpkinSeed",
		Title:    "Test title",
		Content:  "Test content",
	}
	res, err := Map(&ts, func(v reflect.Value) error {
		if v.Type() == stringType {
			v.SetString(strings.ToLower(v.String()))
		}
		return nil
	})
	if err != nil {
		t.Errorf("Error should be nil, instead of %s", err.Error())
	}
	if res.(*testStruct).Username != "pumpkinseed" {
		t.Errorf("Username should be 'pumpkinseed', instead of %s", res.(*testStruct).Username)
	}

	var testErr = errors.New("Test")
	res, err = Map(&ts, func(v reflect.Value) error {
		return testErr
	})
	if err != testErr {
		t.Errorf("Error should be %s, instead of nil", testErr.Error())
	}
}

/*
	Benchmarks
*/

func BenchmarkContains(b *testing.B) {
	testStruct := struct {
		TestInt     int
		TestInt8    int8
		TestInt16   int16
		TestInt32   int32
		TestInt64   int64
		TestString  string
		TestBool    bool
		TestFloat32 float32
		TestFloat64 float64
	}{
		TestInt:     12,
		TestInt8:    42,
		TestInt16:   55,
		TestInt32:   33,
		TestInt64:   78,
		TestString:  "test",
		TestBool:    false,
		TestFloat32: 13.444,
		TestFloat64: 16.444,
	}

	for n := 0; n < b.N; n++ {
		Contains(testStruct, float64(16.444))
	}
}

func BenchmarkCompareEqual(b *testing.B) {
	testStructA := struct {
		TestInt   int
		TestInt8  int8
		TestInt16 int16
	}{
		TestInt:   12,
		TestInt8:  42,
		TestInt16: 55,
	}

	testStructB := struct {
		TestInt   int
		TestInt8  int8
		TestInt16 int16
	}{
		TestInt:   12,
		TestInt8:  42,
		TestInt16: 55,
	}

	for n := 0; n < b.N; n++ {
		Compare(testStructA, testStructB)
	}
}

func BenchmarkCompareNotEqual(b *testing.B) {
	testStructA := struct {
		TestInt   int
		TestInt8  int8
		TestInt16 int16
	}{
		TestInt:   12,
		TestInt8:  42,
		TestInt16: 56,
	}

	testStructB := struct {
		TestInt   int
		TestInt8  int8
		TestInt16 int16
	}{
		TestInt:   12,
		TestInt8:  42,
		TestInt16: 55,
	}

	for n := 0; n < b.N; n++ {
		Compare(testStructA, testStructB)
	}
}

func BenchmarkIndex(b *testing.B) {
	testStruct := struct {
		TestInt     int
		TestInt8    int8
		TestInt16   int16
		TestInt32   int32
		TestInt64   int64
		TestString  string
		TestBool    bool
		TestFloat32 float32
		TestFloat64 float64
	}{
		TestInt:     12,
		TestInt8:    42,
		TestInt16:   55,
		TestInt32:   33,
		TestInt64:   78,
		TestString:  "test",
		TestBool:    false,
		TestFloat32: 13.444,
		TestFloat64: 16.444,
	}

	for n := 0; n < b.N; n++ {
		Index(testStruct, "test")
	}
}

func BenchmarkFieldNameByValue(b *testing.B) {
	testStruct := struct {
		TestInt     int
		TestInt8    int8
		TestInt16   int16
		TestInt32   int32
		TestInt64   int64
		TestString  string
		TestBool    bool
		TestFloat32 float32
		TestFloat64 float64
	}{
		TestInt:     12,
		TestInt8:    42,
		TestInt16:   55,
		TestInt32:   33,
		TestInt64:   78,
		TestString:  "test",
		TestBool:    false,
		TestFloat32: 13.444,
		TestFloat64: 16.444,
	}

	for n := 0; n < b.N; n++ {
		FieldNameByValue(testStruct, "test")
	}
}

func BenchmarkReplace(b *testing.B) {
	type testStruct struct {
		TestInt64      int64
		TestString1    string
		TestString2    string
		TestString3    string
		TestString4    string
		TestBool       bool
		TestFloat32    float32
		TestFloat64    float64
		TestComplex64  complex64
		TestComplex128 complex128
	}
	ts := testStruct{
		TestInt64:      78,
		TestString1:    "test",
		TestString2:    "test",
		TestString3:    "test",
		TestString4:    "test",
		TestBool:       false,
		TestFloat32:    13.444,
		TestFloat64:    16.444,
		TestComplex64:  12333,
		TestComplex128: 123444455,
	}
	for n := 0; n < b.N; n++ {
		Replace(&ts, "test", "new", 2)
	}
}

func BenchmarkMap(b *testing.B) {
	type testStruct struct {
		Username string
		Title    string
		Content  string
	}
	ts := testStruct{
		Username: "PumpkinSeed",
		Title:    "Test title",
		Content:  "Test content",
	}
	for n := 0; n < b.N; n++ {
		Map(&ts, func(v reflect.Value) error {
			if v.Type() == stringType {
				v.SetString(strings.ToLower(v.String()))
			}
			return nil
		})
	}
}
