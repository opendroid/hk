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
	mux.HandleFunc("/records-xhr-sources", execTemplate(recordsSourcesNav, "records-xhr-sources.gohtml"))
	mux.HandleFunc("/records-xhr-types", execTemplate(recordsTypesNav, "records-xhr-types.gohtml"))
	mux.HandleFunc("/records-xhr-all", execTemplate(recordsAllNav, "records-xhr-all.gohtml"))
	mux.HandleFunc("/summary-table", execTemplate(summaryTableNav, "summary-table.gohtml"))
	mux.HandleFunc("/summary-activity", execTemplate(summaryGraphNav, "summary-activity.gohtml"))
	mux.HandleFunc("/summary-mass", execTemplate(summaryBodyMassNav, "summary-mass.gohtml"))
	mux.HandleFunc("/summary-exposure", execTemplate(summaryExposureNav, "summary-exposure.gohtml"))
	mux.HandleFunc("/summary-walks", execTemplate(summaryWalksNav, "summary-walks.gohtml"))
	mux.HandleFunc("/error", errorsPage)

	// v1 records APIs. The input "user" is a HTTP cookie
	mux.HandleFunc("/v1/recordsSources", recordsData(recordsSource))
	mux.HandleFunc("/v1/recordsTypes", recordsData(recordsTypes))
	mux.HandleFunc("/v1/recordsAll", recordsData(recordsAll))
	mux.HandleFunc("/v1/activitySummary", recordsData(activitySummary))
	mux.HandleFunc("/v1/bodyMass", recordsData(bodyMass))
	mux.HandleFunc("/v1/exposure", recordsData(exposure))
	mux.HandleFunc("/v1/walks", recordsData(walks))
	n := NewSessionHandler(mux) // Wrapper mux
	address := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(address, n))
}
