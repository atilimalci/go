package excercises

import (
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func PermutationBySorting(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return sortString(s) == sortString(t)
}
