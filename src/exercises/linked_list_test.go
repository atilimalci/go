package exercises

import (
	"testing"
)

func TestLinkedListNode_removeDuplicates(t *testing.T) {
	//12->11->12->21->12->43->21
	head := NewNode(12)
	head.next = NewNode(11)
	head.next.next = NewNode(12)
	head.next.next.next = NewNode(21)
	head.next.next.next.next = NewNode(12)
	head.next.next.next.next.next = NewNode(43)
	head.next.next.next.next.next.next = NewNode(21)

	head.removeDuplicates()

	want := "12,11,21,43"

	if s := head.toString(); s != want {
		t.Errorf("Node.dfsInorder() = %v, want %v", s, want)
	}

}
