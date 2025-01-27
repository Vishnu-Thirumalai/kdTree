package node

import (
	"cmp"
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

func (n *node[T]) GetValue() []T {
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

func (n *node[T]) TreeDepth() int {
	curr := n
	depth := 0
	for curr != nil {
		depth += 1
		curr = curr.left
	}
	return depth
}

func NewTreeFromList[T cmp.Ordered](vals [][]T) *node[T] {

	var makeNode func(idx int) *node[T]

	makeNode = func(idx int) *node[T] {
		if idx >= len(vals) || vals[idx] == nil {
			return nil
		}

		curr := NewNode(vals[idx])

		curr.SetLeft(makeNode(idx*2 + 1))
		curr.SetRight(makeNode(idx*2 + 2))

		return curr
	}

	return makeNode(0)
}
