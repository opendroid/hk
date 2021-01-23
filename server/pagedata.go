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

// PageHeader defines the Title of a page and Name of  Active Primary Nav page. One of "records", "activity" or "summary"
type PageHeader struct {
	Title  string
	Active string // Nav Primary data
	NS     []NameActiveHREF
}

// RecordsPage data for Page records
type IndexPage struct {
	PageHeader
	UserID string
}

// ErrorPage data for Page records
type ErrorPage struct {
	PageHeader
	Error string
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
	recordsSourcesNav = PageHeader{ // For record-xhr-sources.gohtml page
		Title:  "Your Health Records Sources",
		Active: "records",
		NS: []NameActiveHREF{
			{Name: "Sources", Active: true, HREF: recordsDevicesHREF},
			{Name: "Types", Active: false, HREF: recordsTypesHREF},
			{Name: "All", Active: false, HREF: recordsAllHREF},
		},
	}

	recordsTypesNav = PageHeader{ // For record-xhr-types.gohtml page
		Title:  "Your Health Records Types",
		Active: "records",
		NS: []NameActiveHREF{
			{Name: "Sources", Active: false, HREF: recordsDevicesHREF},
			{Name: "Types", Active: true, HREF: recordsTypesHREF},
			{Name: "All", Active: false, HREF: recordsAllHREF},
		},
	}

	recordsAllNav = PageHeader{ // For record-xhr-all.gohtml page
		Title:  "All Your Health Records",
		Active: "records",
		NS: []NameActiveHREF{
			{Name: "Sources", Active: false, HREF: recordsDevicesHREF},
			{Name: "Types", Active: false, HREF: recordsTypesHREF},
			{Name: "All", Active: true, HREF: recordsAllHREF},
		},
	}

	summaryTableNav = PageHeader{ // For summary.gohtml page
		Title:  "Activity Summary Table",
		Active: "summary",
		NS: []NameActiveHREF{
			{Name: "Table", Active: true, HREF: summaryTableHREF},
			{Name: "Graph", Active: false, HREF: summaryGraphHREF},
		},
	}

	summaryGraphNav = PageHeader{ // For summary.gohtml page
		Title:  "Your Activity Graph",
		Active: "summary",
		NS: []NameActiveHREF{
			{Name: "Table", Active: false, HREF: summaryTableHREF},
			{Name: "Graph", Active: true, HREF: summaryGraphHREF},
		},
	}

	defaultSecondaryNav = []NameActiveHREF{ // Default
		{Name: "Sources", Active: false, HREF: recordsDevicesHREF},
		{Name: "Types", Active: false, HREF: recordsTypesHREF},
		{Name: "All", Active: false, HREF: recordsAllHREF},
	}
)

// getRecordsData gets template data for page "Records - data"
func getIndexPageData(userID string) *IndexPage {
	return &IndexPage{
		PageHeader: PageHeader{
			Title: "Your Health Records",
			NS:    defaultSecondaryNav,
		},
		UserID: userID,
	}
}

// getRecordsData gets template data for page "Records - data"
func getRecordsPageData(userID string) *RecordsPage {
	d := RecordsPage{
		PageHeader: PageHeader{
			Title:  "Your records data",
			Active: "records",
			NS:     defaultSecondaryNav,
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
			NS:    defaultSecondaryNav,
		},
		Error: err,
	}
}
