package export

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"time"
)

// HKDateTime converts a HK date string '2006-01-02 15:04:05 -0700' to Golang date
func HKDateTime(hkt string) (time.Time, error) {
	if hkt == "" {
		return time.Now(), fmt.Errorf("nothing to coonvert")
	}
	if t, err := time.Parse("2006-01-02 15:04:05 -0700", hkt); err != nil {
		return time.Now(), fmt.Errorf("error converting %q: %v", hkt, err)
	} else {
		return t, nil
	}
}

// yyyymmdd converts a APPLE time stamp "2020-04-19 06:53:41 -0800" to Golang time
// returns yyyy-mm-dd
func yyyymmdd(t string) string {
	if t == "" {
		return t
	}
	// a new 't' overrides incoming t. Look scope
	if t, err := time.Parse("2006-01-02 15:04:05 -0700", t); err != nil {
		return fmt.Sprintf("%v", err)
	} else {
		y, m, d := t.Date()
		return fmt.Sprintf("%4d-%02d-%02d", y, m, d)
	}
}

// shortenHKKey remove prefixes from a Key name.
func shortenHKKey(k string) string {
	sk := strings.Replace(k, "HKCategoryTypeIdentifier", "", 1)
	sk = strings.Replace(sk, "HKQuantityTypeIdentifier", "", 1)
	sk = strings.Replace(sk, "HKMetadataKey", "", 1)
	return sk
}

// printUnknownKeys prints slice of strings with a prefix message "m"
func printUnknownKeys(m string, u []string) {
	if len(u) != 0 {
		if m != "" {
			fmt.Print(m + " ")
		}
		for i, k := range u {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%q", k)
		}
		fmt.Printf("\n")
	}
}

// printUnknownSourcesTypes prefix with message "m"
func printUnknownSourcesTypes(m string, u map[string][]string) {
	if len(u) > 0 {
		if m != "" {
			fmt.Println(m)
		}
		for k, ts := range u {
			fmt.Printf("%s:[", k)
			for i, t := range ts {
				if i != 0 {
					fmt.Print(" ")
				}
				fmt.Printf("%q", t)
			}
			fmt.Println("]")
		}
	}
}

// maps2SortedSlice: Convert map to []KeyCount sorted highest to lowest
func maps2SortedSlice(sim map[string]int) (KeyCounts, error) {
	if len(sim) <= 0 {
		return nil, fmt.Errorf("nothing to process")
	}
	// Convert map to []KeyCounts so we can sort and print highest to lowest
	kv := make(KeyCounts, 0, len(sim))
	for k, v := range sim {
		kv = append(kv, KeyCount{Key: k, Count: v})
	}
	sort.Sort(kv)
	return kv, nil
}

// printStringInts a map with highest count to lowest count.
func printStringInts(m string, sim map[string]int) {
	if len(sim) <= 0 {
		return
	}
	if m != "" {
		fmt.Println(m)
	}
	kv, err := maps2SortedSlice(sim)
	if err != nil {
		return
	}
	total := 0
	for _, v := range kv {
		// Remove long prefixes
		fmt.Printf("\t%s: %d\n", shortenHKKey(v.Key), v.Count)
		total += v.Count
	}
	fmt.Printf("Total Count: %d\n", total)
}

// writeSortedRows
func writeSortedRows(w io.Writer, c1 string, c2 string, sim map[string]int, sep string) {
	if len(sim) <= 0 {
		return
	}
	// Convert map to []KeyCounts so we can sort and print highest to lowest
	kv, err := maps2SortedSlice(sim)
	if err != nil {
		return
	}
	for _, v := range kv {
		_, _ = fmt.Fprintf(w, "%s%s%s%s%s%s%d\n", shortenHKKey(c1), sep, shortenHKKey(c2), sep, shortenHKKey(v.Key), sep, v.Count)
	}

}

// workoutKeys
func workoutKeys(ws []Workout) []string {
	if ws == nil {
		return nil
	}
	kCount := map[string]int{}
	for _, w := range ws {
		md := w.MetadataEntries
		for _, m := range md {
			kCount[m.Key]++
		}
	}

	// Iterate kCount and return names
	keys := make(sort.StringSlice, 0, len(kCount))
	for k := range kCount {
		keys = append(keys, k)
	}
	keys.Sort()
	return keys
}

// unknownWorkoutKeys returns a list if unknown WorkoutMetadataKeys types.
func unknownWorkoutKeys(found []string) []string {
	unknown := sort.StringSlice{}
	// Sort found and known O(nlog(n))
	sort.Strings(found)
	sort.Sort(WorkoutMetadataKeysKnown)
	n := len(WorkoutMetadataKeysKnown)
	for _, f := range found {
		i := sort.Search(n, func(i int) bool { return string(WorkoutMetadataKeysKnown[i]) >= f })
		if i < n && string(WorkoutMetadataKeysKnown[i]) == f {
			continue // Found
		} else {
			unknown = append(unknown, f)
		}
	}
	return unknown
}

// unknownWorkoutActivityTypes returns a list if unknown WorkoutActivityTypes types.
func unknownWorkoutActivityTypes(found []string) []string {
	unknown := sort.StringSlice{}
	// Sort found and known O(nlog(n))
	sort.Strings(found)
	sort.Sort(WorkOutActivityTypesKnown)
	n := len(WorkOutActivityTypesKnown)
	for _, f := range found {
		i := sort.Search(n, func(i int) bool { return string(WorkOutActivityTypesKnown[i]) >= f })
		if i < n && string(WorkOutActivityTypesKnown[i]) == f {
			continue // Found
		} else {
			unknown = append(unknown, f)
		}
	}
	return unknown
}

