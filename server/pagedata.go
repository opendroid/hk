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

var (
	defaultSecondaryNav = []NameActiveHREF{
		{Name: "Devices", Active: false, HREF: recordsDevicesHREF},
		{Name: "Types", Active: false, HREF: recordsTypesHREF},
		{Name: "All", Active: false, HREF: recordsAllHREF},
	}
)

// getRecordsData gets template data for page "Records - data"
func getIndexPageData(userID string) *IndexPage {
	return &IndexPage{
		PageHeader: PageHeader{
			Title: "Your Health Records",
			NS: defaultSecondaryNav,
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
			NS: defaultSecondaryNav,
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
			NS: defaultSecondaryNav,
		},
		Error: err,
	}
}
