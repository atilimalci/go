package main

import (
	"errors"
	"fmt"
)

type node struct {
	data int
	next *node
}

var top *node

func push(newData int) {
	var newNode = node{newData, nil}
	newNode.next = top
	top = &newNode
}

func peek() (int, error) {
	if top != nil {
		return top.data, nil
	}
	return -1, errors.New("Empty list")
}

func pop() (int, error) {

	if top != nil {
		data := top.data
		top = top.next
		return data, nil
	}
	return -1, errors.New("Empty list")
}

func isEmpty() bool {
	return top == nil
}

func main() {
	fmt.Println("Is empty initially", isEmpty())
	push(10)
	fmt.Println("Is empty after push", isEmpty())
	push(15)
	push(20)
	data, _ := peek()
	fmt.Println("Before pop", data)
	var popped, _ = pop()
	fmt.Println("Popped data", popped)
	data, _ = peek()
	fmt.Println("After pop ", data)

	for n := top; ; n = n.next {
		fmt.Println(n.data)
		if n.next == nil {
			break
		}
	}

}
