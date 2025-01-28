package node

import (
	"cmp"
)

// Curr -> Left -> Right
func PreOrder[T cmp.Ordered](head *node[T]) []*node[T] {
	if head == nil {
		return []*node[T]{}
	}

	var curr *node[T]
	ret := make([]*node[T], 0, head.TreeDepth()*2+1)
	stk := NewStack[T](head.TreeDepth())
	stk.Push(head)

	for stk.Len() > 0 {
		curr = stk.Pop()

		ret = append(ret, curr)

		if curr.right != nil {
			stk.Push(curr.right)
		}
		if curr.left != nil {
			stk.Push(curr.left)
		}
	}

	return ret
}

// Left -> Curr -> Right
func InOrder[T cmp.Ordered](head *node[T]) []*node[T] {
	if head == nil {
		return []*node[T]{}
	}

	var curr *node[T]
	ret := make([]*node[T], 0, head.TreeDepth()*2+1)
	stk := NewStack[T](head.TreeDepth())
	stk.Push(head)

	for stk.Len() > 0 {
		curr = stk.Peep()

		// Go down left as far as possible
		for curr.left != nil {
			stk.Push(curr.left)
			curr = stk.Peep()
		}

		// Go up, visiting nodes until there's a right branch
		for curr != nil && curr.right == nil {
			ret = append(ret, stk.Pop())
			curr = stk.Peep()
		}

		// Only if we got back to the head
		if curr == nil {
			break
		}

		// Visit current then move right
		ret = append(ret, curr)
		stk.Pop()
		stk.Push(curr.right)
	}

	return ret
}

// Left -> Right -> Curr
func PostOrder[T cmp.Ordered](head *node[T]) []*node[T] {
	if head == nil {
		return []*node[T]{}
	}

	var curr *node[T]
	ret := make([]*node[T], 0, head.TreeDepth()*2+1)
	stk := NewStack[T](head.TreeDepth())
	curr = head

	for stk.Len() > 0 {

		// Go as far left as possible
		for curr != nil {
			stk.Push(curr.GetRight())
			stk.Push(curr)
			curr = curr.GetLeft()
		}

		curr = stk.Pop() // Has no left

		// If no right child, then visit
		// If right child:
		// If right child is on the stack, it hasn't been visited yet: so visit right first and save this node for later
		// If right child isn't on the stack, it's already been seen so visit

		if curr.GetRight() != nil && curr.GetRight() == stk.Peep() {
			right := stk.Pop()
			stk.Push(curr)
			curr = right
		} else {
			ret = append(ret, curr)
		}
	}
	return ret
}
