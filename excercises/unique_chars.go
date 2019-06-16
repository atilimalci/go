package excercises

// IsUniqueChars determines if a string has all unique characters
func IsUniqueChars(str string) bool {
	if len(str) > 128 {
		return false
	}

	var charSet [128]bool

	for i := 0; i < len(str); i++ {
		val := str[i]
		if charSet[val] {
			return false
		}
		charSet[val] = true
	}

	return true
}
