package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"runtime/trace"

	"github.com/PumpkinSeed/structs/v2"
)

var profileBool bool
var traceBool bool

func init() {
	flag.BoolVar(&profileBool, "profile", false, "Start pprof")
	flag.BoolVar(&traceBool, "trace", false, "Start trace")
}

func main() {
	flag.Parse()

	if profileBool {
		pprof.StartCPUProfile(os.Stdout)
		defer pprof.StopCPUProfile()
	}

	if traceBool {
		trace.Start(os.Stdout)
		defer trace.Stop()
	}

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

	if !profileBool && !traceBool {
		fmt.Println("Start test...")
	}

	for n := 0; n < 1000000; n++ {
		structs.Contains(testStruct, float64(16.444))
	}

	if !profileBool && !traceBool {
		fmt.Println("Stop test...")
	}
}
