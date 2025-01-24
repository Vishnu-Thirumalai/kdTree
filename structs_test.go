package main

import (
	"fmt"
	"testing"
)

func printNodeChildren(curr *node[string]) {
	v := curr.getValue()
	var l, r []string
	if curr.GetLeft() != nil {
		l = curr.GetLeft().getValue()
	}
	if curr.GetRight() != nil {
		r = curr.GetRight().getValue()
	}
	fmt.Println(v, l, r)
}

func makeTree() *node[string] {
	//    a
	//  b   c
	// d e f g

	vals := []string{"a", "b", "c", "d", "e", "f", "g"}

	var makeNode func(idx int) *node[string]

	makeNode = func(idx int) *node[string] {
		if idx >= len(vals) || vals[idx] == "" {
			return nil
		}

		curr := NewNode([]string{vals[idx]})

		curr.SetLeft(makeNode(idx*2 + 1))
		curr.SetRight(makeNode(idx*2 + 2))

		return curr
	}

	return makeNode(0)
}

func Test_Inorder(t *testing.T) {
	InOrder(makeTree())
}
