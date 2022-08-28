package export

import (
	"testing"
	"time"
)

// Yyyymmdd For unit testing, not exported as part of package
var Yyyymmdd = yyyymmdd // func(string) string

// dateTest convenience struct
type dateTest struct {
	title string
	in    string
	want  string
}

// TestYyyymmdd tests static func
// go test -run=TestYyyymmdd -v
func TestYyyymmdd(t *testing.T) {
	data := []dateTest{{
		title: "Test 2016-12-09",
		in:    "2016-12-09 09:52:21 -0800",
		want:  "2016-12-09",
	}, {
		title: "Test 1999-06-01",
		in:    "1999-06-01 09:52:21 -0800",
		want:  "1999-06-01",
	},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			r := Yyyymmdd(d.in)
			if r != d.want {
				t.Fail()
			}
		})
	}
}

// TestHKDateTime test HKDateTime utility
// go test -run=TestHKDateTime -v
func TestHKDateTime(t *testing.T) {
	data := []dateTest{{
		title: "Test Timezone -0800",
		in:    "2016-12-09 09:52:21 -0800",
		want:  "Local",
	}}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			d, err := HKDateTime(d.in)
			if err != nil {
				t.Fail()
				return
			}
			t.Logf("Time: %v, Location: %s, Kitchen: %s\n", d, d.Location(), d.Format(time.Kitchen))
		})
	}
}
