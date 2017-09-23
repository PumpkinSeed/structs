package structs

import "testing"

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

	result := Contains(testStruct, float64(16.444))
	if !result {
		t.Error("Should be true instead of false")
	}

}

/*
	Benchmark
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
