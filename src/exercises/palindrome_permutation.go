package exercises

func IsPermutationOfPalindrome(str string) bool {
	var cb [26]int
	a := "a"[0]
	for i := 0; i < len(str); i++ {
		cb[str[i]-a]++
	}

	var hangingCenter bool
	for i := 0; i < len(cb); i++ {
		if cb[i] == 1 {
			if hangingCenter {
				return false
			}
			hangingCenter = true
			continue
		}
		if cb[i]%2 != 0 {
			return false
		}
	}
	return true
}
