// Package main sample test utility to process HK data
package main

import (
	"flag"
	"github.com/opendroid/hk/server"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

const (
	_ = "./export/sampleexport.xml" // Example to test
)

// go tool pprof -http=:8080 /Users/ajayt/Downloads/mem.prof
// Example from https://golang.org/pkg/runtime/pprof/

// cpuProfile name of output CPU Profile file name
var (
	cpuProfile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	memProfile = flag.String("memprofile", "", "write memory profile to `file`")
	file       = flag.String("file", "public/sampleexport.xml", "Apple health kit exported data")
	_          = flag.String("record", "StepCount", "export a record csv to a `file`")
	port       = flag.Int("port", 8080, "Http server port")
)

// main enter the program here
func main() {
	flag.Parse()
	if *cpuProfile != "" { // Save CPU  profile
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer func() { _ = f.Close() }() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// Save memory profile
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer func() { _ = f.Close() }() // error handling omitted for example
		runtime.GC()                     // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
	server.Start(*port, *file) // Start the http server
}
