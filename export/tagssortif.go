package export

// This section implements sort.Interface for various known types
// Len of WorkoutMetadataKeys for sort
func (w WorkoutMetadataKeys) Len() int { return len(w) }

// Less of WorkoutMetadataKeys for sort
func (w WorkoutMetadataKeys) Less(i, j int) bool { return string(w[i]) < string(w[j]) }

// Swap two WorkoutMetadataKeys for sort
func (w WorkoutMetadataKeys) Swap(i, j int) { w[i], w[j] = w[j], w[i] }

// Len of WorkOutActivityTypes for sort
func (w WorkOutActivityTypes) Len() int { return len(w) }

// Less of WorkOutActivityTypes for sort
func (w WorkOutActivityTypes) Less(i, j int) bool { return string(w[i]) < string(w[j]) }

// Swap two WorkOutActivityTypes for sort
func (w WorkOutActivityTypes) Swap(i, j int) { w[i], w[j] = w[j], w[i] }

// Len of RecordMetadataKeys for sort
func (r RecordMetadataKeys) Len() int { return len(r) }

// Less of RecordMetadataKeys for sort
func (r RecordMetadataKeys) Less(i, j int) bool { return string(r[i]) < string(r[j]) }

// Swap two RecordMetadataKeys for sort
func (r RecordMetadataKeys) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

// Len of RecordTypes for sort
func (r RecordTypes) Len() int { return len(r) }

// Less of RecordTypes for sort
func (r RecordTypes) Less(i, j int) bool { return string(r[i]) < string(r[j]) }

// Swap two RecordTypes sort interface
func (r RecordTypes) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

// Len of RecordSources for sort
func (r RecordSources) Len() int { return len(r) }

// Less of RecordSources for sort
func (r RecordSources) Less(i, j int) bool { return string(r[i]) < string(r[j]) }

// Swap two RecordTypes sort interface
func (r RecordSources) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

// keyCount helper to sort data in decreasing order
type KeyCount struct {
	Key   string
	Count int
}
type KeyCounts []KeyCount

// Len of KeyCounts for sort
func (k KeyCounts) Len() int { return len(k) }

// Less of KeyCounts for sort in decreasing order
func (k KeyCounts) Less(i, j int) bool { return k[i].Count > k[j].Count }

// Swap two KeyCounts sort interface
func (k KeyCounts) Swap(i, j int) { k[i], k[j] = k[j], k[i] }
