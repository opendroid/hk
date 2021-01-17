// Package server hosts the HealthKit data on http port
package server

import (
	"fmt"
	"log"
	"net/http"
)

// Start the http server
func Start(port int, f string) {
	file = f
	mux := http.NewServeMux()

	// Static content handlers
	css := http.FileServer(http.Dir(css))
	images := http.FileServer(http.Dir(images))
	js := http.FileServer(http.Dir(js))
	mux.Handle(cssAbs+"/", http.StripPrefix(cssAbs, css))
	mux.Handle(imagesAbs+"/", http.StripPrefix(imagesAbs, images))
	mux.Handle(jsAbs+"/", http.StripPrefix(jsAbs, js))

	// URL handlers
	mux.HandleFunc("/", index)
	mux.HandleFunc("/records", records)
	n := NewSessionHandler(mux) // Wrapper mux
	address := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(address, n))
}
