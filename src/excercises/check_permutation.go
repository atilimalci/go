package excercises

import (
	"sort"
	"strings"
)

func sortString(w string) <-chan string {
	r := make(chan string)
	go func() {
		s := strings.Split(w, "")
		sort.Strings(s)
		r <- strings.Join(s, "")
		close(r)
	}()
	return r
}

// PermutationBySorting checks permutation by sorting strings
func PermutationBySorting(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	r1 := <-sortString(s)
	r2 := <-sortString(t)

	return r1 == r2
}
