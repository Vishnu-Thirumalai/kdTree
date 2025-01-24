package main

import (
	"cmp"
	"fmt"
)

type node[T cmp.Ordered] struct {
	value []T
	left  *node[T]
	right *node[T]
}

func NewNode[T cmp.Ordered](val []T) *node[T] {
	return &node[T]{
		value: val,
	}
}

func (n *node[T]) getValue() []T {
	return n.value
}

func (n *node[T]) SetLeft(child *node[T]) {
	n.left = child
}

func (n *node[T]) SetRight(child *node[T]) {
	n.right = child
}

func (n *node[T]) GetLeft() *node[T] {
	return n.left
}

func (n *node[T]) GetRight() *node[T] {
	return n.right
}

func InOrder[T cmp.Ordered](head *node[T]) {
	if head == nil {
		return
	}

	stk := []*node[T]{head}
	l := 1

	for l > 0 {
		l -= 1
		curr := stk[l]
		stk = stk[:l]

		fmt.Println(curr.getValue())

		if curr.right != nil {
			stk = append(stk, curr.right)
			l += 1
		}
		if curr.left != nil {
			stk = append(stk, curr.left)
			l += 1
		}
	}
}
