package main

import (
	"errors"
	"fmt"
)

type node struct {
	data int
	next *node
}

func push(newData int, nodeList *node) *node {
	newNode := &node{newData, nil}
	newNode.next = nodeList
	return newNode
}

func pop(nodeList *node) *node {
	if nodeList != nil {
		nodeList = nodeList.next
	}
	return nodeList
}

func peek(nodeList *node) (int, error) {
	if nodeList != nil {
		return nodeList.data, nil
	}
	return -1, errors.New("Empty list")
}

func main() {
	a := &node{5, nil}
	numberList := a
	numberList = push(10, numberList)
	numberList = push(15, numberList)
	numberList = push(20, numberList)

	data, _ := peek(numberList)
	fmt.Println("Before pop ", data)

	numberList = pop(numberList)

	data, _ = peek(numberList)
	fmt.Println("After pop ", data)

	for n := numberList; ; n = n.next {
		fmt.Println(n.data)
		if n.next == nil {
			break
		}
	}

}
