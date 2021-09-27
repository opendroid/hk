package server

import (
	"github.com/opendroid/hk/export"
	"github.com/opendroid/hk/logger"
	"go.uber.org/zap"
)

// NameActiveHREF a tuple that saves name and is a section is active.
type NameActiveHREF struct {
	Name   string
	Active bool
	HREF   string
}

// NavPrimary data used to diplay "navprimary" elements
type NavPrimary struct {
	Active string // Name of active CSS class in primary navigation links
	State  string // Show "login" or "data" <if available>
}

// PageHeader defines the Title of a page and Name of  Active Primary Nav page. One of "records", "activity" or "summary"
type PageHeader struct {
	Title string
	NP    NavPrimary
	NS    []NameActiveHREF // Nav secondary
}

// IndexPage data
type IndexPage struct {
	PageHeader
	UserID string
}

// ErrorPage data for Page records
type ErrorPage struct {
	PageHeader
	Error string
}

// InfoPage data for Page records
type InfoPage struct {
	PageHeader
}

// RecordsPage data for Page records
type RecordsPage struct {
	PageHeader
	*export.Me
	Records export.NameTypeKeyCounts
	Devices export.KeyCounts
	Types   export.KeyCounts
}

// Default data for template pages
var (
	loginNav = PageHeader{ // For record-xhr-sources.gohtml page
		Title: "Login With Google",
		NP:    NavPrimary{Active: "login", State: "login"},
	}

	recordsSourcesNav = PageHeader{ // For record-xhr-sources.gohtml page
		Title: "Your Health Records Sources",
		NP:    NavPrimary{Active: "records", State: "data"},
		NS: []NameActiveHREF{
			{Name: "Sources", Active: true, HREF: recordsDevicesHREF},
			{Name: "Types", Active: false, HREF: recordsTypesHREF},
			{Name: "All", Active: false, HREF: recordsAllHREF},
		},
	}

	recordsTypesNav = PageHeader{ // For record-xhr-types.gohtml page
		Title: "Your Health Records Types",
		NP:    NavPrimary{Active: "records", State: "data"},
		NS: []NameActiveHREF{
			{Name: "Sources", Active: false, HREF: recordsDevicesHREF},
			{Name: "Types", Active: true, HREF: recordsTypesHREF},
			{Name: "All", Active: false, HREF: recordsAllHREF},
		},
	}

	recordsAllNav = PageHeader{ // For record-xhr-all.gohtml page
		Title: "All Your Health Records",
		NP:    NavPrimary{Active: "records", State: "data"},
		NS: []NameActiveHREF{
			{Name: "Sources", Active: false, HREF: recordsDevicesHREF},
			{Name: "Types", Active: false, HREF: recordsTypesHREF},
			{Name: "All", Active: true, HREF: recordsAllHREF},
		},
	}

	// Summary health data section pages primary and secondary navigation data
	summaryTableNav = PageHeader{ // For summary.gohtml page
		Title: "Activity Summary Table",
		NP:    NavPrimary{Active: "summary", State: "data"},
		NS: []NameActiveHREF{
			{Name: "Table", Active: true, HREF: summaryTableHREF},
			{Name: "Activity", Active: false, HREF: summaryActivityHREF},
			{Name: "Body Mass", Active: false, HREF: summaryBodyMassHREF},
			{Name: "Exposure", Active: false, HREF: summaryExposureHREF},
			{Name: "Walks", Active: false, HREF: summaryWalksHREF},
		},
	}

	summaryGraphNav = PageHeader{ // For summary.gohtml page
		Title: "Your Activity Graph",
		NP:    NavPrimary{Active: "summary", State: "data"},
		NS: []NameActiveHREF{
			{Name: "Table", Active: false, HREF: summaryTableHREF},
			{Name: "Activity", Active: true, HREF: summaryActivityHREF},
			{Name: "Body Mass", Active: false, HREF: summaryBodyMassHREF},
			{Name: "Exposure", Active: false, HREF: summaryExposureHREF},
			{Name: "Walks", Active: false, HREF: summaryWalksHREF},
		},
	}

	summaryBodyMassNav = PageHeader{ // For summary.gohtml page
		Title: "Body Mass Data",
		NP:    NavPrimary{Active: "summary", State: "data"},
		NS: []NameActiveHREF{
			{Name: "Table", Active: false, HREF: summaryTableHREF},
			{Name: "Activity", Active: false, HREF: summaryActivityHREF},
			{Name: "Body Mass", Active: true, HREF: summaryBodyMassHREF},
			{Name: "Exposure", Active: false, HREF: summaryExposureHREF},
			{Name: "Walks", Active: false, HREF: summaryWalksHREF},
		},
	}

	summaryExposureNav = PageHeader{ // For summary.gohtml page
		Title: "Body Mass Data",
		NP:    NavPrimary{Active: "summary", State: "data"},
		NS: []NameActiveHREF{
			{Name: "Table", Active: false, HREF: summaryTableHREF},
			{Name: "Activity", Active: false, HREF: summaryActivityHREF},
			{Name: "Body Mass", Active: false, HREF: summaryBodyMassHREF},
			{Name: "Exposure", Active: true, HREF: summaryExposureHREF},
			{Name: "Walks", Active: false, HREF: summaryWalksHREF},
		},
	}

	summaryWalksNav = PageHeader{ // For summary.gohtml page
		Title: "Body Mass Data",
		NP:    NavPrimary{Active: "summary", State: "data"},
		NS: []NameActiveHREF{
			{Name: "Table", Active: false, HREF: summaryTableHREF},
			{Name: "Activity", Active: false, HREF: summaryActivityHREF},
			{Name: "Body Mass", Active: false, HREF: summaryBodyMassHREF},
			{Name: "Exposure", Active: false, HREF: summaryExposureHREF},
			{Name: "Walks", Active: true, HREF: summaryWalksHREF},
		},
	}

	// Info section pages primary and secondary navigation data
	infoPrivacyNav = PageHeader{
		Title: "Usense.io Privacy Policy",
		NP:    NavPrimary{Active: "info", State: "data"},
		NS: []NameActiveHREF{
			{Name: "Privacy", Active: true, HREF: infoPrivacy},
			{Name: "Contact", Active: false, HREF: infoContactUs},
			{Name: "About", Active: false, HREF: infoAboutUs},
		},
	}

	infoContactUsNav = PageHeader{
		Title: "Contact Us",
		NP:    NavPrimary{Active: "info", State: "data"},
		NS: []NameActiveHREF{
			{Name: "Privacy", Active: false, HREF: infoPrivacy},
			{Name: "Contact", Active: true, HREF: infoContactUs},
			{Name: "About", Active: false, HREF: infoAboutUs},
		},
	}

	infoAboutUsNav = PageHeader{
		Title: "About US",
		NP:    NavPrimary{Active: "info", State: "data"},
		NS: []NameActiveHREF{
			{Name: "Privacy", Active: false, HREF: infoPrivacy},
			{Name: "Contact", Active: false, HREF: infoContactUs},
			{Name: "About", Active: true, HREF: infoAboutUs},
		},
	}
	defaultSecondaryNav = []NameActiveHREF{ // Default
		{Name: "Sources", Active: false, HREF: recordsDevicesHREF},
		{Name: "Types", Active: false, HREF: recordsTypesHREF},
		{Name: "All", Active: false, HREF: recordsAllHREF},
		{Name: "Exposure", Active: false, HREF: summaryExposureHREF},
		{Name: "Walks", Active: false, HREF: summaryWalksHREF},
	}
)

// getRecordsData gets template data for page "Records - data"
func getIndexPageData(userID string) *IndexPage {
	return &IndexPage{
		PageHeader: PageHeader{
			Title: "Your Health Records",
			NP:    NavPrimary{Active: "records", State: "data"},
			NS:    defaultSecondaryNav,
		},
		UserID: userID,
	}
}

// getRecordsData gets template data for page "Records - data"
func getRecordsPageData(userID string) *RecordsPage {
	d := RecordsPage{
		PageHeader: PageHeader{
			Title: "Your records data",
			NP:    NavPrimary{Active: "records", State: "data"},
			NS:    defaultSecondaryNav,
		},
	}
	u, ok := users[userID]
	if !ok {
		logger.Error("No user data", zap.String("user", userID))
		return nil
	}
	d.Records = u.summary
	d.Devices = u.sources
	d.Types = u.recordTypes
	return &d
}

// getErrorPageData gets template data for page "Records - data"
func getErrorPageData(err string) *ErrorPage {
	return &ErrorPage{
		PageHeader: PageHeader{
			Title: "System Error Message",
			NP:    NavPrimary{Active: "records", State: "data"}, // TODO: need appropriate header
			NS:    defaultSecondaryNav,
		},
		Error: err,
	}
}
