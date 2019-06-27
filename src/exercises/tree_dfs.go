package exercises

import (
	"container/list"
	"strconv"
	"strings"
)

type Node struct {
	i           int
	left, right *Node
}

func (node *Node) dfs() strings.Builder {
	var sb strings.Builder

	l := list.New()
	curr := node
	for curr != nil || l.Len() > 0 {
		for curr != nil {
			l.PushFront(curr)
			curr = curr.left
		}

		curr = l.Front().Value.(*Node)
		l.Remove(l.Front())
		sb.WriteString(strconv.Itoa(curr.i))
		curr = curr.right
	}
	return sb
}
