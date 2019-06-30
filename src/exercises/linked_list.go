package exercises

import (
	"strconv"
	"strings"
)

// LinkedListNode is intended to hold a linked list node
type LinkedListNode struct {
	i    int
	next *LinkedListNode
}

// NewNode creates a new LinkedListNode
func NewNode(i int) *LinkedListNode {
	return &LinkedListNode{i, nil}
}

func (head *LinkedListNode) toString() string {
	var r strings.Builder
	p := head
	for p != nil {
		r.WriteString(strconv.Itoa(p.i))
		p = p.next
		if p != nil {
			r.WriteString(",")
		}
	}

	return r.String()
}

func (head *LinkedListNode) removeDuplicates() {
	h := head
	for h != nil {

		nodeOuter := h
		inll := h
		h = h.next

		for inll != nil && inll.next != nil {

			if nodeOuter.i == inll.next.i {
				inll.next = inll.next.next
			}

			inll = inll.next
		}

	}
}
