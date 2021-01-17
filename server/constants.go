package server

import (
	"time"
)

const (
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
	robots    = "public/robots.txt"
	sitemap   = "public/sitemap.xml"
	favicon   = "public/images/icons/favicon-16x16.png"
	favicon32 = "public/images/icons/favicon-32x32.png"
)

// User cookie
const (
	sessionIDKey = "user"
	dayHrs       = 1 * 24 * time.Hour
	daySeconds   = int(dayHrs) * 3600
)
