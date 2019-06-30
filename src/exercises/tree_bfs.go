package exercises

import (
	"container/list"
	"strconv"
	"strings"
)

func (root *Node) bfs() string {
	var r strings.Builder
	q := list.New()

	curr := root

	for curr != nil {
		r.WriteString(strconv.Itoa(curr.i))
		q.PushBack(curr.left)
		q.PushBack(curr.right)
		curr = pop(q)
	}

	return r.String()
}

func (root *Node) bfsRecursive() string {
	r := new(strings.Builder)
	q := list.New()
	q.PushBack(root)
	bfsRecursive(q, r)
	return r.String()
}

func bfsRecursive(q *list.List, r *strings.Builder) {
	if q.Len() == 0 {
		return
	}
	curr := pop(q)
	if curr == nil {
		return
	}
	r.WriteString(strconv.Itoa(curr.i))
	q.PushBack(curr.left)
	q.PushBack(curr.right)
	bfsRecursive(q, r)
}
