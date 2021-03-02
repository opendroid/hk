// Package main sample test utility to process HK data
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/opendroid/hk/export"
	"github.com/opendroid/hk/server"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
)

const (
	_ = "./export/sampleexport.xml" // Example to test
)

// go tool pprof -http=:8080 /Downloads/mem.prof
// Example from https://golang.org/pkg/runtime/pprof/

// cpuProfile name of output CPU Profile file name
var (
	cpuProfile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	memProfile = flag.String("memprofile", "", "write memory profile to `file`")
	file       = flag.String("file", "public/sampleexport.xml", "Apple health kit exported data")
	out        = flag.String("out", "/Downloads/hk/ds/steps.csv", "Export steps data")
	tag        = flag.String("tag", "All", "Tag type")
	date       = flag.String("date", "2020-01-03", "Export data for date")
	port       = flag.Int("port", 8080, "Http server port")
	start      = flag.Bool("start", true, "Export to CSV only, dont start the server")
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
	_ = saveSteps // Suppress warning
	if *start {
		server.Start(*port, *file) // Start the http server
	}
}

// saveSteps in a CSV file
func saveSteps() {
	fmt.Printf("Starting processing ...\n")
	if health, err := export.Parse(*file); err != nil {
		fmt.Printf("Error parsing %q, %v\n", *file, err)
	} else {
		fmt.Printf("Processing finished. Saving file\n")
		f, err := os.Create(*out)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		defer func() { _ = f.Close() }()
		cw := csv.NewWriter(f)
		for i, r := range health.GetRecords(export.RecordType(*tag), *date) {
			row := []string{strconv.Itoa(i + 1), r.Type, r.SourceName, r.Device, r.SourceVersion, r.CreationDate,
				r.StartDate, r.EndDate, r.Unit, r.Value}
			if err := cw.Write(row); err != nil {
				fmt.Printf("Error writing CSV: %v", err)
				continue
			}
		}
		cw.Flush()
		_ = f.Sync()
		fmt.Printf("Saved file\n")

	}
}
