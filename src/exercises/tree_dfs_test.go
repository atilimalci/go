package exercises

import "testing"

func Test_dfs(t *testing.T) {
	n0 := &Node{0, nil, nil}
	n1 := &Node{2, nil, nil}
	n2 := &Node{1, n0, n1}
	n3 := &Node{3, n2, nil}
	n41 := &Node{5, nil, nil}
	n42 := &Node{7, nil, nil}
	n4 := &Node{6, n41, n42}
	n5 := &Node{4, n3, n4}

	r := n5.dfs()
	s := r.String()
	if s != "01234567" {
		t.Errorf("Expected 01234567 but got %s", s)
	}
}
