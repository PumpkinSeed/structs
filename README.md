# Golang structs
Package structs implements simple functions to manipulate structs in Golang.

## Contains
Contains check the struct is contain the given value or not.

```
package main

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

    result := Contains(tst, float64(16.444)) // true
    result = Contains(tst, float32(13.444)) // true
}
```

#### Benchmark

```
BenchmarkContains-4   	 3000000	       492 ns/op
```

## Compare
_Not yet implemented_
Compare two struct, two different compare, compare by fields and compare by fields and values

## Index
_Not yet implemented_
Return the index of the field where the second parameter value located

## Map
_Not yet implemented_
The second parameter is a function, apply the function on each field on the struct, or on the condition determined in the third argument

## Replace
_Not yet implemented_
Replace elements in the struct
