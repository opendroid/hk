package server

import (
	"github.com/opendroid/hk/export"
	"github.com/opendroid/hk/logger"
	"go.uber.org/zap"
)

// NameActive a tuple that saves name and is a section is active.
type NameActive struct {
	Name   string
	Active bool
}

type PageHeader struct {
	Title  string
	Active string // Nav Primary data
	NS     []NameActive
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
}

// getRecordsData gets template data for page "Records - data"
func getIndexPageData(userID string) *IndexPage {
	return &IndexPage{
		PageHeader: PageHeader{
			Title: "Your Health Records",
			NS: []NameActive{
				{Name: "Devices", Active: false},
				{Name: "Types", Active: false},
				{Name: "Metadata", Active: false},
			},
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
			NS: []NameActive{
				{Name: "Devices", Active: true},
				{Name: "Types", Active: false},
				{Name: "Metadata", Active: false},
			},
		},
	}
	u, ok := users[userID]
	if !ok {
		logger.Error("No user data", zap.String("user", userID))
		return nil
	}
	d.Records = u.health.RecordsSummary()
	return &d
}

// getErrorPageData gets template data for page "Records - data"
func getErrorPageData(err string) *ErrorPage {
	return &ErrorPage{
		PageHeader: PageHeader{
			Title: "System Error Message",
			NS: []NameActive{
				{Name: "Devices", Active: false},
				{Name: "Types", Active: false},
				{Name: "Metadata", Active: false},
			},
		},
		Error: err,
	}
}