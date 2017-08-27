package structs

import "testing"

func TestContains(t *testing.T) {
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

	result := Contains(testStruct, float64(16.444))
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
