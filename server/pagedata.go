package server

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

// RecordsPage data for Page records
type RecordsPage struct {
	PageHeader
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
func getRecordsPageData() *RecordsPage {
	return &RecordsPage{
		PageHeader{
			Title:  "Your records data",
			Active: "records",
			NS: []NameActive{
				{Name: "Devices", Active: true},
				{Name: "Types", Active: false},
				{Name: "Metadata", Active: false},
			},
		},
	}
}
