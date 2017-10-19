# Golang structs
Package structs implements simple functions to manipulate structs in Golang.

[![Documentation](https://godoc.org/github.com/PumpkinSeed/structs?status.svg)](https://godoc.org/github.com/PumpkinSeed/structs)


## Get it

```
go get github.com/PumpkinSeed/structs
```

## Contains
Contains reports whether value is within struct

```
package main

import "github.com/PumpkinSeed/structs"

type Tst struct {
    TestString  string
    TestFloat32 float32
    TestFloat64 float64
}

func main() {
    tst := Tst{
        TestString:  "test",
        TestFloat32: 13.444,
        TestFloat64: 16.444,
    }

    result := structs.Contains(tst, float64(16.444)) // true
    result = structs.Contains(tst, float32(13.444)) // true
}
```

#### Benchmark

```
BenchmarkContains-4   	 3000000	       492 ns/op
```

## Compare
Compare returns a boolean comparing two struct

```
package main

import "github.com/PumpkinSeed/structs"

type TstA struct {
	TestInt   int
	TestInt8  int8
	TestInt16 int16
}

type TstB struct {
	TestInt   int
	TestInt8  int8
	TestInt16 int16
}

func main() {
    tstA := TstA{
        TestInt:   12,
	    TestInt8:  42,
	    TestInt16: 55,
    }

    tstB := TstB{
        TestInt:   12,
	    TestInt8:  42,
	    TestInt16: 55,
    }

    result := structs.Compare(testStructA, testStructB) // true
}


```

#### Benchmark

```
BenchmarkCompareEqual-4      	 5000000	       379 ns/op
BenchmarkCompareNotEqual-4   	 5000000	       372 ns/op
```

## Index
Index returns the index of the first instance of the value in struct

```
package main

import "github.com/PumpkinSeed/structs"

type Tst struct {
    TestInt     int
	TestInt8    int8
	TestInt16   int16
	TestInt32   int32
	TestInt64   int64
	TestString  string
	TestBool    bool
	TestFloat32 float32
	TestFloat64 float64
}

func main() {
    tst := Tst{
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

    result := Index(testStruct, "test") // 5
}
```

#### Benchmark

```
BenchmarkIndex-4             	 5000000	       242 ns/op
```

## FieldNameByValue
FieldNameByValue returns the field's name of the first instance of the value in struct

```
package main

import "github.com/PumpkinSeed/structs"

type Tst struct {
    TestInt     int
	TestInt8    int8
	TestInt16   int16
	TestInt32   int32
	TestInt64   int64
	TestString  string
	TestBool    bool
	TestFloat32 float32
	TestFloat64 float64
}

func main() {
    tst := Tst{
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

    result := FieldNameByValue(testStruct, "test") // TestString
}
```

#### Benchmark

```
BenchmarkFieldNameByValue-4   	 5000000	       293 ns/op
```

## Map
_Not yet implemented_
The second parameter is a function, apply the function on each field on the struct, or on the condition determined in the third argument

## Replace
Replace returns a copy of the struct with the first non-overlapping instance of old replaced by new, the last param (n) is the limit, if n < 0, there is no limit on the number of replacements

```
package main

import "github.com/PumpkinSeed/structs"

type Tst struct {
    TestInt     int
	TestInt8    int8
	TestInt16   int16
	TestInt32   int32
	TestInt64   int64
	TestString1    string
	TestString2    string
	TestString3    string
	TestString4    string
	TestBool    bool
	TestFloat32 float32
	TestFloat64 float64
}

func main() {
    tst := Tst{
        TestInt:     12,
		TestInt8:    42,
		TestInt16:   55,
		TestInt32:   33,
		TestInt64:   78,
		TestString1:    "test",
		TestString2:    "test",
		TestString3:    "test",
		TestString4:    "test",
		TestBool:    false,
		TestFloat32: 13.444,
		TestFloat64: 16.444,
    }

    result, err := Replace(&ts, "test", "new", 2)
}
```

#### Benchmark

```
BenchmarkReplace-4            	 2000000	       655 ns/op
```
