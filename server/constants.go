package server

import (
	"time"
)

const (
	root      = "root"          // Root template name
	css       = "public/css"    // CSS directory
	images    = "public/images" // images directory
	js        = "public/js"     // javascript directory
	templates = "public/templates"
	htmlGlob  = templates + "/*.gohtml"
)

const (
	cssAbs       = "/" + css
	imagesAbs    = "/" + images
	templatesAbs = "/" + templates
	jsAbs        = "/" + js
)

const (
	recordsDevicesHREF = "/records-xhr-sources" // Relative template Records URLs
	recordsTypesHREF   = "/records-xhr-types"
	recordsAllHREF     = "/records-xhr-all"
	summaryTableHREF   = "/summary-table" // Relative template Summary URLs
	summaryGraphHREF   = "/summary-graph"
)

type RecordsDataCategory string

const (
	recordsSource   RecordsDataCategory = "source"
	recordsTypes    RecordsDataCategory = "types"
	recordsAll      RecordsDataCategory = "all"
	activitySummary RecordsDataCategory = "summary"
)

const (
	robots    = "public/robots.txt"
	sitemap   = "public/sitemap.xml"
	favicon   = "public/images/icons/favicon-16x16.png"
	favicon32 = "public/images/icons/favicon-32x32.png"
)

// User cookie
const (
	sessionIDKey = "user" // User UUID is stored in this scoped request
	dayHrs       = 1 * 24 * time.Hour
	daySeconds   = int(dayHrs) * 3600
)

// User Messages in case of error
const (
	umNoError      = "Dont worry be happy"
	umNoData       = "Sorry could not find your data"
	umNotProcessed = "Sorry your data is not processed yet"
)
