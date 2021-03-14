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
	recordsDevicesHREF  = "/records-xhr-sources" // Relative template Records URLs
	recordsTypesHREF    = "/records-xhr-types"
	recordsAllHREF      = "/records-xhr-all"
	summaryTableHREF    = "/summary-table" // Relative template Summary URLs
	summaryActivityHREF = "/summary-activity"
	summaryBodyMassHREF = "/summary-mass"
	summaryExposureHREF = "/summary-exposure"
	summaryWalksHREF    = "/summary-walks"
	infoPrivacy         = "/info-privacy"
	infoContactUs       = "/info-contact-us"
	infoAboutUs         = "/info-about-us"
)

type RecordsDataCategory string

const (
	recordsSource   RecordsDataCategory = "source"
	recordsTypes    RecordsDataCategory = "types"
	recordsAll      RecordsDataCategory = "all"
	activitySummary RecordsDataCategory = "summary"
	bodyMass        RecordsDataCategory = "mass"
	exposure        RecordsDataCategory = "exposure"
	walks           RecordsDataCategory = "walks"
)

const (
	robots    = "public/robots.txt"
	sitemap   = "public/sitemap.xml"
	favicon   = "public/images/icons/favicon-16x16.png"
	favicon32 = "public/images/icons/favicon-32x32.png"
)

// User cookie
const (
	userCookieKey = "user" // "user" cookie key stored
	dayHrs        = 1 * 24 * time.Hour
	daySeconds    = int(dayHrs) * 3600
)

// User Messages in case of error
const (
	umNoError      = "Dont worry be happy."
	umNoData       = "Sorry could not find your data."
	umNotProcessed = "Sorry your data is not processed yet."
)

//  Context keys used in context request.
type contextKey int

const (
	contextKeyUserID contextKey = iota // User UUID is stored in this scoped request
)
