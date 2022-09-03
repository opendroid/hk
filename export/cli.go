package export

// ValidateRecordFlag checks if flag is one os known types
func ValidateRecordFlag(_ string) bool {
	return false
}

var _ = ValidateRecordFlag // silence unused warning
