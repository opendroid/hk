package export

import (
	"github.com/opendroid/hk/logger"
	"go.uber.org/zap"
	"sort"
)

// RecordsSummary returns a summary of HealthData in a sorted slice.
//
func (h *HealthData) RecordsSummary() NameTypeKeyCounts {
	if h == nil {
		return nil
	}
	// d holds map of summary data
	d := make(map[string]map[string]map[string]int)

	// Get unique [SourceNames][Type][MetadataEntryKey] count
	entries := 0
	for _, v := range h.Records {
		if _, ok := d[v.SourceName]; !ok {
			d[v.SourceName] = make(map[string]map[string]int)
		}
		if _, ok := d[v.SourceName][v.Type]; !ok {
			d[v.SourceName][v.Type] = make(map[string]int)
		}
		for _, md := range v.MetadataEntries {
			mdKey := shortenHKKey(md.Key)
			d[v.SourceName][v.Type][mdKey]++
			entries++
		}
		if len(v.HRV) > 0 {
			for _, h := range v.HRV {
				d[v.SourceName][v.Type]["HRV"] += len(h.BPM)
				entries++
			}
		} else if len(v.MetadataEntries) == 0 {
			d[v.SourceName][v.Type]["-"]++
			entries++
		}
	}
	// Convert map to an array NameTypeKeyCounts{{Name, Type, Key, Count},{}}
	rs := make(NameTypeKeyCounts, 0, entries)
	total := 0
	for ks, vs := range d { // SourceNames
		for kt, vt := range vs { // Types
			if kc, err := maps2SortedSlice(vt); err != nil {
				logger.Error("RecordsSummary maps2SortedSlice error",
					zap.String("info", err.Error()),
					zap.Int("entries", entries))
				continue
			} else {
				for _, v := range kc {
					rs = append(rs, NameTypeKeyCount{ks, shortenHKKey(kt), v})
					total += v.Count
				}
			}
		}
	}
	sort.Sort(rs) // Sort each type by decreasing order
	// Add total as last element
	rs = append(rs, NameTypeKeyCount{"Total", "-", KeyCount{"-", total}})
	return rs
}