func workoutActivityTypes(ws []Workout) []string {
	if ws == nil {
		return nil
	}
	kCount := map[string]int{}
	for _, w := range ws {
		md := w.ActivityType
		kCount[md]++
	}

	// Iterate kCount and return names
	keys := make(sort.StringSlice, 0, len(kCount))
	for k := range kCount {
		keys = append(keys, k)
	}
	keys.Sort()
	return keys
}

// recordKeys returns all unique keys in the HK record data
func recordKeys(rs []Record) []string {
	if rs == nil {
		return nil
	}
	kCount := map[string]int{}
	for _, r := range rs {
		md := r.MetadataEntries
		for _, m := range md {
			kCount[m.Key]++
		}
	}

	// Iterate kCount and return names
	keys := make(sort.StringSlice, 0, len(kCount))
	for k := range kCount {
		keys = append(keys, k)
	}
	keys.Sort()
	return keys
}

// recordSources returns all unique sourceName in the HK record data.
// AppleWatch provide millions of HKQuantityTypeIdentifierHeartRate records.
func recordSources(rs []Record) []string {
	if rs == nil {
		return nil
	}
	kCount := map[string]int{}
	for _, r := range rs {
		s := r.SourceName
		kCount[s]++
	}

	// Iterate kCount and return names
	keys := make(sort.StringSlice, 0, len(kCount))
	for k := range kCount {
		keys = append(keys, k)
	}
	keys.Sort()
	return keys
}

// recordTypes returns all possible record types.
func recordTypes(rs []Record) []string {
	if rs == nil {
		return nil
	}
	kCount := map[string]int{}
	for _, r := range rs {
		t := r.Type
		kCount[t]++
	}

	// Iterate kCount and return names
	keys := make(sort.StringSlice, 0, len(kCount))
	for k := range kCount {
		keys = append(keys, k)
	}
	keys.Sort()
	return keys
}

// unknownRecordKeys returns  unseen RecordMetadataKeys
func unknownRecordKeys(found []string) []string {
	unknown := sort.StringSlice{}
	// Sort found and known O(nlog(n))
	sort.Strings(found)
	sort.Sort(RecordMetadataKeysKnown)
	n := len(RecordMetadataKeysKnown)
	for _, f := range found {
		i := sort.Search(n, func(i int) bool { return string(RecordMetadataKeysKnown[i]) >= f })
		if i < n && string(RecordMetadataKeysKnown[i]) == f {
			continue // Found
		} else {
			unknown = append(unknown, f)
		}
	}
	return unknown
}

// unknownRecordTypes returns a list if unknown record types.
func unknownRecordTypes(found []string) []string {
	unknown := sort.StringSlice{}
	// Sort found and known O(nlog(n))
	sort.Strings(found)
	sort.Sort(RecordTypesKnown)
	n := len(RecordTypesKnown)
	for _, f := range found {
		i := sort.Search(n, func(i int) bool { return string(RecordTypesKnown[i]) >= f })
		if i < n && string(RecordTypesKnown[i]) == f {
			continue // Found
		} else {
			unknown = append(unknown, f)
		}
	}
	return unknown
}

// unknownRecordSources returns a list if unknown record Sources.
func unknownRecordSources(found []string) []string {
	unknown := sort.StringSlice{}
	// Sort found and known O(nlog(n))
	sort.Strings(found)
	sort.Sort(RecordSourcesKnown)
	n := len(RecordSourcesKnown)
	for _, f := range found {
		i := sort.Search(n, func(i int) bool { return string(RecordSourcesKnown[i]) >= f })
		if i < n && string(RecordSourcesKnown[i]) == f {
			continue // Found
		} else {
			unknown = append(unknown, f)
		}
	}
	return unknown
}

// unknownSourcesTypes returns a list if unknown record "SourceName" and "Type" they usually return.
func unknownSourcesTypes(found map[string][]string) map[string][]string {
	unknown := make(map[string][]string)

	for k, v := range found { // Iterate through 'found' map
		sn := RecordSource(k)
		if kv, ok := RecordKnownSourcesTypes[sn]; !ok { // SourceName is not Known
			nt := make([]string, 0, len(v)) // new tags
			nt = append(nt, v...)           // Copy slice
			unknown[k] = nt
		} else { // SourceName is found, compare if there are any unknown tags in it
			var nt []string       // News tags, if any
			for _, t := range v { // check each type t in v against known
				match := false
				for _, tt := range kv {
					if t == string(tt) {
						match = true
						break
					}
				}
				if !match {
					nt = append(nt, t)
				}
			}
			if len(nt) > 0 {
				unknown[k] = nt
			}
		}
	}

	return unknown
}

// recordSourcesTypes returns a map of all devices and keys found
func recordSourcesTypes(rs []Record) map[string][]string {
	// Create a map of [device][keys they see]
	recordTypes := make(map[string]map[string]int)
	for _, r := range rs {
		s, t := r.SourceName, r.Type
		if _, ok := recordTypes[s]; !ok {
			recordTypes[s] = make(map[string]int)
			recordTypes[s][t] = 1
		} else {
			if _, ok := recordTypes[s][t]; !ok {
				recordTypes[s][t] = 1
			} else {
				recordTypes[s][t]++
			}
		}
	}

	// Create a map of all keys
	allTypes := make(map[string][]string)
	for k, v := range recordTypes {
		types := make([]string, 0, len(v))
		for k := range v {
			types = append(types, k)
		}
		allTypes[k] = types
	}

	return allTypes
}
