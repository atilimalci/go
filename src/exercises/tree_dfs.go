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

//	    1
//	  /   \
//   2     3
// /   \
//4     5
//
// 4 5 2 3 1
func (root *Node) dfsPostorder() string {
	var sb strings.Builder
	s1 := list.New()
	s2 := list.New()

	s1.PushFront(root)
	for s1.Len() > 0 {
		curr := pop(s1)
		s2.PushFront(curr)
		if curr.left != nil {
			s1.PushFront(curr.left)
		}
		if curr.right != nil {
			s1.PushFront(curr.right)
		}
	}

	for e := s2.Front(); e != nil; e = e.Next() {
		n := e.Value.(*Node).i
		sb.WriteString(strconv.Itoa(n))
	}

	return sb.String()
}

func (root *Node) dfsPreorder() string {
	var sb strings.Builder
	s := list.New()

	curr := root
	for curr != nil || s.Len() > 0 {
		for curr != nil {
			s.PushFront(curr)
			sb.WriteString(strconv.Itoa(curr.i))
			curr = curr.left
		}

		curr = pop(s)
		curr = curr.right
	}
	return sb.String()
}

func pop(l *list.List) *Node {
	top := l.Front().Value.(*Node)
	l.Remove(l.Front())
	return top
}

func (root *Node) dfsInorder() string {
	var sb strings.Builder

	s := list.New()
	curr := root
	for curr != nil || s.Len() > 0 {
		for curr != nil {
			s.PushFront(curr)
			curr = curr.left
		}

		curr = pop(s)
		sb.WriteString(strconv.Itoa(curr.i))
		curr = curr.right
	}
	return sb.String()
}

func (root *Node) bfs() strings.Builder {
	var r strings.Builder
	s1, s2 := list.New(), list.New()
	s1.PushFront(root)

	for s1.Len() > 0 {
		node := s1.Front().Value.(*Node)
		s1.Remove(s1.Front())

		s2.PushFront(node)

		if node.left != nil {
			s1.PushFront(node.left)
		}
		if node.right != nil {
			s1.PushFront(node.right)
		}
	}

	for s2.Len() > 0 {
		node := s2.Front().Value.(*Node)
		s2.Remove(s2.Front())
		r.WriteString(strconv.Itoa(node.i))
	}
	return r
}
