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

func (root *Node) dfsPosterderRecursive() string {
	sb := new(strings.Builder)
	dfsPosterderRecursive(root, sb)
	return sb.String()
}

func dfsPosterderRecursive(node *Node, sb *strings.Builder) {
	if node == nil {
		return
	}
	dfsPosterderRecursive(node.left, sb)
	dfsPosterderRecursive(node.right, sb)
	sb.WriteString(strconv.Itoa(node.i))
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

func (root *Node) dfsPreorderRecursive() string {
	sb := new(strings.Builder)
	dfsPreorderRecursive(root, sb)
	return sb.String()
}

func dfsPreorderRecursive(node *Node, sb *strings.Builder) {
	if node == nil {
		return
	}
	sb.WriteString(strconv.Itoa(node.i))
	dfsPreorderRecursive(node.left, sb)
	dfsPreorderRecursive(node.right, sb)
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

func (root *Node) dfsInorderRecursive() string {
	sb := new(strings.Builder)
	dfsInorderRecursive(root, sb)
	return sb.String()
}

func dfsInorderRecursive(node *Node, sb *strings.Builder) {
	if node == nil {
		return
	}
	dfsInorderRecursive(node.left, sb)
	sb.WriteString(strconv.Itoa(node.i))
	dfsInorderRecursive(node.right, sb)
}
