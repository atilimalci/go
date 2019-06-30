package exercises

import (
	"container/list"
)

//IsParanthesesBalanced checks if the given inputs paranteses are balanced or not
func IsParanthesesBalanced(s string) bool {
	st := list.New()

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '{' || c == '[' {
			st.PushFront(c)
			continue
		}
		switch c {
		case '}':
			if !checkStack(st, '{') {
				return false
			}
		case ']':
			if !checkStack(st, '[') {
				return false
			}
		case ')':
			if !checkStack(st, '(') {
				return false
			}
		}
	}
	size := st.Len()
	return size == 0
}

func checkStack(st *list.List, b byte) bool {
	h := st.Front()
	st.Remove(st.Front())
	if h == nil || h.Value.(byte) == b {
		return true
	}
	return false
}
